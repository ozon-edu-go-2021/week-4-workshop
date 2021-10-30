package product_service

import (
	desc "github.com/ozonmp/week-3-workshop/product-service/pkg/product-service"
)

type Implementation struct {
	desc.UnimplementedProductServiceServer
}

func NewProductService() desc.ProductServiceServer {
	return &Implementation{}
}
