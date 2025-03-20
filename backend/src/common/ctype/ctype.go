package ctype

type Dict map[string]interface{}

type StrDict map[string]string

func (v StrDict) Get(key string) string {
	if val, ok := v[key]; ok {
		return val
	}
	return ""
}

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
	Joins    []string
	Order    string
}

type Fields []string

type SimpleSelectOption[T any] struct {
	Value T      `json:"value"`
	Label string `json:"label"`
}

type SelectOption[T any] struct {
	Value       T                       `json:"value"`
	Label       string                  `json:"label"`
	Description string                  `json:"description"`
	Group       string                  `json:"group"`
	Options     []SimpleSelectOption[T] `json:"options"`
}

type EmailBody struct {
	HmtlPath string
	Data     Dict
}
