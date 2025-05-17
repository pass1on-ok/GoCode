package handler

import (
	"online-learning-platform/domain"
	"online-learning-platform/utils/middlewares"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

type CategoryHandler struct {
	CategoryUsecase domain.CategoryUsecase
}

var (
	validate = validator.New()
)

func (ch *CategoryHandler) CreateCategory() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(fasthttp.StatusUnauthorized, err.Error())
		}

		if role != "admin" {
			return c.JSON(fasthttp.StatusUnauthorized, "admin only")
		}

		var input domain.CategoryRequest
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		err = validate.Struct(input)
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		err = ch.CategoryUsecase.CreateCategory(c.Request().Context(), input.Name)
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, "success")
	}
}

func (ch *CategoryHandler) GetCategory() echo.HandlerFunc {
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

		res, err := ch.CategoryUsecase.GetCategoryByID(c.Request().Context(), int64(id))
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (ch *CategoryHandler) GetCategories() echo.HandlerFunc {
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

		if page == 0 {
			page = 1
		}

		var limit int
		if c.QueryParam("limit") != "" {
			limit, err = strconv.Atoi(c.QueryParam("limit"))
			if err != nil {
				return c.JSON(fasthttp.StatusBadRequest, err.Error())
			}
		}

		if limit == 0 {
			limit = 10
		}

		sort := c.QueryParam("sort")

		res, err := ch.CategoryUsecase.GetAllCategory(c.Request().Context(), int64(page), int64(limit), sort)
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (ch *CategoryHandler) UpdateCategory() echo.HandlerFunc {
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

		var input domain.CategoryRequest
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		err = ch.CategoryUsecase.UpdateCategory(c.Request().Context(), input.Name, int64(id))
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, "success")
	}
}
