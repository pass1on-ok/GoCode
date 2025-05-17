package domain

import (
	"context"
	"time"
)

type Category struct {
	ID     int64      `json:"id,omitempty"`
	Name   string     `json:"name,omitempty"`
	Count  *int       `json:"most_popular,omitempty"`
	DtmCrt *time.Time `json:"dtm_crt,omitempty"`
	DtmUpd *time.Time `json:"dtm_upd,omitempty"`
}

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

// CategoryMySQLRepository is Category repository in MySQL
type CategoryMySQLRepository interface {
	InsertCategory(ctx context.Context, name string) (err error)
	SelectCategoryByID(ctx context.Context, id int64) (category Category, err error)
	SelectAllCategory(ctx context.Context, page, limit int64, sort string) (category []Category, err error)
	EditCategory(ctx context.Context, name string, id int64) (err error)
}

// CategoryUsecase is Category usecase
type CategoryUsecase interface {
	CreateCategory(ctx context.Context, name string) (err error)
	GetCategoryByID(ctx context.Context, id int64) (category Category, err error)
	GetAllCategory(ctx context.Context, page, limit int64, sort string) (category []Category, err error)
	UpdateCategory(ctx context.Context, name string, id int64) (err error)
}
