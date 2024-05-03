package internal

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

func (srv *Service) CreatePlayerRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var player Player
		data := ctx.Body()
		if err := json.Unmarshal(data, &player); err != nil {
			return errors.Wrap(err, "CreatePlayerRoute")
		}

		p, err := srv.CreatePlayer(player)
		if err != nil {
			return errors.Wrap(err, "CreatePlayerRoute")
		}

		return ctx.JSON(p)
	}
}

func (srv *Service) GetPlayerRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(ErrIDNotNumber, "GetPlayerRoute")
		}

		p := srv.GetPlayerByID(id)
		if p == nil {
			return errors.Wrap(ErrNoPlayer, "GetPlayerRoute")
		}
		return ctx.JSON(p)
	}
}

func (srv *Service) GetPlayersRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetPlayers())
	}
}

func (srv *Service) AddItemRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		playerID, err := strconv.Atoi(ctx.Params("playerID"))
		if err != nil {
			return errors.Wrap(err, "AddItemRoute player id")
		}

		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "AddItemRoute item id")
		}

		return srv.AddItem(playerID, itemID)
	}
}

func (srv *Service) DropItemRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		playerID, err := strconv.Atoi(ctx.Params("playerID"))
		if err != nil {
			return errors.Wrap(err, "DropItemRoute player id")
		}

		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "DropItemRoute player id")
		}

		return srv.DropItem(playerID, itemID)
	}
}

func (srv *Service) UpdatePlayerRoute() func(ctx fiber.Ctx) error {
	// do poprawienia
	return func(ctx fiber.Ctx) error {
		return nil
	}
}
