package infra

import (
	"src/module/account/schema"
)

type ListOutput = schema.Role
type DetailOutput = schema.Role

func ListPres(items []schema.Role) []ListOutput {
	return items
}

func DetailPres(item schema.Role) DetailOutput {
	return item
}
