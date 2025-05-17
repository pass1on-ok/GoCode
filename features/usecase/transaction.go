package usecase

import (
	"context"
	"online-learning-platform/domain"

	log "github.com/sirupsen/logrus"
)

type transactionUsecase struct {
	transactionMySQLRepo domain.TransactionMySQLRepository
	courseMySQLRepo      domain.CourseMySQLRepository
}

func NewTransactionUsecase(transactionMySQLRepo domain.TransactionMySQLRepository, courseMySQLRepo domain.CourseMySQLRepository) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionMySQLRepo: transactionMySQLRepo,
		courseMySQLRepo:      courseMySQLRepo,
	}
}

func (tu *transactionUsecase) CreateTransaction(ctx context.Context, req domain.TransactionRequest) (transaction domain.Transaction, err error) {
	course, err := tu.courseMySQLRepo.SelectCourseByID(ctx, req.CourseID)
	if err != nil {
		log.Error(err)
		return
	}
	req.Price = course.Price

	id, err := tu.transactionMySQLRepo.InsertTransaction(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	transaction, err = tu.transactionMySQLRepo.SelectTransactionByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (tu *transactionUsecase) GetTransactionByID(ctx context.Context, id int64) (transaction domain.Transaction, err error) {
	transaction, err = tu.transactionMySQLRepo.SelectTransactionByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
