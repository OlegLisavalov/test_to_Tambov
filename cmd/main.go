package main

import (
	"log"
	"test_Task_New_server/config"
	"test_Task_New_server/database"
	"test_Task_New_server/handlers"
	"test_Task_New_server/logger"
	"test_Task_New_server/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()

	cfg := config.LoadConfig()

	db := database.ConnectDB(cfg)

	newsRepo := repository.NewNewsRepository(db)
	newsHandler := handlers.NewNewsHandler(newsRepo)

	app := fiber.New()

	app.Put("/news/:id", newsHandler.UpdateNews)
	app.Get("/news", newsHandler.GetAllNews)
	app.Get("/news/:id", newsHandler.GetNewsById) 

	log.Fatal(app.Listen(":8080"))
}
