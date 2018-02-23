package specs

import (
	"context"

	"github.com/alibaba/pouch/daemon/mgr"
)

func setupNetwork(ctx context.Context, c *mgr.ContainerMeta, spec *SpecWrapper) error {
	s := spec.s

	s.Hostname = c.Config.Hostname.String()
	//TODO setup network parameters

	return nil
}
