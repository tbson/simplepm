package infra

import (
	"fmt"
	"src/module/event/repo/change"
	"src/util/dbutil"
	"src/util/dictutil"
	"src/util/vldtutil"
)

func LogCreateTask(msg []byte) {
	changeRepo := change.New(dbutil.Db(nil))
	structData, err := vldtutil.BytesToStruct(msg, InputData{})
	if err != nil {
		return
	}

	data := dictutil.StructToDict(structData)

	_, err = changeRepo.Create(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
