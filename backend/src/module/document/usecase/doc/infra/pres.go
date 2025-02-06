package infra

import (
	"src/module/document/schema"
)

type ListOutput = schema.Doc
type DetailOutput = schema.Doc

func ListPres(items []schema.Doc) []ListOutput {
	return items
}

func DetailPres(item schema.Doc) DetailOutput {
	return item
}
