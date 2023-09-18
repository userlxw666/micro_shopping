package utils

import (
	"github.com/gin-gonic/gin"
	"micro_shopping/idl/pb"
	"strconv"
)

var (
	DefaultPageSize = 100
	MaxPageSize     = 1000
	PageVar         = "page"
	PageSizeVar     = "pageSize"
)

type Pages struct {
	Page      int64
	PageSize  int64
	Total     int64
	PageCount int64
}

func New(page, pageSize, total int) *Pages {
	if pageSize < 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}
	}
	if page <= 0 {
		page = 1
	}

	return &Pages{
		Page:      int64(page),
		PageSize:  int64(pageSize),
		Total:     int64(total),
		PageCount: int64(pageCount),
	}
}

func NewFromGinRequest(c *gin.Context, count int) *Pages {
	page := ParseInt(c.Query(PageVar), 1)
	pageSize := ParseInt(c.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// page转换categoryPage
func NewCategoryPages(page *Pages) *pb.Page {
	return &pb.Page{
		Page:      page.Page,
		PageSize:  page.PageSize,
		PageCount: page.PageCount,
		Total:     page.Total,
	}
}

// page转换productPage
func NewProductPages(page *Pages) *pb.ProductPage {
	return &pb.ProductPage{
		Page:      page.Page,
		PageSize:  page.PageSize,
		PageCount: page.PageCount,
		Total:     page.Total,
	}
}
