package routes

import (
	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/group-chat/src/controllers"
	"github.com/louissaadgo/group-chat/src/middlewares"
)

func Setup(app *fiber.App) {

	app.Static("/public", "./src/public")

	app.Get("/login", controllers.RenderLogin)
	app.Post("/login", controllers.Login)

	app.Get("/logout", controllers.Logout)

	app.Get("/register", controllers.RenderRegister)
	app.Post("/register", controllers.Register)

	app.Use(middlewares.IsAuthenticated)

	app.Use("/ws", middlewares.UpgradeToWS)

	controllers.RegisterWSEvents()

	app.Get("/", controllers.Home)

	app.Get("/ws", ikisocket.New(controllers.WS))

}
