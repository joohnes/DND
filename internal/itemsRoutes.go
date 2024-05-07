package internal

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

func (srv *Service) CreateItemRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var item Item
		data := ctx.Body()
		fmt.Println(string(data))
		if err := json.Unmarshal(data, &item); err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		i, err := srv.CreateItem(item)
		if err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		return ctx.JSON(i)
	}
}

func (srv *Service) GetItemsIDsRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetItemsIDs())
	}
}

func (srv *Service) GetItemRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return errors.Wrap(err, "GetItemRoute")
		}

		i := srv.GetItemByID(id)
		if i == nil {
			return errors.Wrap(ErrNoItem, "GetItemRoute")
		}

		return ctx.JSON(i)
	}
}

func (srv *Service) UpdateItemRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var item Item
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

func (srv *Service) EquipItemRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "EquipItemRoute item id")
		}

		playerID, err := strconv.Atoi(ctx.Params("playerID"))
		if err != nil {
			return errors.Wrap(err, "EquipItemRoute player id")
		}

		err = srv.EquipItem(itemID, playerID)
		if err != nil {
			return errors.Wrap(err, "EquipItemRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

// bag routes
func (srv *Service) AddItemToBagRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute item id")
		}

		playerID, err := strconv.Atoi(ctx.Params("playerID"))
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute player id")
		}

		bagID, err := srv.GetBagID()
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute")
		}

		err = srv.DropItem(playerID, itemID)
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute")
		}

		err = srv.AddItem(bagID, itemID)
		if err != nil {
			return errors.Wrap(err, "AddItemToBagRoute")
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (srv *Service) TransferItemFromBagRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		itemID, err := strconv.Atoi(ctx.Params("itemID"))
		if err != nil {
			return errors.Wrap(err, "TransferItemFromBagRoute item id")
		}

		playerID, err := strconv.Atoi(ctx.Params("playerID"))
		if err != nil {
			return errors.Wrap(err, "TransferItemFromBagRoute player id")
		}

		bagID, err := srv.GetBagID()
		if err != nil {
			return errors.Wrap(err, "TransferItemFromBagRoute")
		}

		err = srv.DropItem(bagID, itemID)
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

func (srv *Service) GetBagRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetBag())
	}
}
