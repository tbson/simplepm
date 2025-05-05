package queueworker

import (
	"log"

	logcreatetask "src/module/event/usecase/logtask/infra"
	"src/util/vldtutil"
)

const TEST_MESSAGE = "TEST_MESSAGE"
const LOG_CREATE_TASK = "LOG_CREATE_TASK"
const LOG_EDIT_TASK = "LOG_EDIT_TASK"
const LOG_DELETE_TASK = "LOG_DELETE_TASK"

type TestMessage struct {
	UserID uint   `json:"user_id"`
	Action string `json:"action"`
}

// Handler for audit log messages.
func handleTestMessage(msg []byte) {
	data, err := vldtutil.BytesToStruct(msg, &TestMessage{})
	if err != nil {
		return
	}
	log.Printf("[AuditLog] User ID: %d, Action: %s", data.UserID, data.Action)
}

var Handlers map[string]func([]byte) = map[string]func([]byte){
	TEST_MESSAGE:    handleTestMessage,
	LOG_CREATE_TASK: logcreatetask.LogCreateTask,
	LOG_EDIT_TASK:   logcreatetask.LogEditTask,
	LOG_DELETE_TASK: logcreatetask.LogDeleteTask,
}
