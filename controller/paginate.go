package controller

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginatedResponse struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
}

func NewPaginatedResponse(c *gin.Context) PaginatedResponse {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	sort := c.Query("sort")

	return PaginatedResponse{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func (p *PaginatedResponse) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginatedResponse) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PaginatedResponse) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginatedResponse) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func Paginate(model interface{}, pr *PaginatedResponse) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(model).Count(&totalRows)

	pr.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pr.Limit)))

	if totalPages <= 0 {
		totalPages = 1
	}

	pr.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pr.GetOffset()).Limit(pr.GetLimit()).Order(pr.GetSort())
	}
}
