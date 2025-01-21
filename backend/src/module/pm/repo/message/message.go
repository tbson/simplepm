package message

import (
	"src/client/skyllaclient"

	"github.com/gocql/gocql"
)

type Schema struct {
	ID        string `json:"id"`
	UserID    int    `json:"user_id"`
	TaskID    int    `json:"task_id"`
	ProjectID int    `json:"project_id"`
	Content   string `json:"content"`
}

type Repo struct {
	client *skyllaclient.Client
}

func New(client *skyllaclient.Client) Repo {
	return Repo{client: client}
}

func (r Repo) List(taskID uint) ([]Schema, error) {
	client := skyllaclient.NewClient()
	defer client.Close()
	rows, err := client.Query("SELECT * FROM event.messages WHERE task_id = ?", taskID)
	if err != nil {
		return nil, err
	}
	result := make([]Schema, 0)
	for _, row := range rows {
		id := row["id"].(gocql.UUID).String()
		result = append(result, Schema{
			ID:        id,
			UserID:    row["user_id"].(int),
			TaskID:    row["task_id"].(int),
			ProjectID: row["project_id"].(int),
			Content:   row["content"].(string),
		})
	}
	return result, nil
}

func (r Repo) Create(message Schema) (string, error) {
	client := skyllaclient.NewClient()
	defer client.Close()
	id := skyllaclient.GenerateID()
	err := client.Exec(
		"INSERT INTO event.messages (id, user_id, task_id, project_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, toTimestamp(now()), toTimestamp(now()))",
		id, message.UserID, message.TaskID, message.ProjectID, message.Content,
	)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (r Repo) Delete(id string) error {
	client := skyllaclient.NewClient()
	defer client.Close()
	err := client.Exec("DELETE FROM event.messages WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
