package handler

import (
	"automatic-guacamole/internal/model"
	"automatic-guacamole/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductAddonHandler struct {
	produtService service.ProductAddonService
}

func NewProductAddonHandler(productAddonService service.ProductAddonService) *ProductAddonHandler {
	return &ProductAddonHandler{produtService: productAddonService}
}

func (h *ProductAddonHandler) GetAll(c *fiber.Ctx) error {
	productAddons, err := h.produtService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productAddons)
}

func (h *ProductAddonHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	productAddon, err := h.produtService.GetByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ProductAddon not found"})
	}
	return c.JSON(productAddon)
}

func (h *ProductAddonHandler) Create(c *fiber.Ctx) error {
	productAddon := new(model.ProductAddon)
	if err := c.BodyParser(productAddon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.produtService.Create(productAddon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productAddon)
}

func (h *ProductAddonHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	productAddon, err := h.produtService.GetByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ProductAddon not found"})
	}

	if err := c.BodyParser(productAddon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.produtService.Update(productAddon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productAddon)
}

func (h *ProductAddonHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.produtService.Delete(uint64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
