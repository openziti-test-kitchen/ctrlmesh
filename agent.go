package ctrlmesh

import (
	"github.com/hashicorp/memberlist"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Agent struct {
	cfg           *Config
	agentDelegate memberlist.Delegate
	memberlist    *memberlist.Memberlist
}

func NewAgent(cfg *Config) (*Agent, error) {
	mlCfg := memberlist.DefaultLocalConfig()

	bindAddress, bindPort, err := splitAddress(cfg.BindAddress)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid bind_address '%s'", cfg.BindAddress)
	}
	mlCfg.BindAddr = bindAddress
	mlCfg.BindPort = bindPort

	if cfg.AdvertiseAddress != "" {
		advertiseAddress, advertisePort, err := splitAddress(cfg.AdvertiseAddress)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid advertise_address '%s'", cfg.AdvertiseAddress)
		}
		mlCfg.AdvertiseAddr = advertiseAddress
		mlCfg.AdvertisePort = advertisePort
	}

	mlCfg.Name = cfg.Identity
	mlCfg.Delegate = &agentDelegate{cfg: cfg}
	mlCfg.Logger = logrus.StandardLogger()

	ml, err := memberlist.Create(mlCfg)
	if err != nil {
		return nil, errors.Wrap(err, "error creating agent")
	}

	return &Agent{
		cfg:           cfg,
		agentDelegate: mlCfg.Delegate,
		memberlist:    ml,
	}, nil
}

func (self *Agent) Join() error {
	initialPeers := strings.Split(self.cfg.InitialPeerList, " ")
	if self.cfg.InitialPeerList != "" && len(initialPeers) > 0 {
		count, err := self.memberlist.Join(initialPeers)
		if err != nil {
			return errors.Wrap(err, "error joining control plane")
		}
		logrus.Infof("joined control plane with [%d] nodes", count)
	}
	return nil
}

func (self *Agent) Status() {
	nodes := self.memberlist.Members()
	logrus.Infof("%d nodes:", len(nodes))
	for i, node := range nodes {
		logrus.Infof("#%d: %s/%s: %s", i, node.Name, node.Address(), string(node.Meta))
	}
}

func splitAddress(addr string) (string, int, error) {
	tokens := strings.Split(addr, ":")
	if len(tokens) != 2 {
		return "", -1, errors.Errorf("malformed address '%s'", addr)
	}
	port, err := strconv.Atoi(tokens[1])
	if err != nil {
		return "", -1, errors.Wrapf(err, "bad port '%s'", tokens[1])
	}
	return tokens[0], port, nil
}

type agentDelegate struct{
	cfg *Config
}

func (self *agentDelegate) NodeMeta(limit int) []byte {
	return nil
}

func (self *agentDelegate) NotifyMsg(msg []byte) {
	logrus.Infof("received msg [%s]", string(msg))
}

func (self *agentDelegate) GetBroadcasts(overhead, limit int) [][]byte {
	return nil
}

func (self *agentDelegate) LocalState(join bool) []byte {
	return nil
}

func (self *agentDelegate) MergeRemoteState(buf []byte, join bool) {
	logrus.Infof("merge [%d] bytes, join = %t", len(buf), join)
}

type State struct {
	DataListener string `json:"dl"`
}
