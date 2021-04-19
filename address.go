package ctrlmesh

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func splitAddress(addr string) (string, int, error) {
	tokens := strings.Split(addr, ":")
	if len(tokens) != 2 {
		return "", -1, errors.Errorf("malformed address '%s'", addr)
	}
	port, err := strconv.Atoi(tokens[1])
	if err != nil {
		return "", -1, errors.Wrapf(err, "bad port '%s'", tokens[1])
	}
	return tokens[0], port, nil
}
