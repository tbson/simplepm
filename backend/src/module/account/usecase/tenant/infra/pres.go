package infra

import (
	"src/module/account/schema"
)

type ListOutput = schema.Tenant
type DetailOutput = schema.Tenant

func ListPres(items []schema.Tenant) []ListOutput {
	return items
}

func DetailPres(item schema.Tenant) DetailOutput {
	return item
}
