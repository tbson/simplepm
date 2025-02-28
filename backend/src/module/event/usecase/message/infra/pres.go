package infra

import (
	accountSchema "src/module/account/schema"
	pmSchema "src/module/event/schema"
	"src/util/dictutil"
)

type UserInfo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Color  string `json:"color"`
}

type Attachment struct {
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileURL  string `json:"file_url"`
	FileSize int    `json:"file_size"`
}

type ListOutput struct {
	ID          string                 `json:"id"`
	Content     string                 `json:"content"`
	Editable    bool                   `json:"editable"`
	Type        string                 `json:"type"`
	GitData     map[string]interface{} `json:"git_data"`
	User        UserInfo               `json:"user"`
	Attachments []Attachment           `json:"attachments"`
}

type ListResult struct {
	Items     []ListOutput `json:"items"`
	PageState []byte       `json:"page_state"`
}

func presItem(
	item pmSchema.Message,
	attachmentMap map[string][]pmSchema.Attachment,
	currentUser accountSchema.User,
) ListOutput {
	editable := false
	if item.UserID != currentUser.ID {
		editable = true
	}
	userInfo := UserInfo{
		ID:     uint(item.UserID),
		Name:   item.UserName,
		Avatar: item.UserAvatar,
		Color:  item.UserColor,
	}
	rawAttachments := attachmentMap[item.ID]
	attachments := make([]Attachment, 0)
	for _, rawAttachment := range rawAttachments {
		attachment := Attachment{
			FileName: rawAttachment.FileName,
			FileType: rawAttachment.FileType,
			FileURL:  rawAttachment.FileURL,
			FileSize: rawAttachment.FileSize,
		}
		attachments = append(attachments, attachment)
	}

	result := ListOutput{
		ID:          item.ID,
		Content:     item.Content,
		Editable:    editable,
		Type:        item.Type,
		GitData:     dictutil.StructToDict(item.GitPush),
		User:        userInfo,
		Attachments: attachments,
	}
	return result
}

func ListPres(
	items []pmSchema.Message,
	pageState []byte,
	attachmentMap map[string][]pmSchema.Attachment,
	currentUser accountSchema.User,
) ListResult {
	result := make([]ListOutput, 0)
	for _, item := range items {
		result = append(result, presItem(item, attachmentMap, currentUser))
	}
	return ListResult{
		Items:     result,
		PageState: pageState,
	}
}
