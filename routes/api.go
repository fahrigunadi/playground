package routes

import (
	"github.com/goravel/framework/facades"

	"github.com/fahrigunadi/playground/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()

	facades.Route().Resource("/api/persons", controllers.NewPersonController())
	facades.Route().Get("/users/{id}", userController.Show)
}
