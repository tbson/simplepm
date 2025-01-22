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

type ListOutput struct {
	ID       string   `json:"id"`
	Message  string   `json:"message"`
	Status   string   `json:"status"`
	UserInfo UserInfo `json:"user_info"`
}

func presItem(item pmSchema.Message, user accountSchema.User) ListOutput {
	status := "local"
	if item.UserID != item.UserID {
		status = "ai"
	}
	name := strings.TrimSpace(user.FirstName + " " + user.LastName)
	userInfo := UserInfo{
		ID:     int(user.ID),
		Name:   name,
		Avatar: user.Avatar,
	}
	result := ListOutput{
		ID:       item.ID,
		Message:  item.Content,
		Status:   status,
		UserInfo: userInfo,
	}
	return result
}

func ListPres(items []pmSchema.Message, user accountSchema.User) []ListOutput {
	result := make([]ListOutput, 0)
	for _, item := range items {
		result = append(result, presItem(item, user))
	}
	return result
}
