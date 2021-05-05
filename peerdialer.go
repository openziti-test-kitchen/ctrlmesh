package ctrlmesh

import (
	"fmt"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"github.com/sirupsen/logrus"
)

type peerDialer struct {
	id   *identity.TokenId
	addr transport.Address
	ch   channel2.Channel
}

func newPeerDialer(id string, addr transport.Address) *peerDialer {
	return &peerDialer{&identity.TokenId{Token: id}, addr, nil}
}

func (self *peerDialer) run() {
	dialer := channel2.NewReconnectingDialerWithHandler(self.id, self.addr, nil, self.reconnect)
	ch, err := channel2.NewChannel(fmt.Sprintf("peer_%s", self.addr), dialer, nil)
	if err != nil {
		logrus.Errorf("dial failed (%v)", err)
	}
	self.ch = ch
	logrus.Infof("dial succeeded for [%s]", ch.LogicalName())
}

func (self *peerDialer) reconnect() {
	logrus.Warnf("reconnect handling for [%s]", self.ch.LogicalName())
}
