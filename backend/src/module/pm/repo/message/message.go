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
			ID:         id,
			UserID:     uint(row["user_id"].(int)),
			TaskID:     uint(row["task_id"].(int)),
			ProjectID:  uint(row["project_id"].(int)),
			Content:    row["content"].(string),
			UserName:   row["user_name"].(string),
			UserAvatar: row["user_avatar"].(string),
			UserColor:  row["user_color"].(string),
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		})
	}
	return result, nil
}

func (r Repo) Create(message schema.Message) (schema.Message, error) {
	defaultResult := schema.Message{}
	client := skyllaclient.NewClient()
	// defer client.Close()
	id := skyllaclient.GenerateID()
	err := client.Exec(
		"INSERT INTO event.messages (id, user_id, task_id, project_id, content, user_name, user_avatar, user_color, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, toTimestamp(now()), toTimestamp(now()))",
		id, message.UserID, message.TaskID, message.ProjectID, message.Content, message.UserName, message.UserAvatar, message.UserColor,
	)
	if err != nil {
		return defaultResult, err
	}
	result := schema.Message{
		ID:         id.String(),
		UserID:     message.UserID,
		TaskID:     message.TaskID,
		ProjectID:  message.ProjectID,
		Content:    message.Content,
		UserName:   message.UserName,
		UserAvatar: message.UserAvatar,
		UserColor:  message.UserColor,
		CreatedAt:  dateutil.TimeToStr(time.Now()),
		UpdatedAt:  dateutil.TimeToStr(time.Now()),
	}
	return result, nil
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

func (r Repo) CreateAttachment(
	messageID string,
	fileName string,
	fileType string,
	fileURL string,
) (schema.Attachment, error) {
	emptyResult := schema.Attachment{}
	client := skyllaclient.NewClient()
	// defer client.Close()
	id := skyllaclient.GenerateID()
	err := client.Exec(
		"INSERT INTO event.attachments (id, message_id, file_name, file_type, file_url, created_at) VALUES (?, ?, ?, ?, ?, toTimestamp(now()))",
		id, messageID, fileName, fileType, fileURL,
	)
	if err != nil {
		return emptyResult, err
	}
	result := schema.Attachment{
		ID:        id.String(),
		MessageID: messageID,
		FileName:  fileName,
		FileType:  fileType,
		FileURL:   fileURL,
	}
	return result, nil
}

func (r Repo) GetAttachmentMap(
	messages []schema.Message,
) (map[string][]schema.Attachment, error) {
	client := skyllaclient.NewClient()
	// defer client.Close()
	messageIDs := make([]string, 0)
	for _, message := range messages {
		messageIDs = append(messageIDs, message.ID)
	}
	rows, err := client.Query(
		"SELECT * FROM event.attachments WHERE message_id IN ?",
		messageIDs,
	)
	if err != nil {
		return nil, err
	}
	attachments := make(map[string][]schema.Attachment)
	for _, row := range rows {
		messageID := row["message_id"].(gocql.UUID).String()
		attachment := schema.Attachment{
			MessageID: messageID,
			FileName:  row["file_name"].(string),
			FileType:  row["file_type"].(string),
			FileURL:   row["file_url"].(string),
		}
		attachments[messageID] = append(attachments[messageID], attachment)
	}

	return attachments, nil
}
