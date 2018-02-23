package specs

import (
	"context"

	"github.com/alibaba/pouch/daemon/mgr"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// SpecWrapper wraps the container's specs and add manager operations.
type SpecWrapper struct {
	s *specs.Spec

	ctrMgr mgr.ContainerMgr
	volMgr mgr.VolumeMgr
	netMgr mgr.NetworkMgr
}

// SetupFunc defines spec setup function type.
type SetupFunc func(ctx context.Context, m *mgr.ContainerMeta, s *SpecWrapper) error

var setupFunc = []SetupFunc{
	// process
	setupProcessArgs,
	setupProcessCwd,
	setupProcessEnv,
	setupProcessTTY,
	setupProcessUser,
	setupCap,

	// cgroup
	setupCgroupCPUShare,
	setupCgroupCPUSet,
	setupCgroupMemory,
	setupCgroupMemorySwap,
	setupCgroupMemorySwappiness,

	// namespaces
	setupUserNamespace,
	setupNetworkNamespace,
	setupIpcNamespace,
	setupPidNamespace,
	setupUtsNamespace,

	// volume spec
	setupMounts,

	// network spec
	setupNetwork,

	// host device spec
	setupDevices,

	// linux-platform-specifc spec
	setupSysctl,
	setupAppArmor,
	setupCapabilities,
	setupSeccomp,

	// blkio spec
	setupBlkio,

	// IntelRdtL3Cbm
	setupIntelRdt,

	// alios options
	setupAliOsOption,
}

// Register is used to registe spec setup function.
func Register(f SetupFunc) {
	if setupFunc == nil {
		setupFunc = make([]SetupFunc, 0)
	}
	setupFunc = append(setupFunc, f)
}

// SetupFuncs returns all the spec setup functions.
func SetupFuncs() []SetupFunc {
	return setupFunc
}
