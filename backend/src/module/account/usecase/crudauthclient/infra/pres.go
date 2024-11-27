package infra

import (
	"src/module/account/schema"
)

type ListOutput = schema.AuthClient
type DetailOutput = schema.AuthClient

func ListPres(items []schema.AuthClient) []ListOutput {
	return items
}

func DetailPres(item schema.AuthClient) DetailOutput {
	return item
}
