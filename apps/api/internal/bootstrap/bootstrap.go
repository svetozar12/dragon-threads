package bootstrap

import (
	"dragon-threads/apps/api/internal/database"
	"dragon-threads/apps/api/internal/pkg/env"
	"dragon-threads/apps/api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/qinains/fastergoding"
)

func Bootstrap() {
	env.InitConfig()
	fastergoding.Run()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	database.Open()
	routes.InitRoutes(app)

	app.Listen(":" + env.Envs.Port)
}
