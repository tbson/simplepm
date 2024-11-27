package infra

type InputData struct {
	Key         string `json:"key" form:"key" validate:"required"`
	Value       string `json:"value" form:"value"`
	Description string `json:"description" form:"description"`
	DataType    string `json:"data_type" form:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}
