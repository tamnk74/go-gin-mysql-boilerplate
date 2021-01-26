package dto

import "github.com/gin-gonic/gin"

type Pagination struct {
	Limit   int `form:"limit"`
	Page    int `form:"page"`
	PerPage int
	Total   int64
	Offset  int
}

func (a *Pagination) FillDefault() {
	if a.Page == 0 {
		a.Page = 1
	}
	if a.Limit == 0 {
		a.Limit = 15
	}
	a.PerPage = a.Limit
	a.Offset = (a.Page - 1) * a.Limit
	if a.Limit == -1 {
		a.Offset = -1
	}
}

func (a *Pagination) Update() {
	if a.Limit == -1 {
		a.PerPage = int(a.Total)
	}
}

func (a *Pagination) GetMeta() gin.H {
	return gin.H{
		"total":    a.Total,
		"page":     a.Page,
		"per_page": a.PerPage,
	}
}
