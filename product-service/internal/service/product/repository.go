package product_service

import (
	"context"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

var nextID int64 = 1

type repository struct{
	DB *sqlx.DB
}

func newRepo(db *sqlx.DB) IRepository {
	return repository{
		DB: db,
	}
}

func (r repository) SaveProduct(ctx context.Context, product *Product) error {
	query := sq.Insert("products").PlaceholderFormat(sq.Dollar).Columns("name", "category_id").Values(
		product.Name, product.CategoryID).RunWith(r.DB)

	_, err := query.ExecContext(ctx)

	return err
}
