// +build !linux

package libpod

import (
	cnitypes "github.com/containernetworking/cni/pkg/types/current"
	"github.com/containers/podman/v2/libpod/define"
)

func (r *Runtime) setupRootlessNetNS(ctr *Container) error {
	return define.ErrNotImplemented
}

func (r *Runtime) setupSlirp4netns(ctr *Container) error {
	return define.ErrNotImplemented
}

func (r *Runtime) setupNetNS(ctr *Container) error {
	return define.ErrNotImplemented
}

func (r *Runtime) teardownNetNS(ctr *Container) error {
	return define.ErrNotImplemented
}

func (r *Runtime) createNetNS(ctr *Container) error {
	return define.ErrNotImplemented
}

func (c *Container) getContainerNetworkInfo() (*define.InspectNetworkSettings, error) {
	return nil, define.ErrNotImplemented
}

func (r *Runtime) reloadContainerNetwork(ctr *Container) ([]*cnitypes.Result, error) {
	return nil, define.ErrNotImplemented
}

func getCNINetworksDir() (string, error) {
	return "", define.ErrNotImplemented
}
