package product_service

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) *Service {
	ctrl := gomock.NewController(t)
	repo := NewMockIRepository(ctrl)

	service := &Service{
		repo: repo,
	}

	repo.EXPECT().
		SaveProduct(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, product *Product) error {
			product.ID = 124
			return nil
		})

	return service
}

func TestCreateProduct_Success_ReturnName(t *testing.T) {
	service := setup(t)

	product, err := service.CreateProduct(
		context.Background(),
		"new product",
		4,
	)

	require.Nil(t, err)
	require.Equal(t, "new product", product.Name)
}

func TestCreateProduct_Success_ReturnCategoryID(t *testing.T) {
	service := setup(t)

	product, err := service.CreateProduct(
		context.Background(),
		"new product",
		25,
	)

	require.Nil(t, err)
	require.Equal(t, int64(25), product.CategoryID)
}

func TestCreateProduct_Success_ReturnID(t *testing.T) {
	service := setup(t)

	product, err := service.CreateProduct(
		context.Background(),
		"another product",
		3,
	)

	require.Nil(t, err)
	require.Equal(t, int64(124), product.ID)
}
