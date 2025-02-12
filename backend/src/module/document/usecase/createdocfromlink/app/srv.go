package app

import (
	"src/module/document"
	"src/module/document/schema"
	"src/util/dictutil"
)

type Service struct {
	docRepo DocRepo
	netRepo NetRepo
}

func New(docRepo DocRepo, netRepo NetRepo) Service {
	return Service{docRepo, netRepo}
}

func (srv Service) Create(data InputData) (*schema.Doc, error) {
	link := data.Link
	linkData, err := srv.netRepo.GetHTMLMeta(link)
	if err != nil {
		return nil, err
	}
	data.Title = linkData.Title
	data.Description = linkData.Description
	data.Type = document.DOC_TYPE_LINK

	dictData := dictutil.StructToDict(data)
	tenant, err := srv.docRepo.Create(dictData)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}
