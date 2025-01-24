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
	Status      string       `json:"status"`
	UserInfo    UserInfo     `json:"user_info"`
	Attachments []Attachment `json:"attachments"`
}

func presItem(
	item pmSchema.Message,
	attachmentMap map[string][]pmSchema.Attachment,
	user accountSchema.User,
) ListOutput {
	status := "local"
	if item.UserID != user.ID {
		status = "ai"
	}
	name := strings.TrimSpace(user.FirstName + " " + user.LastName)
	userInfo := UserInfo{
		ID:     int(user.ID),
		Name:   name,
		Avatar: user.Avatar,
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
		Status:      status,
		UserInfo:    userInfo,
		Attachments: attachments,
	}
	return result
}

func ListPres(
	items []pmSchema.Message,
	attachmentMap map[string][]pmSchema.Attachment,
	user accountSchema.User,
) []ListOutput {
	result := make([]ListOutput, 0)
	for _, item := range items {
		result = append(result, presItem(item, attachmentMap, user))
	}
	return result
}
