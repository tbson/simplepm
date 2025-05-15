package vltd

type CreateVariableInput struct {
	Key         string `json:"key" validate:"required"`
	Value       string `json:"value"`
	Description string `json:"description"`
	DataType    string `json:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

type UpdateVariableInput struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
	DataType    string `json:"data_type" validate:"oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

func (i CreateVariableInput) ToUpdate() UpdateVariableInput {
	return UpdateVariableInput(i)
}
