package ctrlmesh

import (
	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type SerfAgent struct {
	cfg     *Config
	serf    *serf.Serf
	eventCh chan serf.Event
}

func NewSerfAgent(cfg *Config) (*SerfAgent, error) {
	sCfg := serf.DefaultConfig()

	bindAddress, bindPort, err := splitAddress(cfg.BindAddress)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid bind_address '%s'", cfg.BindAddress)
	}
	sCfg.MemberlistConfig.BindAddr = bindAddress
	sCfg.MemberlistConfig.BindPort = bindPort

	if cfg.AdvertiseAddress != "" {
		advertiseAddress, advertisePort, err := splitAddress(cfg.AdvertiseAddress)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid advertise_address '%s'", cfg.AdvertiseAddress)
		}
		sCfg.MemberlistConfig.AdvertiseAddr = advertiseAddress
		sCfg.MemberlistConfig.AdvertisePort = advertisePort
	}
	sCfg.NodeName = cfg.Identity

	eventCh := make(chan serf.Event, 4)
	sCfg.EventCh = eventCh

	sf, err := serf.Create(sCfg)
	if err != nil {
		return nil, errors.Wrap(err, "error creating serf")
	}

	agent := &SerfAgent{
		cfg:     cfg,
		serf:    sf,
		eventCh: eventCh,
	}
	go agent.handleEvents()

	return agent, nil
}

func (self *SerfAgent) Join() error {
	tags := make(map[string]string)
	tags["data_listener"] = self.cfg.DataListener

	initialPeers := strings.Split(self.cfg.InitialPeerList, " ")
	if self.cfg.InitialPeerList != "" && len(initialPeers) > 0 {
		count, err := self.serf.Join(initialPeers, false)
		if err != nil {
			return errors.Wrap(err, "error joining control plane")
		}
		if err := self.serf.SetTags(tags); err != nil {
			return errors.Wrap(err, "error setting node tags")
		}
		logrus.Infof("joined control plane with [%d] nodes", count)
	}
	return nil
}

func (self *SerfAgent) Status() {
	nodes := self.serf.Members()
	logrus.Infof("%d nodes:", len(nodes))
	for i, node := range nodes {
		logrus.Infof("#%d %s/%s (%v)", i, node.Name, node.Addr, node.Tags)
	}
}

func (self *SerfAgent) Query() {
	response, err := self.serf.Query("hello", []byte("oh, wow!"), self.serf.DefaultQueryParams())
	if err != nil {
		logrus.Errorf("query failed (%v)", err)
	}
	select {
	case rv := <-response.ResponseCh():
		logrus.Infof("response: @%s: %s", rv.From, string(rv.Payload))
		response.Close()

	case <-time.After(2 * time.Second):
		logrus.Warnf("no response")
		response.Close()
	}
}

func (self *SerfAgent) handleEvents() {
	for {
		select {
		case event := <-self.eventCh:
			logrus.Infof("received [%s]", event)
		}
	}
}
