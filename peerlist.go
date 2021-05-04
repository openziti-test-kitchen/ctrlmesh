package ctrlmesh

type PeerList struct {
	cfg   *PeerListConfig
	self  *peer
	peers []*peer
}

func NewPeerList(cfg *PeerListConfig) (*PeerList, error) {
	return nil, nil
}