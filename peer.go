package ctrlmesh

import (
	"github.com/openziti/foundation/transport"
)

type peer struct {
	id      string
	ads     []transport.Address
	updated uint32
}
