package infra

import (
	"src/module/document/schema"
)

type ListOutput = schema.DocAttachment
type DetailOutput = schema.DocAttachment

func ListPres(items []schema.DocAttachment) []ListOutput {
	return items
}

func DetailPres(item schema.DocAttachment) DetailOutput {
	return item
}
