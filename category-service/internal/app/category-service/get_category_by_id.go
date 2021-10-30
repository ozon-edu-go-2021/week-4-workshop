package category_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	category_service "github.com/ozonmp/week-3-workshop/category-service/pkg/category-service"
)

func (i *Implementation) GetCategoryById(ctx context.Context, req *category_service.GetCategoryByIdRequest) (*category_service.GetCategoryByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
