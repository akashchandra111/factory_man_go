package main

import (
	"os"

	"github.com/akashchandra111/factory_man/src/configs" // Configs

	"github.com/gofiber/fiber/v2" // Fiber import
)

func main() {
	app := fiber.New()         // Intializing fiber app
	_, err := configs.InitDB() // Intializing DB

	if err != nil {
		os.Stderr.WriteString("Failed to init database, exiting")
		panic(err)
	}

	app.Get("/app/login", configs.GetName)

	app.Listen(":8911")
}
