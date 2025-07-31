package controllers

import (
	"github.com/fahrigunadi/playground/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type PersonController struct {
	// Dependent services
}

func NewPersonController() *PersonController {
	return &PersonController{
		// Inject services
	}
}

func (r *PersonController) Index(ctx http.Context) http.Response {
	var persons []models.Person
	err := facades.Orm().Query().Get(&persons)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":   err.Error(),
			"message": "Internal Server Error",
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "Success",
		"data":    persons,
	})
}

func (r *PersonController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")

	var person models.Person
	err := facades.Orm().Query().FindOrFail(&person, id)

	if err != nil {
		return ctx.Response().Json(404, http.Json{
			"error":   err.Error(),
			"message": "Person Not Found",
		})
	}

	return ctx.Response().Success().Json(person)
}

func (r *PersonController) Store(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		"name":    "required|max_len:255",
		"email":   "required|email|max_len:255",
		"phone":   "required|max_len:255",
		"age":     "required|numeric",
		"address": "max_len:6000",
	})

	if err != nil {
		return ctx.Response().Json(500, http.Json{
			"error":   err.Error(),
			"message": "Internal Server Error",
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"errors":  validator.Errors().All(),
			"message": "Unprocessable Entity",
		})
	}

	var person models.Person
	person.Name = ctx.Request().Input("name")
	person.Email = ctx.Request().Input("email")
	person.Phone = ctx.Request().Input("phone")
	person.Age = ctx.Request().InputInt("age")
	person.Address = ctx.Request().Input("address")

	if err := facades.Orm().Query().Create(&person); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":   err.Error(),
			"message": "Failed to create post",
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"data":    person,
		"message": "Success",
	})
}

func (r *PersonController) Update(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")

	var person models.Person

	if err := facades.Orm().Query().FindOrFail(&person, id); err != nil {
		return ctx.Response().Json(404, http.Json{
			"error":   err.Error(),
			"message": "Person Not Found",
		})
	}

	validator, err := ctx.Request().Validate(map[string]string{
		"name":    "required|max_len:255",
		"email":   "required|email|max_len:255",
		"phone":   "required|max_len:255",
		"age":     "required|numeric",
		"address": "max_len:6000",
	})

	if err != nil {
		return ctx.Response().Json(500, http.Json{
			"error":   err.Error(),
			"message": "Internal Server Error",
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"errors":  validator.Errors().All(),
			"message": "Unprocessable Entity",
		})
	}

	person.Name = ctx.Request().Input("name")
	person.Email = ctx.Request().Input("email")
	person.Phone = ctx.Request().Input("phone")
	person.Age = ctx.Request().InputInt("age")
	person.Address = ctx.Request().Input("address")

	if err := facades.Orm().Query().Save(&person); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":   err.Error(),
			"message": "Failed to create post",
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"data":    person,
		"message": "Success",
	})
}

func (r *PersonController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")

	var person models.Person

	if err := facades.Orm().Query().FindOrFail(&person, id); err != nil {
		return ctx.Response().Json(404, http.Json{
			"message": "Person Not Found",
		})
	}

	facades.Orm().Query().Delete(&person)

	return ctx.Response().Success().Json(http.Json{
		"message": "Success",
	})
}
