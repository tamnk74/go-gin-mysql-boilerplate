package items

type CreateItem struct {
	Name string `form:"name" json:"name" binding:"required"`
}
