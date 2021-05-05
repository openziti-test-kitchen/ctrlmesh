package ctrlmesh

import (
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/transport"
)

type peer struct {
	ctrl    channel2.Channel
	id      string
	ads     []transport.Address
	updated uint32
}
