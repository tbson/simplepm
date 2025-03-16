package pres

import (
	"src/module/config"
	"src/module/config/schema"
)

type outputList struct {
	ID            uint   `json:"id"`
	Key           string `json:"key"`
	Value         string `json:"value"`
	Description   string `json:"description"`
	DataTypeLabel string `json:"data_type"`
}
type outputItem struct {
	ID          uint   `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
	DataType    string `json:"data_type"`
}

func ListPres(items []schema.Variable) []outputList {
	var result []outputList
	for _, item := range items {
		result = append(result, outputList{
			ID:            item.ID,
			Key:           item.Key,
			Value:         item.Value,
			Description:   item.Description,
			DataTypeLabel: config.VariableDataTypeDict.Get(item.DataType),
		})
	}
	return result
}

func DetailPres(item schema.Variable) outputItem {
	return outputItem{
		ID:          item.ID,
		Key:         item.Key,
		Value:       item.Value,
		Description: item.Description,
		DataType:    item.DataType,
	}
}
