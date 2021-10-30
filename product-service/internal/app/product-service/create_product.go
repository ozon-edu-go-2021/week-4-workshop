package product_service

import (
	"context"

	product_service "github.com/ozonmp/week-3-workshop/product-service/internal/service/product"
	desc "github.com/ozonmp/week-3-workshop/product-service/pkg/product-service"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	res, err := i.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId())
	if err != nil {
		return nil, err
	}

	return &desc.CreateProductResponse{
		Result: convertProductToPb(res),
	}, nil
}

func convertProductToPb(res *product_service.Product) *desc.Product {
	return &desc.Product{
		Id:         res.ID,
		Name:       res.Name,
		CategoryId: res.CategoryID,
	}
}
