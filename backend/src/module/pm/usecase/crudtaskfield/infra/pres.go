package infra

import (
	"src/module/pm/schema"
)

type ListOutput = schema.TaskField
type DetailOutput = schema.TaskField

func ListPres(items []schema.TaskField) []ListOutput {
	return items
}

func DetailPres(item schema.TaskField) DetailOutput {
	return item
}
