package infra

import (
	"src/module/pm/schema"
)

type ListOutput = schema.Project
type DetailOutput = schema.Project

func ListPres(items []schema.Project) []ListOutput {
	return items
}

func DetailPres(item schema.Project) DetailOutput {
	return item
}
