package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

func (srv *Service) CreateItemRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		var item Item
		if err := json.Unmarshal(ctx.Body(), &item); err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		i, err := srv.CreateItem(item)
		if err != nil {
			return errors.Wrap(err, "CreateItemRoute")
		}

		return ctx.JSON(i)
	}
}

func (srv *Service) GetItemsRoute() func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(srv.GetItems())
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

		return ctx.SendStatus(http.StatusOK)
	}
}
