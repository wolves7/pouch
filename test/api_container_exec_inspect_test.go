package main

import (
	"fmt"
	"time"

	"github.com/alibaba/pouch/test/environment"
	"github.com/alibaba/pouch/test/request"

	"github.com/go-check/check"
)

// APIContainerExecInspectSuite is the test suite for container exec inspect API.
type APIContainerExecInspectSuite struct{}

func init() {
	check.Suite(&APIContainerExecInspectSuite{})
}

// SetUpTest does common setup in the beginning of each test.
func (suite *APIContainerExecInspectSuite) SetUpTest(c *check.C) {
	SkipIfFalse(c, environment.IsLinux)

	PullImage(c, busyboxImage)
}

// TestContainerCreateExecOk tests execing containers is OK.
func (suite *APIContainerExecInspectSuite) TestContainerExecInspectOk(c *check.C) {
	cname := "TestContainerExecInspectOk"

	CreateBusyboxContainerOk(c, cname)
	defer DelContainerForceMultyTime(c, cname)

	StartContainerOk(c, cname)

	// create an exec and get the execID
	// and inspect the exec before exec start

	execid := CreateExecCmdOk(c, cname, "sleep", "9")
	{
		execInspectResp := InspectExecOk(c, execid)
		c.Assert(execInspectResp.Running, check.Equals, false)
		c.Assert(execInspectResp.ExitCode, check.Equals, int64(0))
	}

	// set the detach to be true
	// and start the exec
	{
		obj := map[string]interface{}{
			"Detach": true,
		}
		body := request.WithJSONBody(obj)
		resp, err := request.Post(fmt.Sprintf("/exec/%s/start", execid), body)
		c.Assert(err, check.IsNil)
		CheckRespStatus(c, resp, 200)
	}

	// inspect the exec after exec start
	{
		execInspectResp := InspectExecOk(c, execid)
		c.Assert(execInspectResp.Running, check.Equals, true)
		c.Assert(execInspectResp.ExitCode, check.Equals, int64(0))
	}

	// sleep 10s to wait the process exit
	{
		time.Sleep(10 * time.Second)
		execInspectResp := InspectExecOk(c, execid)
		c.Assert(execInspectResp.Running, check.Equals, false)
		c.Assert(execInspectResp.ExitCode, check.Equals, int64(0))
	}
}
