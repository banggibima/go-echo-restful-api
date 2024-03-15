package main

import (
	"fmt"
	"log"

	"github.com/banggibima/go-echo-restful-api/internal/app"
	"github.com/banggibima/go-echo-restful-api/internal/config"
	"github.com/banggibima/go-echo-restful-api/internal/database"
	"github.com/banggibima/go-echo-restful-api/internal/handlers"
	"github.com/banggibima/go-echo-restful-api/internal/repositories"
	"github.com/banggibima/go-echo-restful-api/internal/usecases"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatalf("error establishing database connection: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	e := echo.New()
	myApp := app.NewApp(userHandler)
	myApp.SetupRoutes(e)

	port := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := e.Start(port); err != nil {
		log.Fatalf("error starting the application: %v", err)
	}
}
