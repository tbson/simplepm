package infra

import (
	"src/module/config/schema"
)

type ListOutput = schema.Variable
type DetailOutput = schema.Variable

func ListPres(items []schema.Variable) []ListOutput {
	return items
}

func DetailPres(item schema.Variable) DetailOutput {
	return item
}
