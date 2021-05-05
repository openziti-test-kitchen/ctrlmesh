package ctrlmesh

import (
	"github.com/openziti/foundation/transport"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type PeerListConfig struct {
	initialPeers []transport.Address
	listeners    []transport.Address
	advertise    []transport.Address
}

func LoadPeerListConfigYaml(path string) (*PeerListConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading data from [%s]", path)
	}
	dataMap := make(map[string]interface{})
	if err := yaml.Unmarshal(data, dataMap); err != nil {
		return nil, errors.Wrapf(err, "error unmarshaling from [%s]", path)
	}
	return LoadPeerListConfig(dataMap)
}

func LoadPeerListConfig(data map[string]interface{}) (*PeerListConfig, error) {
	plc := &PeerListConfig{}

	if v, found := data["initial_peers"]; found {
		subarr, ok := v.([]interface{})
		if !ok {
			return nil, errors.Errorf("mailformed 'initial_peers' list (%s)", reflect.TypeOf(subarr))
		}
		for _, v := range subarr {
			initialPeer, ok := v.(string)
			if !ok {
				return nil, errors.Errorf("malformed initial peer (%s)", reflect.TypeOf(v))
			}
			initialPeerAddr, err := transport.ParseAddress(initialPeer)
			if err != nil {
				return nil, errors.Wrapf(err, "error parsing 'initial_peers' address [%s]", initialPeer)
			}
			plc.initialPeers = append(plc.initialPeers, initialPeerAddr)
		}
	} else {
		return nil, errors.New("no 'initial_peers' specified")
	}

	if v, found := data["listeners"]; found {
		subarr, ok := v.([]interface{})
		if !ok {
			return nil, errors.Errorf("malformed 'listeners' list (%s)", reflect.TypeOf(subarr))
		}
		for _, v := range subarr {
			listener, ok := v.(string)
			if !ok {
				return nil, errors.Errorf("malformed listener (%s)", reflect.TypeOf(v))
			}
			listenerAddr, err := transport.ParseAddress(listener)
			if err != nil {
				return nil, errors.Wrapf(err, "error parsing 'listener' address [%s]", listener)
			}
			plc.listeners = append(plc.listeners, listenerAddr)
		}
	}

	if v, found := data["ads"]; found {
		subarr, ok := v.([]interface{})
		if !ok {
			return nil, errors.Errorf("malformed 'advertise' list (%s)", reflect.TypeOf(subarr))
		}
		for _, v := range subarr {
			advertise, ok := v.(string)
			if !ok {
				return nil, errors.Errorf("malformed 'advertise' list (%s)", reflect.TypeOf(subarr))
			}
			advertiseAddr, err := transport.ParseAddress(advertise)
			if err != nil {
				return nil, errors.Wrapf(err, "error parsing 'advertise' address [%s]", advertise)
			}
			plc.advertise = append(plc.advertise, advertiseAddr)
		}
	}

	return plc, nil
}
