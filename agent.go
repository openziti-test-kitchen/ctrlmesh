package ctrlmesh

import "github.com/openziti/foundation/transport"

type Agent interface {
	Start()
	AddListener(address transport.Address)
	Stop()
}
