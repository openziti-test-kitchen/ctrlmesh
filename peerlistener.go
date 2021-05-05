package ctrlmesh

import (
	"crypto/x509"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type peerListener struct {
	id *identity.TokenId
	l  channel2.UnderlayListener
	a  *accepter
}

func newPeerListener(id string, bind transport.Address) (*peerListener, error) {
	tid := &identity.TokenId{Token: id}
	l := channel2.NewClassicListener(tid, bind, channel2.DefaultConnectOptions(), nil)
	if err := l.Listen(&connectHandler{}); err != nil {
		return nil, errors.Wrap(err, "error creating listener")
	}
	a := newAccepter(l)
	go a.run()
	return &peerListener{tid, l, a}, nil
}

type connectHandler struct{}

func newConnectHandler() *connectHandler {
	return &connectHandler{}
}

func (self *connectHandler) HandleConnection(hello *channel2.Hello, certificates []*x509.Certificate) error {
	return nil
}

type accepter struct {
	l channel2.UnderlayListener
}

func newAccepter(l channel2.UnderlayListener) *accepter {
	return &accepter{l}
}

func (self *accepter) run() {
	logrus.Info("started")
	defer logrus.Info("exited")

	for {
		if ch, err := channel2.NewChannel("ctrl", self.l, nil); err == nil {
			_ = ch.Close()
			logrus.Infof("accepted and closed")
		}
	}
}
