package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type HttpController struct {
	// Dependent services
}

func NewHttpController() *HttpController {
	return &HttpController{
		// Inject services
	}
}

func (r *HttpController) Status(ctx http.Context) http.Response {
	status := ctx.Request().RouteInt("status")

	return ctx.Response().Status(status).String(http.StatusText(status))
}
