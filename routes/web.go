package routes

import (
	"github.com/fahrigunadi/playground/app/http/controllers"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl")
	})

	facades.Route().Get("/editor", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("editor.tmpl")
	})

	facades.Route().Any("/http/status/{status}", controllers.NewHttpController().Status)

	facades.Route().Get("/image/{widthXHeight}", controllers.NewImageController().Index)

	facades.Route().StaticFile("/generate-images.sh", "public/generate-images.sh")
}
