package infra

import (
	"fmt"
	"net/http"
	"src/util/dbutil"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/numberutil"
	"src/util/vldtutil"

	"src/module/document/repo/docattachment"
	"src/module/document/schema"

	"github.com/labstack/echo/v4"
)

type Schema = schema.DocAttachment

var NewRepo = docattachment.New
var folder = "docattachment"

func Create(c echo.Context) error {
	userID := c.Get("UserID").(uint)
	taskID := numberutil.StrToUint(c.QueryParam("task_id"), 0)
	repo := NewRepo(dbutil.Db(nil))
	structData := InputData{
		UserID: userID, TaskID: taskID,
	}

	folder = fmt.Sprintf("%s/%d", folder, taskID)

	files, err := vldtutil.UploadAndGetMetadata(c, folder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	file := files[0]
	structData.FileName = file.FileName
	structData.FileSize = file.FileSize
	structData.FileType = file.FileType
	structData.FileURL = file.FileURL

	data := dictutil.StructToDict(structData)
	result, err := repo.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusCreated, result)
}
