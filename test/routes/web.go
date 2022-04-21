package routes

import "test/controllers"

func webRoutes() {
	a.StaticDir("/assets ", "./static")

	a.GET("/", controllers.HomeController, "home")
}
