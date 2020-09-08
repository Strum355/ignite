package flag

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/weaveworks/ignite/pkg/network"
)

type DockerNetworkNameFlag struct {
	value *network.DockerNetworkName
}

func (nf *DockerNetworkNameFlag) Set(val string) error {
	*nf.value = network.DockerNetworkName(val)
	return nil
}

func (nf *DockerNetworkNameFlag) String() string {
	if nf.value == nil {
		return ""
	}
	return nf.value.String()
}

func (nf *DockerNetworkNameFlag) Type() string {
	return "docker-network-name"
}

var _ pflag.Value = &DockerNetworkNameFlag{}

func DockerNetworkNameVar(fs *pflag.FlagSet, ptr *network.DockerNetworkName) {
	fs.Var(&DockerNetworkNameFlag{value: ptr}, "docker-network-name", fmt.Sprintf("Name of the docker network to use."))
}
