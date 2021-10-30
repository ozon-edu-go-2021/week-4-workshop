package product_service

import (
	"context"

	desc "github.com/ozonmp/week-3-workshop/product-service/pkg/product-service"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	return &desc.CreateProductResponse{
		Result: &desc.Product{
			Id: 42,
		},
	}, nil
}
