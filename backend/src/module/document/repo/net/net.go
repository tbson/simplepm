package net

import (
	"io"
	"net/http"
	"src/util/errutil"
	"src/util/i18nmsg"

	"golang.org/x/net/html"
)

type HTMLMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	SiteName    string `json:"site_name"`
}

type Repo struct {
	client *http.Client
}

func New(client *http.Client) Repo {
	return Repo{
		client: client,
	}
}

func (r Repo) GetHTMLMeta(link string) (HTMLMeta, error) {
	result := HTMLMeta{}
	resp, err := r.client.Get(link)

	errObj := errutil.New(i18nmsg.CanNotParseLinkMetaData)
	if err != nil || resp.StatusCode != http.StatusOK {
		return result, errObj
	}
	defer resp.Body.Close()

	result = extract(resp.Body)

	if result.Title == "" {
		return result, errObj
	}

	return result, nil
}

func extract(resp io.Reader) HTMLMeta {
	z := html.NewTokenizer(resp)

	titleFound := false

	hm := HTMLMeta{}

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return hm
			}
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				desc, ok := extractMetaProperty(t, "description")
				if ok {
					hm.Description = desc
				}

				ogTitle, ok := extractMetaProperty(t, "og:title")
				if ok {
					hm.Title = ogTitle
				}

				ogDesc, ok := extractMetaProperty(t, "og:description")
				if ok {
					hm.Description = ogDesc
				}

				ogImage, ok := extractMetaProperty(t, "og:image")
				if ok {
					hm.Image = ogImage
				}

				ogSiteName, ok := extractMetaProperty(t, "og:site_name")
				if ok {
					hm.SiteName = ogSiteName
				}
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				hm.Title = t.Data
				titleFound = false
			}
		}
	}
}

func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if attr.Key == "property" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
		}
	}

	return
}
