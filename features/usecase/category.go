package usecase

import (
	"context"
	"online-learning-platform/domain"

	log "github.com/sirupsen/logrus"
)

type categoryUsecase struct {
	categoryMySQLRepo domain.CategoryMySQLRepository
}

func NewCategoryUsecase(categoryMySQLRepo domain.CategoryMySQLRepository) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryMySQLRepo: categoryMySQLRepo,
	}
}

func (cu *categoryUsecase) CreateCategory(ctx context.Context, name string) (err error) {
	err = cu.categoryMySQLRepo.InsertCategory(ctx, name)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *categoryUsecase) GetCategoryByID(ctx context.Context, id int64) (category domain.Category, err error) {
	category, err = cu.categoryMySQLRepo.SelectCategoryByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *categoryUsecase) GetAllCategory(ctx context.Context, page, limit int64, sort string) (category []domain.Category, err error) {
	offset := (page - 1) * limit
	category, err = cu.categoryMySQLRepo.SelectAllCategory(ctx, offset, limit, sort)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *categoryUsecase) UpdateCategory(ctx context.Context, name string, id int64) (err error) {
	err = cu.categoryMySQLRepo.EditCategory(ctx, name, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
