package ctrlmesh

import (
	"github.com/openziti/dilithium/cf"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Identity         string `cf:"identity"`
	BindAddress      string `cf:"bind_address"`
	AdvertiseAddress string `cf:"advertise_address"`
	InitialPeerList  string `cf:"initial_peer"`
}

func Load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading config '%s'", path)
	}
	dataMap := make(map[interface{}]interface{})
	if err = yaml.Unmarshal(data, dataMap); err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal config '%s'", path)
	}
	cfg := &Config{}
	if err := cf.Load(cf.MapIToMapS(dataMap), cfg); err != nil {
		return nil, errors.Wrapf(err, "unable to load config '%s'", path)
	}
	return cfg, nil
}
