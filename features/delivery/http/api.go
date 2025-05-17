package http

import (
	"online-learning-platform/domain"
	"online-learning-platform/features/delivery/http/handler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteAPI(e *echo.Echo, category domain.CategoryUsecase, user domain.UserUsecase, course domain.CourseUsecase, transaction domain.TransactionUsecase) {
	handlerCategory := handler.CategoryHandler{CategoryUsecase: category}
	handlerUser := handler.UserHandler{UserUsecase: user}
	handlerCourse := handler.CourseHandler{CourseUsecase: course}
	handlerTransaction := handler.TransactionHandler{TransactionUsecase: transaction}

	// Login
	e.POST("/login", handlerUser.Login())

	// Category
	e.POST("/category", handlerCategory.CreateCategory(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/category", handlerCategory.GetCategories(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/category/:id", handlerCategory.GetCategory(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.PUT("/category/:id", handlerCategory.UpdateCategory(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// User
	e.POST("/user", handlerUser.CreateUser())
	e.DELETE("/user/:id", handlerUser.DeleteUser(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Course
	e.POST("/course", handlerCourse.CreateCourse(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/course/:id", handlerCourse.GetCourseByID(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/course", handlerCourse.GetAllCourse(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.PATCH("/course/:id", handlerCourse.UpdateCourse(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.DELETE("/course/:id", handlerCourse.DeleteCourse(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Transaction
	e.POST("/transaction", handlerTransaction.CreateTransaction(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/transaction/:id", handlerTransaction.GetTransactionByID(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Statistics
	e.GET("/statistics", handlerUser.GetSimpleStatistics(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}
