package ctype

type Dict map[string]interface{}

type Pem struct {
	ProfileTypes []string
	Title        string
	Module       string
	Action       string
}

type PemMap map[string]Pem

type QueryOptions struct {
	Filters  Dict
	Preloads []string
	Order    string
}

type SelectOption[T any] struct {
	Value       T      `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Group       string `json:"group"`
}
