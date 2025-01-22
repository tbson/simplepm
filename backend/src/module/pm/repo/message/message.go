package message

import (
	"src/client/skyllaclient"
	"src/util/dateutil"
	"time"

	"src/module/pm/schema"

	"github.com/gocql/gocql"
)

type Repo struct {
	client *skyllaclient.Client
}

func New(client *skyllaclient.Client) Repo {
	return Repo{client: client}
}

func (r Repo) List(taskID uint) ([]schema.Message, error) {
	client := skyllaclient.NewClient()
	// defer client.Close()
	rows, err := client.Query("SELECT * FROM event.messages WHERE task_id = ?", taskID)
	if err != nil {
		return nil, err
	}
	result := make([]schema.Message, 0)
	for _, row := range rows {
		id := row["id"].(gocql.UUID).String()
		createdAt := dateutil.TimeToStr(row["created_at"].(time.Time))
		updatedAt := dateutil.TimeToStr(row["updated_at"].(time.Time))
		result = append(result, schema.Message{
			ID:        id,
			UserID:    row["user_id"].(int),
			TaskID:    row["task_id"].(int),
			ProjectID: row["project_id"].(int),
			Content:   row["content"].(string),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return result, nil
}

func (r Repo) Create(message schema.Message) (string, error) {
	client := skyllaclient.NewClient()
	// defer client.Close()
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
	// defer client.Close()
	err := client.Exec("DELETE FROM event.messages WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
