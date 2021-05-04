package ctrlmesh

import "github.com/openziti/foundation/transport"

type PeerList struct {
	cfg   *PeerListConfig
	self  *peer
	peers []*peer
}

type PeerListConfig struct {
	initialPeerAds []transport.Address
}
