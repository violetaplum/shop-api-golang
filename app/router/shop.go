package router

import (
	"github.com/gin-gonic/gin"
	grpc_shop "github.com/violetaplum/shop-grpc/proto/public_gen/go/shop"
	"net/http"
	shop_error "shop-api-golang/internal/error"
	model "shop-api-golang/internal/model"
)

func (r *Router) GetProductList(c *gin.Context) {
	param := model.GetProductListParam{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, shop_error.ErrorResponse{
			Code:    "INVALID_ARGUMENT",
			Message: err.Error(),
		})
		return
	}
	resp, err := r.shopClient.GetProductList(c.Request.Context(),
		&grpc_shop.GetProductListRequest{})
	if err != nil {
		shop_error.HandleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
