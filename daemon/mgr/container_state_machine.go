package mgr

import "github.com/alibaba/pouch/apis/types"

var stateMachine = map[types.Status]map[types.Status]bool{
	types.StatusCreated:    createdTransitionRule,
	types.StatusRunning:    runningTransitionRule,
	types.StatusPaused:     pausedTransitionRule,
	types.StatusRestarting: restartingTransitionRule,
	types.StatusRemoving:   removingTransitionRule,
	types.StatusExited:     exitedTransitionRule,
	types.StatusStopped:    stoppedTransitionRule,
	types.StatusDead:       deadTransitionRule,
}

// createdStateMachine is ...
var createdTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}

var runningTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}

var pausedTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}

var restartingTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}

var removingTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}

var exitedTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}

var stoppedTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}

var deadTransitionRule = map[types.Status]bool{
	types.StatusCreated:    false,
	types.StatusRunning:    false,
	types.StatusPaused:     false,
	types.StatusRestarting: false,
	types.StatusRemoving:   false,
	types.StatusExited:     false,
	types.StatusDead:       false,
	types.StatusStopped:    false,
}
