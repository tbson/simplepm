package infra

import (
	accountSchema "src/module/account/schema"
	pmSchema "src/module/pm/schema"
	"strings"
)

type UserInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Attachment struct {
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileURL  string `json:"file_url"`
}

type ListOutput struct {
	ID          string       `json:"id"`
	Content     string       `json:"content"`
	Editable    bool         `json:"editable"`
	UserInfo    UserInfo     `json:"user_info"`
	Attachments []Attachment `json:"attachments"`
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
	name := strings.TrimSpace(currentUser.FirstName + " " + currentUser.LastName)
	userInfo := UserInfo{
		ID:     int(currentUser.ID),
		Name:   name,
		Avatar: currentUser.Avatar,
	}
	rawAttachments := attachmentMap[item.ID]
	attachments := make([]Attachment, 0)
	for _, rawAttachment := range rawAttachments {
		attachment := Attachment{
			FileName: rawAttachment.FileName,
			FileType: rawAttachment.FileType,
			FileURL:  rawAttachment.FileURL,
		}
		attachments = append(attachments, attachment)
	}
	result := ListOutput{
		ID:          item.ID,
		Content:     item.Content,
		Editable:    editable,
		UserInfo:    userInfo,
		Attachments: attachments,
	}
	return result
}

func ListPres(
	items []pmSchema.Message,
	attachmentMap map[string][]pmSchema.Attachment,
	currentUser accountSchema.User,
) []ListOutput {
	result := make([]ListOutput, 0)
	for _, item := range items {
		result = append(result, presItem(item, attachmentMap, currentUser))
	}
	return result
}
