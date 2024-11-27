package repoutil

import (
	"src/util/restlistutil"

	"gorm.io/gorm"
)

type Repo[T any] struct {
	schema T
	client *gorm.DB
}

func (r Repo[T]) New(client *gorm.DB) Repo[T] {
	return Repo[T]{
		schema: r.schema,
		client: client,
	}
}

func (r Repo[T]) ListPaging(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[T], error) {
	db := r.client
	preloads := options.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}
	pageSize := restlistutil.DEFAULT_PAGE_SIZE
	var items []T
	emptyResult := restlistutil.ListRestfulResult[T]{
		Items:      items,
		Total:      0,
		PageSize:   pageSize,
		TotalPages: 0,
		Pages: restlistutil.Pages{
			Next: 0,
			Prev: 0,
		},
	}
	query := db.Model(new(*T))

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
	result := query.Find(&items)
	if result.Error != nil {
		return emptyResult, result.Error
	}
	return restlistutil.ListRestfulResult[T]{
		Items:      items,
		Total:      total,
		Pages:      pages,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}
