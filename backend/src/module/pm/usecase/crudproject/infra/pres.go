package infra

import (
	"src/common/ctype"
	"src/module/pm/repo/project"
	"src/module/pm/schema"
	"src/util/dbutil"
)

type ListOutput struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Layout         string `json:"layout"`
	Status         string `json:"status"`
	WorkspaceID    *uint  `json:"workspace_id"`
	WorkspaceLabel string `json:"workspace_label"`
	Order          int    `json:"order"`
}

type DetailOutput = schema.Project

func listPresItem(item schema.Project) ListOutput {
	workspaceLabel := ""
	if item.Workspace != nil {
		workspaceLabel = item.Workspace.Title
	}

	return ListOutput{
		ID:             item.ID,
		Title:          item.Title,
		Description:    item.Description,
		Layout:         item.Layout,
		Status:         item.Status,
		WorkspaceID:    item.WorkspaceID,
		WorkspaceLabel: workspaceLabel,
		Order:          item.Order,
	}
}

func ListPres(items []schema.Project) []ListOutput {
	var result []ListOutput
	for _, item := range items {
		result = append(result, listPresItem(item))
	}
	return result
}

func DetailPres(item schema.Project) DetailOutput {
	return item
}

func MutatePres(item schema.Project) ListOutput {
	projectRepo := project.New(dbutil.Db())
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"id": item.ID,
		},
		Preloads: []string{"Workspace"},
	}
	presItem, _ := projectRepo.Retrieve(queryOptions)
	return listPresItem(*presItem)
}
