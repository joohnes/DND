package main

import (
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

	app := fiber.New()
	app.Use(logger.New(), cors.New())
	// srv.TestObjects()
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
	app.Get("/players", srv.GetPlayersIDsRoute())
	app.Get("/players-names", srv.GetPlayerIdsWithNamesRoute())
	app.Delete("/player/:id", srv.DeletePlayerRoute())
	app.Post("/player/:playerID/add-item/:itemID", srv.AddItemRoute())
	app.Delete("/player/add-item/:itemID", srv.DropItemRoute())
	app.Put("/player/hpmana/:id", srv.ChangeHPandManaRoute())

	// item
	app.Post("/item", srv.CreateItemRoute())
	app.Get("/items", srv.GetItemsIDsRoute())
	app.Get("/item/:id", srv.GetItemRoute())
	app.Put("/item/:id", srv.UpdateItemRoute())
	app.Put("/equip/:id", srv.EquipItemRoute())
	app.Delete("/item/:id", srv.DeleteItemRoute())

	// bag
	app.Post("/bag/add/:itemID", srv.AddItemToBagRoute())
	app.Get("/bag", srv.GetBagRoute())
	app.Post("/bag/transfer/:itemID/:playerID", srv.TransferItemFromBagRoute())
}
