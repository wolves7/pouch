package mgr

import (
	"fmt"
	"time"

	"github.com/alibaba/pouch/apis/types"
	"github.com/alibaba/pouch/pkg/utils"
)

// IsRunning returns container is running or not.
func (c *Container) IsRunning() bool {
	c.Lock()
	defer c.Unlock()

	return c.State.Status == types.StatusRunning
}

// IsStopped returns container is stopped or not.
func (c *Container) IsStopped() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusStopped
}

// IsExited returns container is exited or not.
func (c *Container) IsExited() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusExited
}

// IsCreated returns container is created or not.
func (c *Container) IsCreated() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusCreated
}

// IsPaused returns container is paused or not.
func (c *Container) IsPaused() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusPaused
}

// IsRemoving returns container is removing or not.
func (c *Container) IsRemoving() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusRemoving
}

// IsDead returns container is dead or not.
func (c *Container) IsDead() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusDead
}

// IsRunningOrPaused returns true of container is running or paused.
func (c *Container) IsRunningOrPaused() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusRunning || c.State.Status == types.StatusPaused
}

// IsRestarting returns container is restarting or not.
func (c *Container) IsRestarting() bool {
	c.Lock()
	defer c.Unlock()
	return c.State.Status == types.StatusRestarting

}

// ExitCode returns container's ExitCode.
func (c *Container) ExitCode() int64 {
	c.Lock()
	defer c.Unlock()
	return c.State.ExitCode
}

// SetStatusRunning sets a container to be status running.
// When a container's status turns to StatusStopped, the following fields need updated:
// Status -> StatusRunning
// StartAt -> time.Now()
// Pid -> input param
// ExitCode -> 0
func (c *Container) SetStatusRunning(pid int64) error {
	c.Lock()
	defer c.Unlock()

	if !stateMachine[c.State.Status][types.StatusRunning] {
		return fmt.Errorf("container %s cannot transit from status %v to %v", c.ID, c.State.Status, types.StatusRunning)
	}
	c.State.Status = types.StatusRunning
	c.State.StartedAt = time.Now().UTC().Format(utils.TimeLayout)
	c.State.Pid = pid
	c.State.ExitCode = 0

	return nil
}

// SetStatusStopped sets a container to be status stopped.
// When a container's status turns to StatusStopped, the following fields need updated:
// Status -> StatusStopped
// FinishedAt -> time.Now()
// Pid -> -1
// ExitCode -> input param
// Error -> input param
func (c *Container) SetStatusStopped(exitCode int64, errMsg string) error {
	c.Lock()
	defer c.Unlock()

	if !stateMachine[c.State.Status][types.StatusStopped] {
		return fmt.Errorf("container %s cannot transit from status %v to %v", c.ID, c.State.Status, types.StatusStopped)
	}

	c.State.Status = types.StatusStopped
	c.State.FinishedAt = time.Now().UTC().Format(utils.TimeLayout)
	c.State.Pid = -1
	c.State.ExitCode = exitCode
	c.State.Error = errMsg

	return nil
}

// SetStatusExited sets a container to be status exited.
func (c *Container) SetStatusExited(exitCode int64, errMsg string) error {
	c.Lock()
	defer c.Unlock()

	if !stateMachine[c.State.Status][types.StatusExited] {
		return fmt.Errorf("container %s cannot transit from status %v to %v", c.ID, c.State.Status, types.StatusExited)
	}

	c.State.Status = types.StatusExited
	c.State.FinishedAt = time.Now().UTC().Format(utils.TimeLayout)
	c.State.Pid = -1
	c.State.ExitCode = exitCode
	c.State.Error = errMsg

	return nil
}

// SetStatusPaused sets a container to be status paused.
func (c *Container) SetStatusPaused() error {
	c.Lock()
	defer c.Unlock()

	if !stateMachine[c.State.Status][types.StatusPaused] {
		return fmt.Errorf("container %s cannot transit from status %v to %v", c.ID, c.State.Status, types.StatusPaused)
	}

	c.State.Status = types.StatusPaused

	return nil
}

// SetStatusRestarting sets a container to be status restarting.
func (c *Container) SetStatusRestarting() error {
	c.Lock()
	defer c.Unlock()

	if !stateMachine[c.State.Status][types.StatusRestarting] {
		return fmt.Errorf("container %s cannot transit from status %v to %v", c.ID, c.State.Status, types.StatusRestarting)
	}

	c.State.Status = types.StatusRestarting
	return nil
}

// SetStatusRemoving sets a container to be status restarting.
func (c *Container) SetStatusRemoving() error {
	c.Lock()
	defer c.Unlock()

	if !stateMachine[c.State.Status][types.StatusRemoving] {
		return fmt.Errorf("container %s cannot transit from status %v to %v", c.ID, c.State.Status, types.StatusRemoving)
	}

	c.State.Status = types.StatusRemoving

	return nil
}
