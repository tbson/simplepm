package infra

type InputData struct {
	Uid         string `json:"uid" validate:"required"`
	Description string `json:"description"`
	Secret      string `json:"secret" validate:"required"`
	Partition   string `json:"partition" validate:"required"`
	Default     bool   `json:"default"`
}
