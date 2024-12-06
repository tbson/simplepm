package pm

import "src/common/ctype"

const PLTable = "TABLE"
const PLKanban = "KANBAN"
const PLRoadmap = "ROADMAP"

type option = ctype.SelectOption[string]

var ProjectLayoutOptions = []option{
	{Value: PLTable, Label: "Table"},
	{Value: PLKanban, Label: "Kanban"},
	{Value: PLRoadmap, Label: "Roadmap"},
}
