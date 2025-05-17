package usecase

import (
	"context"
	"online-learning-platform/domain"

	log "github.com/sirupsen/logrus"
)

type courseUsecase struct {
	courseMySQLRepo domain.CourseMySQLRepository
}

func NewCourseUsecase(courseMySQLRepo domain.CourseMySQLRepository) domain.CourseUsecase {
	return &courseUsecase{
		courseMySQLRepo: courseMySQLRepo,
	}
}

func (cu *courseUsecase) CreateCourse(ctx context.Context, req domain.CourseRequest) (course domain.Course, err error) {
	id, err := cu.courseMySQLRepo.InsertCourse(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	course, err = cu.courseMySQLRepo.SelectCourseByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *courseUsecase) GetCourseByID(ctx context.Context, id int64) (course domain.Course, err error) {
	course, err = cu.courseMySQLRepo.SelectCourseByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *courseUsecase) GetAllCourse(ctx context.Context, page, limit int64, search string, sort string, categoryId ...*string) (course []domain.Course, err error) {
	offset := (page - 1) * limit
	course, err = cu.courseMySQLRepo.SelectAllCourse(ctx, offset, limit, search, sort, categoryId...)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *courseUsecase) UpdateCourse(ctx context.Context, req domain.CourseRequest, id int64) (course domain.Course, err error) {
	data, err := cu.courseMySQLRepo.SelectCourseByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	if req.CategoryID == 0 {
		req.CategoryID = data.CategoryID
	}

	if req.Name == "" {
		req.Name = data.Name
	}

	if req.Detail == nil {
		req.Detail = data.Detail
	}

	if req.Price == 0 {
		req.Price = data.Price
	}

	err = cu.courseMySQLRepo.EditCourse(ctx, req, id)
	if err != nil {
		log.Error(err)
		return
	}

	course, err = cu.courseMySQLRepo.SelectCourseByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *courseUsecase) DeleteCourse(ctx context.Context, id int64) (err error) {
	err = cu.courseMySQLRepo.RemoveCourse(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
