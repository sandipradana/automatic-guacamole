package handler

import (
	"automatic-guacamole/internal/model"
	"automatic-guacamole/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductAddonGroupHandler struct {
	produtService service.ProductAddonGroupService
}

func NewProductAddonGroupHandler(productAddonGroupService service.ProductAddonGroupService) *ProductAddonGroupHandler {
	return &ProductAddonGroupHandler{produtService: productAddonGroupService}
}

func (h *ProductAddonGroupHandler) GetAll(c *fiber.Ctx) error {
	productAddonGroups, err := h.produtService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productAddonGroups)
}

func (h *ProductAddonGroupHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	productAddonGroup, err := h.produtService.GetByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ProductAddonGroup not found"})
	}
	return c.JSON(productAddonGroup)
}

func (h *ProductAddonGroupHandler) Create(c *fiber.Ctx) error {
	productAddonGroup := new(model.ProductAddonGroup)
	if err := c.BodyParser(productAddonGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.produtService.Create(productAddonGroup); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productAddonGroup)
}

func (h *ProductAddonGroupHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	productAddonGroup, err := h.produtService.GetByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ProductAddonGroup not found"})
	}

	if err := c.BodyParser(productAddonGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.produtService.Update(productAddonGroup); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productAddonGroup)
}

func (h *ProductAddonGroupHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.produtService.Delete(uint64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
