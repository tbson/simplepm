package templateutil

import (
	"bytes"
	"path/filepath"
	"src/common/ctype"
	"src/util/errutil"
	"src/util/i18nmsg"
	"text/template"
)

func GetHtmlString(templatePath string, data ctype.Dict) (string, error) {
	path := filepath.Join("/code/public", templatePath)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", errutil.NewWithArgs(
			i18nmsg.FailedToParseTemplate,
			ctype.Dict{
				"Value": templatePath,
			},
		)
	}

	var templateBody bytes.Buffer
	if err := tmpl.Execute(&templateBody, data); err != nil {
		return "", errutil.NewWithArgs(
			i18nmsg.FailedToExecuteTemplate,
			ctype.Dict{
				"Value": templatePath,
			},
		)
	}

	return templateBody.String(), nil
}
