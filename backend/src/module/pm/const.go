package pm

import "src/common/ctype"

const PROJECT_LAYOUT_TABLE = "TABLE"
const PROJECT_LAYOUT_KANBAN = "KANBAN"
const PROJECT_LAYOUT_ROADMAP = "ROADMAP"

const PROJECT_STATUS_ACTIVE = "ACTIVE"
const PROJECT_STATUS_ARCHIVE = "ARCHIVE"

const TASK_FIELD_TYPE_TEXT = "TEXT"
const TASK_FIELD_TYPE_NUMBER = "NUMBER"
const TASK_FIELD_TYPE_DATE = "DATE"
const TASK_FIELD_TYPE_SELECT = "SELECT"
const TASK_FIELD_TYPE_MULTIPLE_SELECT = "MULTIPLE_SELECT"

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

var TaskFieldTypeOptions = []option{
	{Value: TASK_FIELD_TYPE_TEXT, Label: "Text"},
	{Value: TASK_FIELD_TYPE_NUMBER, Label: "Number"},
	{Value: TASK_FIELD_TYPE_DATE, Label: "Date"},
	{Value: TASK_FIELD_TYPE_SELECT, Label: "Select"},
	{Value: TASK_FIELD_TYPE_MULTIPLE_SELECT, Label: "Multiple Select"},
}
