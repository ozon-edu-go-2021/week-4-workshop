package product_service

import "context"

//go:generate mockgen -package=product_service -destination=service_mocks_test.go -self_package=github.com/ozonmp/week-3-workshop/product-service/internal/service/product . IRepository

type IRepository interface {
	SaveProduct(ctx context.Context, product *Product) error
}

type Service struct {
	repo IRepository
}

func NewService() *Service {
	return &Service{
		repo: newRepo(),
	}
}

func (s *Service) CreateProduct(
	ctx context.Context,
	name string,
	categoryID int64,
) (*Product, error) {
	product := &Product{
		Name:       name,
		CategoryID: categoryID,
	}

	if err := s.repo.SaveProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
