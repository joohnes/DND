package routes

import (
	"dndEq/internal"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

func CreateItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var item internal.Item
		data := ctx.Body()
		if err := json.Unmarshal(data, &item); err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		err := internal.IsValidRarity(item.Rarity)
		if err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		i, err := srv.CreateItem(item)
		if err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		return ctx.JSON(i)
	}
}

func GetItemsIDsRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetItemsIDs())
	}
}

func GetItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(err, "GetItemRoute")
		}

		i := srv.GetItemByID(id)
		if i == nil {
			return errors.Wrap(internal.ErrNoItem, "GetItemRoute")
		}

		return ctx.JSON(i)
	}
}

func UpdateItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var item internal.Item
		if err := json.Unmarshal(ctx.Body(), &item); err != nil {
			return errors.Wrap(err, "UpdateItemRoute")
		}

		err := srv.UpdateItem(item)
		if err != nil {
			return errors.Wrap(err, "UpdateItemRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func DeleteItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(err, "DeleteItemRoute")
		}

		err = srv.DeleteItem(id)
		if err != nil {
			return errors.Wrap(err, "DeleteItemRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func EquipItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		type equipItem struct {
			ItemID int `json:"id"`
			Slot   int `json:"slot"`
		}
		var ei equipItem
		if err := json.Unmarshal(ctx.Body(), &ei); err != nil {
			return errors.Wrap(err, "EquipItemRoute")
		}

		if err := srv.EquipItem(ei.ItemID, ei.Slot); err != nil {
			return errors.Wrap(err, "EquipItemRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func UnequipItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "UnequipItemRoute")
		}

		if err := srv.UnequipItem(itemID); err != nil {
			return errors.Wrap(err, "UnequipItemRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
