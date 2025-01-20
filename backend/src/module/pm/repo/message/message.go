package message

import (
	"src/client/skyllaclient"
)

type Schema struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	TaskID  uint   `json:"task_id"`
	Content string `json:"content"`
}

type Repo struct {
	client *skyllaclient.Client
}

func New(client *skyllaclient.Client) Repo {
	return Repo{client: client}
}

func (r Repo) List() ([]Schema, error) {
	client := skyllaclient.NewClient()
	defer client.Close()
	project_id := 1
	rows, err := client.Query("SELECT * FROM event.messages WHERE project_id = ?", project_id)
	if err != nil {
		return nil, err
	}
	result := make([]Schema, 0)
	for _, row := range rows {
		result = append(result, Schema{
			ID:      row["id"].(uint),
			UserID:  row["user_id"].(uint),
			TaskID:  row["task_id"].(uint),
			Content: row["content"].(string),
		})
	}
	return result, nil
}
