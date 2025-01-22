package infra

import "src/module/pm/schema"

type ListOutput struct {
	ID      string `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

func presItem(item schema.Message) ListOutput {
	result := ListOutput{
		ID:      item.ID,
		UserID:  item.UserID,
		Content: item.Content,
	}
	return result
}

func ListPres(items []schema.Message) []ListOutput {
	result := make([]ListOutput, 0)
	for _, item := range items {
		result = append(result, presItem(item))
	}
	return result
}

func DetailPres(item schema.Message) ListOutput {
	return presItem(item)
}
