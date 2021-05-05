package ctrlmesh

import (
	"github.com/pkg/errors"
)

type PeerList struct {
	cfg       *PeerListConfig
	self      *peer
	peers     []*peer
	listeners []*peerListener
}

func NewPeerList(cfg *PeerListConfig) (*PeerList, error) {
	var listeners []*peerListener
	for _, addr := range cfg.listeners {
		listener, err := newPeerListener(cfg.id, addr)
		if err != nil {
			return nil, errors.Wrapf(err, "error creating listener at [%s]", addr)
		}
		listeners = append(listeners, listener)
	}
	pl := &PeerList{cfg, nil, nil, listeners}
	return pl, nil
}
