package templateutil

import (
	"bytes"
	"path/filepath"
	"src/common/ctype"
	"src/util/errutil"
	"src/util/localeutil"
	"text/template"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetHtmlString(templatePath string, data ctype.Dict) (string, error) {
	localizer := localeutil.Get()
	path := filepath.Join("/code/public", templatePath)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToParseTemplate,
			TemplateData: ctype.Dict{
				"Value": templatePath,
			},
		})
		return "", errutil.New("", []string{msg})
	}

	var templateBody bytes.Buffer
	if err := tmpl.Execute(&templateBody, data); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToExecuteTemplate,
			TemplateData: ctype.Dict{
				"Value": templatePath,
			},
		})
		return "", errutil.New("", []string{msg})
	}

	return templateBody.String(), nil
}
