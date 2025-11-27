package main

import (
	"log"
	"os"

	"github.com/descooly/star-catalog/internal/db"
	"github.com/descooly/star-catalog/internal/handler"
	"github.com/descooly/star-catalog/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	e := echo.New()

	rep := service.NewRepo(database)
	hand := handler.NewHandler(rep)

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "./web"
	}

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Static("/static", staticDir+"/static")

	e.File("/*", staticDir+"/index.html")

	e.GET("/api/stars", hand.GetStars)
	e.POST("/api/stars", hand.PostStars)

	e.Logger.Fatal(e.Start(":1323"))
}
