package infra

import (
	"src/module/pm/schema"
)

type ListOutput = schema.Feature
type DetailOutput = schema.Feature

func ListPres(items []schema.Feature) []ListOutput {
	return items
}

func DetailPres(item schema.Feature) DetailOutput {
	return item
}
