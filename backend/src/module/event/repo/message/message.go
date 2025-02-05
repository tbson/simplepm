package message

import (
	"fmt"
	"src/client/scyllaclient"
	"src/util/dateutil"
	"time"

	"src/common/setting"
	"src/module/event/schema"

	"github.com/gocql/gocql"
)

type Repo struct {
	client *scyllaclient.Client
}

func New(client *scyllaclient.Client) Repo {
	return Repo{client: client}
}

func (r Repo) List(taskID uint, pageState []byte) ([]schema.Message, []byte, error) {
	pageSize := setting.MSG_PAGE_SIZE
	client := scyllaclient.NewClient()

	// Use QueryWithPaging to build the query
	q := client.QueryWithPaging(
		"SELECT * FROM event.messages WHERE task_id = ? ORDER BY id DESC",
		pageSize,
		pageState,
		taskID,
	)
	iter := q.Iter()

	var messages []schema.Message
	rowData := map[string]interface{}{}

	// Manually iterate over rows up to pageSize
	count := 0
	for count < pageSize && iter.MapScan(rowData) {
		msg := schema.Message{
			ID:         rowData["id"].(gocql.UUID).String(),
			UserID:     uint(rowData["user_id"].(int)),
			TaskID:     uint(rowData["task_id"].(int)),
			ProjectID:  uint(rowData["project_id"].(int)),
			Content:    rowData["content"].(string),
			UserName:   rowData["user_name"].(string),
			UserAvatar: rowData["user_avatar"].(string),
			UserColor:  rowData["user_color"].(string),
			CreatedAt:  dateutil.TimeToStr(rowData["created_at"].(time.Time)),
			UpdatedAt:  dateutil.TimeToStr(rowData["updated_at"].(time.Time)),
		}
		messages = append(messages, msg)
		count++
		// Reinitialize rowData for the next row
		rowData = map[string]interface{}{}
	}

	// Get the paging state for the next page.
	nextPageState := iter.PageState()

	if err := iter.Close(); err != nil {
		return nil, nil, err
	}
	// check messages length to return empty page state
	if len(messages) < pageSize {
		nextPageState = nil
	}
	// reverse the messages
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages, nextPageState, nil
}

func (r Repo) Create(message schema.Message) (schema.Message, error) {
	defaultResult := schema.Message{}
	client := scyllaclient.NewClient()
	// defer client.Close()
	id := scyllaclient.GenerateID()
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

func (r Repo) Update(id string, taskId uint, message schema.Message) (schema.Message, error) {
	defaultResult := schema.Message{}
	client := scyllaclient.NewClient()
	// defer client.Close()
	err := client.Exec(
		"UPDATE event.messages SET content = ?, updated_at = toTimestamp(now()) WHERE id = ? AND task_id = ?",
		message.Content, id, taskId,
	)
	if err != nil {
		return defaultResult, err
	}
	result := schema.Message{
		ID:        id,
		Content:   message.Content,
		UpdatedAt: dateutil.TimeToStr(time.Now()),
	}
	return result, nil
}

func (r Repo) Delete(id string, task_id uint) error {
	client := scyllaclient.NewClient()
	// defer client.Close()
	err := client.Exec("DELETE FROM event.messages WHERE id = ? AND task_id = ?", id, task_id)
	if err != nil {
		fmt.Println("error deleting message", err)
		return err
	}
	return nil
}

func (r Repo) CreateAttachment(
	messageID string,
	fileName string,
	fileType string,
	fileURL string,
	fileSize int,
) (schema.Attachment, error) {
	emptyResult := schema.Attachment{}
	client := scyllaclient.NewClient()
	// defer client.Close()
	id := scyllaclient.GenerateID()
	err := client.Exec(
		"INSERT INTO event.attachments (id, message_id, file_name, file_type, file_url, file_size, created_at) VALUES (?, ?, ?, ?, ?, ?, toTimestamp(now()))",
		id, messageID, fileName, fileType, fileURL, fileSize,
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
		FileSize:  fileSize,
	}
	return result, nil
}

func (r Repo) GetAttachmentMap(
	messages []schema.Message,
) (map[string][]schema.Attachment, error) {
	client := scyllaclient.NewClient()
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
			FileSize:  row["file_size"].(int),
		}
		attachments[messageID] = append(attachments[messageID], attachment)
	}

	return attachments, nil
}
