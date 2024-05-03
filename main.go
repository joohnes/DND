package main

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"dndEq/internal"
)

func main() {
	log := slog.Default()
	db, err := internal.ConnectDB()
	if err != nil {
		panic(err)
	}
	srv, err := internal.NewService(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(srv.GetPlayers())
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	register(app, srv)

	if err := app.Listen(":8000"); err != nil {
		log.Error(err.Error())
		return
	}
}

func register(app *fiber.App, srv *internal.Service) {
	// player
	app.Post("/player", srv.CreatePlayerRoute())
	app.Put("/player/:id", srv.UpdatePlayerRoute())
	app.Get("/player/:id", srv.GetPlayerRoute())
	app.Get("/players", srv.GetPlayersRoute())
	app.Post("/player/:playerID/add-item/:itemID", srv.AddItemRoute())
	app.Delete("/player/:playerID/add-item/:itemID", srv.DropItemRoute())

	// item
	app.Post("/item", srv.CreateItemRoute())
	app.Get("/items", srv.GetItemsRoute())
	app.Get("/item/:id", srv.GetItemRoute())
	app.Put("/item/:id", srv.UpdateItemRoute())
}
