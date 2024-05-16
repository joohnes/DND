package routes

import (
	"dndEq/internal"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

func AddItemToBagRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute item id")
		}

		err = srv.DropItem(itemID)
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute")
		}

		err = srv.AddItemToBag(itemID)
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func TransferItemFromBagRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "TransferItemFromBagRoute item id")
		}

		playerID, err := strconv.Atoi(ctx.Params("playerID"))
		if err != nil {
			return errors.Wrap(err, "TransferItemFromBagRoute player id")
		}

		if !srv.PlayerExists(playerID) {
			return errors.Wrap(internal.ErrNoPlayer, "TransferItemFromBagRoute")
		}

		err = srv.DropItemFromBag(itemID)
		if err != nil {
			return errors.Wrap(err, "TransferItemFromBagRoute")
		}

		err = srv.AddItem(playerID, itemID)
		if err != nil {
			return errors.Wrap(err, "TransferItemFromBagRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func GetBagRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetBag())
	}
}

func DropItemFromBagRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "DropItemFromBagRoute")
		}

		err = srv.DropItemFromBag(itemID)
		if err != nil {
			return errors.Wrap(err, "DropItemFromBagRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func GetBagHolderRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetBagHolderName())
	}

}

func ChangeBagHolderRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var playerID int
		data := ctx.Body()
		if err := json.Unmarshal(data, &playerID); err != nil {
			return errors.Wrap(err, "ChangeBagHolderRoute")
		}

		err := srv.ChangeBagHolder(playerID)
		if err != nil {
			return errors.Wrap(err, "ChangeBagHolderRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
