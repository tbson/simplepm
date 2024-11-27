package localeutil

import (
	"encoding/json"
	"log"
	"src/common/setting"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle      *i18n.Bundle
	bundleOnce  sync.Once
	currentLang string
)

var languageFiles = []string{
	"/code/src/util/localeutil/locales/active.en.json",
	"/code/src/util/localeutil/locales/active.vi.json",
}

func initBundle() {
	bundleOnce.Do(func() {
		bundle = i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
		for _, file := range languageFiles {
			if _, err := bundle.LoadMessageFile(file); err != nil {
				log.Fatalf("Failed to load message file: %v", err)
			}
		}
	})
}

func Init(lang string) *i18n.Localizer {
	initBundle()
	if lang == "" {
		lang = setting.DEFAULT_LANG
	}
	currentLang = lang
	return i18n.NewLocalizer(bundle, currentLang)
}

func Get() *i18n.Localizer {
	if currentLang == "" {
		currentLang = setting.DEFAULT_LANG
	}
	return i18n.NewLocalizer(bundle, currentLang)
}
