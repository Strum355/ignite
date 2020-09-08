package flag

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/weaveworks/ignite/pkg/network"
)

type DockerNetworkNameFlag struct {
	value *network.NetworkName
}

func (nf *DockerNetworkNameFlag) Set(val string) error {
	// TODO: lol
	/* for _, plugin := range plugins {
		if plugin.String() == val {
			*nf.value = plugin
			return nil
		}
	}
	return fmt.Errorf("invalid network plugin %q, must be one of %v", val, plugins) */

	*nf.value = "ignite"
	return nil
}

func (nf *DockerNetworkNameFlag) String() string {
	if nf.value == nil {
		return ""
	}
	return nf.value.String()
}

func (nf *DockerNetworkNameFlag) Type() string {
	return "networkname"
}

var _ pflag.Value = &DockerNetworkNameFlag{}

func NetworkNameVar(fs *pflag.FlagSet, ptr *network.NetworkName) {
	fs.Var(&DockerNetworkNameFlag{value: ptr}, "network-name", fmt.Sprintf("Docker network to use."))
}
