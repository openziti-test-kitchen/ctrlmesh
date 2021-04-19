package ctrlmesh

import (
	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
)

type SerfAgent struct {
	cfg  *Config
	serf *serf.Serf
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

	sf, err := serf.Create(sCfg)
	if err != nil {
		return nil, errors.Wrap(err, "error creating serf")
	}

	return &SerfAgent{
		cfg:  cfg,
		serf: sf,
	}, nil
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