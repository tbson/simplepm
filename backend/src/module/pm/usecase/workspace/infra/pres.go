package infra

import (
	"src/module/pm/schema"
)

type ListOutput = schema.Workspace
type DetailOutput = schema.Workspace

func ListPres(items []schema.Workspace) []ListOutput {
	return items
}

func DetailPres(item schema.Workspace) DetailOutput {
	return item
}
