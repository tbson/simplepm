package restlistutil

import (
	"fmt"
	"src/common/ctype"
	"strings"

	"src/util/numberutil"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const DEFAULT_PAGE_SIZE = 10

type QueryOrder struct {
	Field     string
	Direction string
}

type ListOptions struct {
	Search   string     // Search term
	Filters  ctype.Dict // Filters as key-value pairs
	Order    QueryOrder // Ordering field and direction
	Page     int        // Page number
	Preloads []string
}

type Pages struct {
	Next int `json:"next"` // Next page
	Prev int `json:"prev"` // Previous page
}

type ListRestfulResult[T any] struct {
	Items      []T   `json:"items"`       // Resulting items
	Total      int64 `json:"total"`       // Total records before applying pagination
	Pages      Pages `json:"pages"`       // Pages
	PageSize   int   `json:"page_size"`   // Number of items per page
	TotalPages int   `json:"total_pages"` // Total pages after pagination
}

type ApplyPagingResult struct {
	Query      *gorm.DB
	Pages      Pages
	TotalPages int
}

func GetOptions(c echo.Context, filterableFields []string, orderableFields []string) ListOptions {
	search := c.QueryParam("q")
	page := numberutil.StrToInt(c.QueryParam("page"), 1)
	defaultOrder := QueryOrder{
		Field:     "id",
		Direction: "DESC",
	}

	filters := make(ctype.Dict)
	for _, field := range filterableFields {
		if value := c.QueryParam(field); value != "" {
			filters[field] = value
		}
	}

	order := defaultOrder

	// id = id ASC
	// +id = id ASC
	// -id = id DESC
	orderParam := c.QueryParam("order")
	if orderParam != "" {
		orderDirection := "ASC"
		orderField := orderParam
		if orderParam[0] == '+' {
			orderField = orderParam[1:]
		} else if orderParam[0] == '-' {
			orderField = orderParam[1:]
			orderDirection = "DESC"
		}
		order = QueryOrder{
			Field:     orderField,
			Direction: orderDirection,
		}
	}
	// if the order field is not in the orderable fields, set it to id DESC
	if !contains(orderableFields, order.Field) {
		order = defaultOrder
	}
	return ListOptions{
		Search:   search,
		Filters:  filters,
		Order:    order,
		Page:     page,
		Preloads: []string{},
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ApplySearch(query *gorm.DB, search string, allowFields []string) *gorm.DB {
	if search == "" || len(allowFields) == 0 {
		return query
	}

	searchTerm := "%" + search + "%"
	conditions := make([]string, 0)
	values := make([]interface{}, 0)

	for _, field := range allowFields {
		conditions = append(conditions, fmt.Sprintf("%s LIKE ?", field))
		values = append(values, searchTerm)
	}

	query = query.Where(strings.Join(conditions, " OR "), values...)
	return query
}

func ApplyFilters(query *gorm.DB, filters ctype.Dict) *gorm.DB {
	// Iterate over the filters and apply only the allowed fields
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	return query
}

func ApplyOrder(query *gorm.DB, order QueryOrder) *gorm.DB {
	orderField := order.Field
	orderDirection := order.Direction

	// Apply the ORDER BY clause
	return query.Order(fmt.Sprintf("\"%s\" %s", orderField, orderDirection))
}

func ApplyPaging(query *gorm.DB, page int, total int64) ApplyPagingResult {
	// if there is no previous page, set it to 0
	// if there is no next page, set it to 0
	pageSize := DEFAULT_PAGE_SIZE

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	totalPages := int(total) / pageSize

	if int(total)%pageSize > 0 {
		totalPages++
	}

	pages := Pages{
		Next: 0,
		Prev: 0,
	}

	if page > 1 {
		pages.Prev = page - 1
	}

	if page < totalPages {
		pages.Next = page + 1
	}

	return ApplyPagingResult{
		Query:      query,
		Pages:      pages,
		TotalPages: totalPages,
	}
}

func GetTotalRecords(query *gorm.DB) (int64, error) {
	var totalRecords int64
	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}
	return totalRecords, nil
}
