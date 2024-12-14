package routes

import (
	"dndEq/internal"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

func CreateItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var item internal.Item
		if err := json.Unmarshal(ctx.Body(), &item); err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		if item.Name == "" {
			return errors.Wrap(internal.ErrEmptyName, "CreateItemRoute")
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
		ids, err := srv.GetItemsIDs()
		if err != nil {
			log.Println(errors.Wrap(err, "GetItemsIDs"))
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
		}
		return ctx.JSON(ids)
	}
}

func GetItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(err, "GetItemRoute")
		}

		i, err := srv.GetItemByID(id)
		if err != nil {
			return errors.Wrap(err, "GetItemRoute")
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

		if item.Name == "" {
			return errors.Wrap(internal.ErrEmptyName, "CreateItemRoute")
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
			PlayerID int `json:"playerID"`
			ItemID   int `json:"itemID"`
		}
		body := equipItem{}
		err := json.Unmarshal(ctx.Body(), &body)
		if err != nil {
			return errors.Wrap(err, "PlayerEquipItemRoute")
		}

		return srv.EquipItem(body.PlayerID, body.ItemID)
	}
}

func UnequipItemRoute(srv *internal.Service) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "PlayerUnequipItemRoute")
		}

		return srv.UnequipItem(itemID)
	}
}
