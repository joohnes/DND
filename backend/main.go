package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"dndEq/internal"
	"dndEq/routes"
)

func main() {
	log := slog.Default()
	db, err := internal.ConnectDB()
	if err != nil {
		panic(err)
	}
	srv := internal.NewService(db)

	app := fiber.New()
	app.Use(logger.New(), cors.New())
	register(app, srv)

	var port string
	flag.StringVar(&port, "port", "8000", "port to listen on")

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Error(err.Error())
		return
	}
}

func register(app *fiber.App, srv *internal.Service) {
	// player
	app.Post("/player", routes.CreatePlayerRoute(srv))
	app.Put("/player/:id", routes.UpdatePlayerRoute(srv))
	app.Get("/player/:id", routes.GetPlayerRoute(srv))
	app.Get("/players", routes.GetPlayersRoute(srv))
	app.Get("/players-names", routes.GetPlayerIdsWithNamesRoute(srv))
	app.Delete("/player/:id", routes.DeletePlayerRoute(srv))
	app.Post("/player/:playerID/add-item/:itemID", routes.AddItemRoute(srv))
	app.Delete("/player/drop-item/:itemID", routes.DropItemRoute(srv))
	app.Put("/player/hpmana/:id", routes.ChangeHPandManaRoute(srv))
	app.Get("/player/items/:playerID", routes.GetPlayerItemsRoute(srv))
	app.Post("/alcohol", routes.AlcoholLevelDownRoute(srv))
	app.Post("/player/zgon/:playerID", routes.ToggleZgonRoute(srv))

	// item
	app.Post("/item", routes.CreateItemRoute(srv))
	app.Get("/items", routes.GetItemsIDsRoute(srv))
	app.Get("/item/:id", routes.GetItemRoute(srv))
	app.Put("/item/:id", routes.UpdateItemRoute(srv))
	app.Delete("/item/:id", routes.DeleteItemRoute(srv))

	app.Post("/item/equip", routes.EquipItemRoute(srv))
	app.Delete("/item/unequip/:itemID", routes.UnequipItemRoute(srv))

	// bag
	app.Post("/bag/add/:itemID", routes.AddItemToBagRoute(srv))
	app.Get("/bag", routes.GetBagRoute(srv))
	app.Post("/bag/transfer/:itemID/:playerID", routes.TransferItemFromBagRoute(srv))
	app.Delete("/bag/drop/:itemID", routes.DropItemFromBagRoute(srv))
	app.Get("/bag/holder", routes.GetBagHolderRoute(srv))
	app.Post("/bag/holder", routes.ChangeBagHolderRoute(srv))
}
