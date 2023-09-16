package integration

import (
	"dragon-threads/apps/api-e2e/integration/usersIntegration"
	"dragon-threads/apps/api/database"
	"dragon-threads/apps/api/pkg/env"
	"dragon-threads/apps/api/routes"
	"fmt"
	"sync"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func TestIntegration(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	startTestServer(&wg)
	// users
	t.Run("GET USER", usersIntegration.GetUserIntegration)
}

func startTestServer(wg *sync.WaitGroup) {
	env.InitConfig()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	database.Open()
	routes.InitRoutes(app)
	fmt.Println("HERE")

}
