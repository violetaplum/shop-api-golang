package model

type GetProductListParam struct {
	Page     int32 `json:"page" form:"page" binding:"required"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required"`
}
