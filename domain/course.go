package domain

import (
	"context"
	"time"
)

type Course struct {
	ID         int64     `json:"id"`
	CategoryID int64     `json:"category_id"`
	Name       string    `json:"name"`
	Detail     *string   `json:"detail"`
	Price      float64   `json:"price"`
	Picture    string    `json:"picture"`
	DtmCrt     time.Time `json:"dtm_crt"`
	DtmUpd     time.Time `json:"dtm_upd"`
}

type CourseRequest struct {
	CategoryID int64   `json:"category_id" form:"category_id"`
	Name       string  `json:"name" form:"name"`
	Detail     *string `json:"detail" form:"detail"`
	Price      float64 `json:"price" form:"price"`
	Picture    string  `json:"picture" form:"picture"`
}

// CourseMySQLRepository is Course repository in MySQL
type CourseMySQLRepository interface {
	InsertCourse(ctx context.Context, req CourseRequest) (id int64, err error)
	SelectCourseByID(ctx context.Context, id int64) (course Course, err error)
	SelectAllCourse(ctx context.Context, page, limit int64, search string, sort string, categoryId ...*string) (course []Course, err error)
	EditCourse(ctx context.Context, req CourseRequest, id int64) (err error)
	RemoveCourse(ctx context.Context, id int64) (err error)
	SelectTotalCourse(ctx context.Context) (count int64, err error)
	SelectTotalFreeCourse(ctx context.Context) (count int64, err error)
}

// CourseUsecase is Course usecase
type CourseUsecase interface {
	CreateCourse(ctx context.Context, req CourseRequest) (course Course, err error)
	GetCourseByID(ctx context.Context, id int64) (course Course, err error)
	GetAllCourse(ctx context.Context, page, limit int64, search string, sort string, categoryId ...*string) (course []Course, err error)
	UpdateCourse(ctx context.Context, req CourseRequest, id int64) (course Course, err error)
	DeleteCourse(ctx context.Context, id int64) (err error)
}
