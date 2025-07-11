package pres

import (
	"src/common/ctype"
	"src/module/config"
	"src/module/config/schema"
	"src/util/restlistutil"

	"github.com/google/uuid"
)

type outputListItem struct {
	ID            uuid.UUID `json:"id"`
	Uid           string    `json:"uid"`
	Value         string    `json:"value"`
	Description   string    `json:"description"`
	DataTypeLabel string    `json:"data_type"`
}
type outputItem struct {
	ID          uuid.UUID `json:"id"`
	Uid         string    `json:"uid"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
	DataType    string    `json:"data_type"`
}

type PageResult restlistutil.ListRestfulResult[outputListItem]

func PagePres(result restlistutil.ListRestfulResult[schema.Variable]) PageResult {
	return PageResult{
		Items:      ListPres(result.Items),
		Total:      result.Total,
		Pages:      result.Pages,
		PageSize:   result.PageSize,
		TotalPages: result.TotalPages,
	}
}

type ListResult []outputListItem

func ListPres(items []schema.Variable) ListResult {
	var result ListResult
	for _, item := range items {
		result = append(result, outputListItem{
			ID:            item.ID,
			Uid:           item.Uid,
			Value:         item.Value,
			Description:   item.Description,
			DataTypeLabel: config.VariableDataTypeDict.Get(item.DataType),
		})
	}
	return result
}

type DetailResult outputItem

func DetailPres(item schema.Variable) DetailResult {
	return DetailResult{
		ID:          item.ID,
		Uid:         item.Uid,
		Value:       item.Value,
		Description: item.Description,
		DataType:    item.DataType,
	}
}

type DeleteResult struct {
	IDs []string `json:"ids"`
}

func DeletePres(ids []string) DeleteResult {
	return DeleteResult{IDs: ids}
}

type OptionResult struct {
	DataType []ctype.SelectOption[string] `json:"data_type"`
}

func OptionPres(dataType []ctype.SelectOption[string]) OptionResult {
	return OptionResult{
		DataType: dataType,
	}
}
