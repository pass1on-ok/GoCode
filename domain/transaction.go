package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"user_id"`
	CourseID int64     `json:"course_id"`
	Price    float64   `json:"price"`
	DtmCrt   time.Time `json:"dtm_crt"`
	DtmUpd   time.Time `json:"dtm_upd"`
}

type TransactionRequest struct {
	UserID   int64   `json:"user_id" form:"user_id"`
	CourseID int64   `json:"course_id" form:"course_id"`
	Price    float64 `json:"price" form:"price"`
}

// TransactionMySQLRepository is Transaction repository in MySQL
type TransactionMySQLRepository interface {
	InsertTransaction(ctx context.Context, req TransactionRequest) (id int64, err error)
	SelectTransactionByID(ctx context.Context, id int64) (transaction Transaction, err error)
}

// TransactionUsecase is Transaction usecase
type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, req TransactionRequest) (transaction Transaction, err error)
	GetTransactionByID(ctx context.Context, id int64) (transaction Transaction, err error)
}
