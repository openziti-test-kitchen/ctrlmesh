module github.com/openziti-incubator/ctrlmesh

go 1.16

replace github.com/hashicorp/memberlist => ../memberlist

replace github.com/hashicorp/serf => ../serf

require (
	github.com/hashicorp/serf v0.9.5
	github.com/michaelquigley/pfxlog v0.3.7
	github.com/openziti/dilithium v0.3.3
	github.com/openziti/foundation v0.15.50
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
)
