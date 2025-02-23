package infra

import (
	"fmt"
	"src/util/vldtutil"
)

func LogCreateTask(msg []byte) {
	fmt.Println("Create task")
	data, err := vldtutil.BytesToStruct(msg, &InputData{})
	if err != nil {
		return
	}
	fmt.Println(data)
}
