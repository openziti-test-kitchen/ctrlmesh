package ctrlmesh

import (
	"github.com/hashicorp/memberlist"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type Agent struct {
	memberlist *memberlist.Memberlist
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

	ml, err := memberlist.Create(mlCfg)
	if err != nil {
		return nil, errors.Wrap(err, "error creating agent")
	}

	return &Agent{memberlist: ml}, nil
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