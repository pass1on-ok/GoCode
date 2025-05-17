package handler

import (
	"online-learning-platform/domain"
	"online-learning-platform/utils/aws"
	"online-learning-platform/utils/middlewares"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

type CourseHandler struct {
	CourseUsecase domain.CourseUsecase
}

func (ch *CourseHandler) CreateCourse() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(fasthttp.StatusUnauthorized, err.Error())
		}

		if role != "admin" {
			return c.JSON(fasthttp.StatusUnauthorized, "admin only")
		}

		var input domain.CourseRequest
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		file, fileheader, _ := c.Request().FormFile("picture")
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		if file != nil {
			input.Picture, err = aws.UploadFile(file, fileheader)
			if err != nil {
				return c.JSON(fasthttp.StatusBadRequest, err.Error())
			}
		}

		res, err := ch.CourseUsecase.CreateCourse(c.Request().Context(), input)
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (ch *CourseHandler) GetCourseByID() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(fasthttp.StatusUnauthorized, err.Error())
		}

		if role != "user" {
			return c.JSON(fasthttp.StatusUnauthorized, "user only")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		res, err := ch.CourseUsecase.GetCourseByID(c.Request().Context(), int64(id))
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (ch *CourseHandler) GetAllCourse() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(fasthttp.StatusUnauthorized, err.Error())
		}

		if role != "user" {
			return c.JSON(fasthttp.StatusUnauthorized, "user only")
		}

		var page int
		if c.QueryParam("page") != "" {
			page, err = strconv.Atoi(c.QueryParam("page"))
			if err != nil {
				return c.JSON(fasthttp.StatusBadRequest, err.Error())
			}
		}

		var limit int
		if c.QueryParam("limit") != "" {
			limit, err = strconv.Atoi(c.QueryParam("limit"))
			if err != nil {
				return c.JSON(fasthttp.StatusBadRequest, err.Error())
			}
		}

		sort := c.QueryParam("sort")
		search := c.QueryParam("search")

		var cat *string
		categoryId := c.QueryParam("category_id")
		if categoryId != "" {
			cat = &categoryId
		}

		res, err := ch.CourseUsecase.GetAllCourse(c.Request().Context(), int64(page), int64(limit), search, sort, cat)
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (ch *CourseHandler) UpdateCourse() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(fasthttp.StatusUnauthorized, err.Error())
		}

		if role != "admin" {
			return c.JSON(fasthttp.StatusUnauthorized, "admin only")
		}

		var input domain.CourseRequest
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		res, err := ch.CourseUsecase.UpdateCourse(c.Request().Context(), input, int64(id))
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (ch *CourseHandler) DeleteCourse() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(fasthttp.StatusUnauthorized, err.Error())
		}

		if role != "admin" {
			return c.JSON(fasthttp.StatusUnauthorized, "admin only")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		err = ch.CourseUsecase.DeleteCourse(c.Request().Context(), int64(id))
		if err != nil {
			if strings.Contains(err.Error(), "no") {
				return c.JSON(fasthttp.StatusNotFound, err.Error())
			}
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, "success")
	}
}
