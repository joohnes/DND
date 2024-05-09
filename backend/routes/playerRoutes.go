package routes

import (
	"dndEq/internal"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

func CreatePlayerRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var player internal.Player
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

func GetPlayerRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(internal.ErrIDNotNumber, "GetPlayerRoute")
		}

		p := srv.GetPlayerByID(id)
		if p == nil {
			return errors.Wrap(internal.ErrNoPlayer, "GetPlayerRoute")
		}

		return ctx.JSON(p)
	}
}

func GetPlayersIDsRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetPlayersIDs())
	}
}

func GetPlayerIdsWithNamesRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetPlayerIdsWithNames())
	}
}

func AddItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
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

func DropItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "DropItemRoute player id")
		}

		return srv.DropItem(itemID)
	}
}

func UpdatePlayerRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var player internal.Player
		data := ctx.Body()
		if err := json.Unmarshal(data, &player); err != nil {
			return errors.Wrap(err, "UpdatePlayerRoute")
		}

		err := srv.UpdatePlayer(player)
		if err != nil {
			return errors.Wrap(err, "UpdatePlayerRoute")
		}

		return ctx.SendStatus(http.StatusOK)
	}
}

func ChangeHPandManaRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		playerID, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(err, "ChangeHPandManaRoute player id")
		}

		var hpMana internal.HPMana
		data := ctx.Body()
		if err := json.Unmarshal(data, &hpMana); err != nil {
			return errors.Wrap(err, "ChangeHPandManaRoute")
		}

		err = srv.ChangeHPandMana(playerID, hpMana)
		if err != nil {
			return errors.Wrap(err, "ChangeHPandManaRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func DeletePlayerRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(internal.ErrIDNotNumber, "DeletePlayerRoute")
		}

		err = srv.DeletePlayer(id)
		if err != nil {
			return errors.Wrap(err, "DeletePlayerRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}