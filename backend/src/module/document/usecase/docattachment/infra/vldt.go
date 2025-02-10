package infra

type InputData struct {
	UserID   uint   `json:"user_id" form:"user_id" validate:"required"`
	TaskID   uint   `json:"task_id" form:"task_id" validate:"required"`
	FileName string `json:"file_name" form:"file_name" validate:"required"`
	FileType string `json:"file_type" form:"file_type" validate:"required"`
	FileSize int    `json:"file_size" form:"file_size" validate:"required"`
	FileURL  string `json:"file_url" form:"file_url" validate:"required"`
}
