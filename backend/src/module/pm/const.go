package pm

import "src/common/ctype"

const PROJECT_LAYOUT_TABLE = "TABLE"
const PROJECT_LAYOUT_KANBAN = "KANBAN"
const PROJECT_LAYOUT_ROADMAP = "ROADMAP"

const PROJECT_STATUS_ACTIVE = "ACTIVE"
const PROJECT_STATUS_ARCHIVE = "ARCHIVE"

type option = ctype.SelectOption[string]

var ProjectLayoutOptions = []option{
	{Value: PROJECT_LAYOUT_TABLE, Label: "Table"},
	{Value: PROJECT_LAYOUT_KANBAN, Label: "Kanban"},
	{Value: PROJECT_LAYOUT_ROADMAP, Label: "Roadmap"},
}

var ProjectStatusOptions = []option{
	{Value: PROJECT_STATUS_ACTIVE, Label: "Active"},
	{Value: PROJECT_STATUS_ARCHIVE, Label: "Archive"},
}
