package router

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc_shop "github.com/violetaplum/shop-grpc/proto/public_gen/go/shop"
	"google.golang.org/grpc"
)

type Router struct {
	engine     *gin.Engine
	shopClient grpc_shop.ShopServiceClient
	gwmux      *runtime.ServeMux
}

func NewRouter(conn *grpc.ClientConn) *Router {
	gwmux := runtime.NewServeMux()
	return &Router{
		engine:     gin.New(),
		shopClient: grpc_shop.NewShopServiceClient(conn),
		gwmux:      gwmux,
	}
}

func (r Router) Register(g *gin.RouterGroup) {
	g.POST("/product/list")
}
