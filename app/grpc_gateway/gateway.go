package grpc_gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc_shop "github.com/violetaplum/shop-grpc/proto/public_gen/go/shop"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterShopGrpc(mux *runtime.ServeMux) {
	svc := "127.0.0.1"
	port := "9095"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := grpc_shop.RegisterShopServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		fmt.Sprintf("%s:%s", svc, port),
		opts,
	)
	if err != nil {
		panic(err)
	}
}
