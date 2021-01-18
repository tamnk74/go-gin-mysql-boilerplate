package schema

type Pagination struct {
	Limit int `form:"limit" binding:"required"`
	Page  int `form:"page" binding:"gte=1"`
}
