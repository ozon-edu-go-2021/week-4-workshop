package task_repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/ozonmp/week-4-workshop/category-service/internal/service/database"
	taskpkg "github.com/ozonmp/week-4-workshop/category-service/internal/service/task"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db: db}
}

func (r Repository) FindNonStartedTask(ctx context.Context) (*taskpkg.Task, error) {
	sb := database.StatementBuilder.
		Update("task").
		Set("started_at", time.Now()).
		Where(sq.Eq{"started_at": nil}).
		Suffix("RETURNING \"id\"")

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	task := new(taskpkg.Task)
	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(task)
	if err != nil {
		return nil, errors.Wrap(err, "db.QueryRowx()")
	}

	return task, nil
}
