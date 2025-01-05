package pm

import "src/common/ctype"

const PROJECT_LAYOUT_TABLE = "TABLE"
const PROJECT_LAYOUT_KANBAN = "KANBAN"
const PROJECT_LAYOUT_ROADMAP = "ROADMAP"

const PROJECT_STATUS_ACTIVE = "ACTIVE"
const PROJECT_STATUS_FINISHED = "FINISHED"
const PROJECT_STATUS_ARCHIVED = "ARCHIVED"

const PROJECT_REPO_TYPE_GITHUB = "GITHUB"
const PROJECT_REPO_TYPE_GITLAB = "GITLAB"

const TASK_FIELD_TYPE_TEXT = "TEXT"
const TASK_FIELD_TYPE_NUMBER = "NUMBER"
const TASK_FIELD_TYPE_DATE = "DATE"
const TASK_FIELD_TYPE_SELECT = "SELECT"
const TASK_FIELD_TYPE_MULTIPLE_SELECT = "MULTIPLE_SELECT"

const TASK_FIELD_OPTION_COLOR_GRAY = "GRAY"
const TASK_FIELD_OPTION_COLOR_BLUE = "BLUE"
const TASK_FIELD_OPTION_COLOR_GREEN = "GREEN"
const TASK_FIELD_OPTION_COLOR_YELLOW = "YELLOW"
const TASK_FIELD_OPTION_COLOR_RED = "RED"
const TASK_FIELD_OPTION_COLOR_ORANGE = "ORANGE"
const TASK_FIELD_OPTION_COLOR_INDIGO = "INDIGO"
const TASK_FIELD_OPTION_COLOR_VIOLET = "VIOLET"

type option = ctype.SelectOption[string]

var ProjectLayoutOptions = []option{
	{Value: PROJECT_LAYOUT_TABLE, Label: "Table"},
	{Value: PROJECT_LAYOUT_KANBAN, Label: "Kanban"},
	{Value: PROJECT_LAYOUT_ROADMAP, Label: "Roadmap"},
}

var ProjectStatusOptions = []option{
	{Value: PROJECT_STATUS_ACTIVE, Label: "Active"},
	{Value: PROJECT_STATUS_FINISHED, Label: "Finished"},
	{Value: PROJECT_STATUS_ARCHIVED, Label: "Archived"},
}

var ProjectRepoTypeOptions = []option{
	{Value: PROJECT_REPO_TYPE_GITHUB, Label: "GitHub"},
	{Value: PROJECT_REPO_TYPE_GITLAB, Label: "GitLab"},
}

var TaskFieldTypeOptions = []option{
	{Value: TASK_FIELD_TYPE_TEXT, Label: "Text"},
	{Value: TASK_FIELD_TYPE_NUMBER, Label: "Number"},
	{Value: TASK_FIELD_TYPE_DATE, Label: "Date"},
	{Value: TASK_FIELD_TYPE_SELECT, Label: "Select"},
	{Value: TASK_FIELD_TYPE_MULTIPLE_SELECT, Label: "Multiple Select"},
}
