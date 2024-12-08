package app

import "fmt"

var (
	BaseRoute            = "/"
	StatusRoute          = "/status"
	TaskListRoute        = "/tasks"
	TaskRoute            = fmt.Sprintf("%s/:%s", TaskListRoute, ParamId)
	EditTaskNameRoute    = fmt.Sprintf("%s/edit-name", TaskRoute)
	DisplayTaskNameRoute = fmt.Sprintf("%s/display-name", TaskRoute)
)

const (
	ParamLocked   = "locked"
	ParamId       = "id"
	ParamTaskName = "task-name"
	ParamChecked  = "checked"
)
