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
	Joins    []string
	Order    string
}

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

type socketData struct {
	ID        string `json:"id"`
	UserID    uint   `json:"user_id"`
	TaskID    uint   `json:"task_id"`
	ProjectID uint   `json:"project_id"`
	Content   string `json:"content"`
}

type SocketMessage struct {
	Channel string     `json:"channel"`
	Data    socketData `json:"data"`
}
