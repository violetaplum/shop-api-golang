package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"shop-api-golang/app/grpc_gateway"
)

func main() {
	mux := runtime.NewServeMux()
	grpc_gateway.RegisterShopGrpc(mux)
	route := gin.New()
	v1 := route.Group("/api/v1")
	v1.Any("/")
}
