package ctrlmesh

import (
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"time"
)

type peer struct {
	ctrl    channel2.Channel
	id      *identity.TokenId
	ads     []transport.Address
	updated time.Time
}
