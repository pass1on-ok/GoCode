package usecase

import (
	"context"
	"errors"
	"online-learning-platform/domain"
	"online-learning-platform/utils/middlewares"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userMySQLRepo   domain.UserMySQLRepository
	courseMySQLRepo domain.CourseMySQLRepository
}

func NewUserUsecase(userMySQLRepo domain.UserMySQLRepository, courseMySQLRepo domain.CourseMySQLRepository) domain.UserUsecase {
	return &userUsecase{
		userMySQLRepo:   userMySQLRepo,
		courseMySQLRepo: courseMySQLRepo,
	}
}

func (uu *userUsecase) GetUserLogin(ctx context.Context, req domain.LoginRequest) (user domain.UserLogin, err error) {
	data, err := uu.userMySQLRepo.SelectUserLogin(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		err = errors.New("password not match")
		log.Error(err)
		return
	}
	user.Token = middlewares.GenerateToken(int(data.ID), data.Role)
	user.ID = data.ID
	user.Email = data.Email
	user.Name = data.Name
	user.Picture = data.Picture
	user.Role = data.Role
	user.DtmCrt = data.DtmCrt
	user.DtmUpd = data.DtmUpd

	return
}

func (uu *userUsecase) CreateUser(ctx context.Context, req domain.UserRequest) (user domain.User, err error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New("cannot encrypt password")
		log.Error(err)
		return
	}
	req.Password = string(generate)
	req.Role = "user"
	req.Deleted = false

	id, err := uu.userMySQLRepo.InsertUser(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	user, err = uu.userMySQLRepo.SelectUserByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (uu *userUsecase) DeleteUser(ctx context.Context, id int64) (err error) {
	err = uu.userMySQLRepo.RemoveUser(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (uu *userUsecase) GetSimpleStatistics(ctx context.Context) (statistic domain.SimpleStatisticsResponse, err error) {
	totalUser, err := uu.userMySQLRepo.SelectTotalUser(ctx)
	if err != nil {
		log.Error(err)
		totalUser = 0
	}

	totalCourse, err := uu.courseMySQLRepo.SelectTotalCourse(ctx)
	if err != nil {
		log.Error(err)
		totalCourse = 0
	}

	totalFreeCourse, err := uu.courseMySQLRepo.SelectTotalFreeCourse(ctx)
	if err != nil {
		log.Error(err)
		totalFreeCourse = 0
	}

	statistic.TotalUser = totalUser
	statistic.TotalCourse = totalCourse
	statistic.TotalFreeCourse = totalFreeCourse

	return
}
