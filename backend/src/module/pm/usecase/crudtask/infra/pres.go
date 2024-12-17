package infra

import (
	"src/module/pm/schema"
)

type ListOutput = schema.Task
type DetailOutput = schema.Task

func ListPres(items []schema.Task) []ListOutput {
	return items
}

func DetailPres(item schema.Task) DetailOutput {
	return item
}
