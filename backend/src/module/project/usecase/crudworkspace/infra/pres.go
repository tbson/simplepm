package infra

import (
	"src/module/project/schema"
)

type ListOutput = schema.Workspace
type DetailOutput = schema.Workspace

func ListPres(items []schema.Workspace) []ListOutput {
	return items
}

func DetailPres(item schema.Workspace) DetailOutput {
	return item
}
