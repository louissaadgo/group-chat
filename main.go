package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/louissaadgo/group-chat/src/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./src/public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Setup(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
