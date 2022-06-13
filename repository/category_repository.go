package repository

import (
	"context"
	"database/sql"
	"service_uaa/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx sql.Tx, categoryID int) domain.Category
	FindAll(ctx context.Context, tx sql.Tx) []domain.Category
}