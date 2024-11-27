package paging

import (
	"src/util/restlistutil"

	"gorm.io/gorm"
)

type Repo[S any, P any] struct {
	client *gorm.DB
	pres   func([]S) []P
}

func New[S any, P any](client *gorm.DB, pres func([]S) []P) Repo[S, P] {
	return Repo[S, P]{
		client: client,
		pres:   pres,
	}
}

func (r Repo[S, P]) Paging(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[P], error) {
	db := r.client
	preloads := options.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}
	pageSize := restlistutil.DEFAULT_PAGE_SIZE
	var items []P
	emptyResult := restlistutil.ListRestfulResult[P]{
		Items:      items,
		Total:      0,
		PageSize:   pageSize,
		TotalPages: 0,
		Pages: restlistutil.Pages{
			Next: 0,
			Prev: 0,
		},
	}
	query := db.Model(new(*S))

	// Apply preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	// Apply search logic
	query = restlistutil.ApplySearch(query, options.Search, searchableFields)

	// Apply filters
	query = restlistutil.ApplyFilters(query, options.Filters)

	// Apply order
	query = restlistutil.ApplyOrder(query, options.Order)

	// Count total records before pagination
	total, err := restlistutil.GetTotalRecords(query)
	if err != nil {
		return emptyResult, err
	}

	// Apply paging
	pagingREsult := restlistutil.ApplyPaging(query, options.Page, total)
	query = pagingREsult.Query
	pages := pagingREsult.Pages
	totalPages := pagingREsult.TotalPages

	// Fetch the results
	schemaItems := []S{}
	result := query.Find(&schemaItems)
	if result.Error != nil {
		return emptyResult, result.Error
	}
	return restlistutil.ListRestfulResult[P]{
		Items:      r.pres(schemaItems),
		Total:      total,
		Pages:      pages,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}
