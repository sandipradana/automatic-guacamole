package handler

import (
	"automatic-guacamole/internal/model"
	"automatic-guacamole/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(service service.CartService) *CartHandler {
	return &CartHandler{cartService: service}
}

func (h *CartHandler) Calculate(c *fiber.Ctx) error {
	cart := new(model.Cart)
	if err := c.BodyParser(cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	total, err := h.cartService.Calculate(cart)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"total": total})
}
