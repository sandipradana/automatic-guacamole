package main

import (
	"automatic-guacamole/internal/handler"
	"automatic-guacamole/internal/middleware"
	"automatic-guacamole/internal/repository"
	"automatic-guacamole/internal/service"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var router *fiber.App

func init() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIMEZONE"))
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            false,
	})

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(db, productRepo)
	productHandler := handler.NewProductHandler(productService)

	productAddonRepo := repository.NewProductAddonRepository()
	productAddonService := service.NewProductAddonService(db, productAddonRepo)
	productAddonHandler := handler.NewProductAddonHandler(productAddonService)

	productAddonGroupRepo := repository.NewProductAddonGroupRepository()
	productAddonGroupService := service.NewProductAddonGroupService(db, productAddonGroupRepo)
	productAddonGroupHandler := handler.NewProductAddonGroupHandler(productAddonGroupService)

	cartService := service.NewCartService(db, productRepo, productAddonGroupRepo, productAddonRepo)
	cartHandler := handler.NewCartHandler(cartService)

	router = fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			errFiber, ok := err.(*fiber.Error)
			if ok {
				return c.Status(errFiber.Code).JSON(fiber.Map{"error": errFiber.Message})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		},
	})

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World !")
	})

	router.Get("/users", middleware.Auth(userService), userHandler.GetAll)
	router.Post("/users/login", userHandler.Login)
	router.Get("/users/:id", middleware.Auth(userService), userHandler.Get)
	router.Post("/users", middleware.Auth(userService), userHandler.Create)
	router.Put("/users/:id", middleware.Auth(userService), userHandler.Update)
	router.Delete("/users/:id", middleware.Auth(userService), userHandler.Delete)

	router.Get("/products", middleware.Auth(userService), productHandler.GetAll)
	router.Get("/products/:id", middleware.Auth(userService), productHandler.Get)
	router.Post("/products", middleware.Auth(userService), productHandler.Create)
	router.Put("/products/:id", middleware.Auth(userService), productHandler.Update)
	router.Delete("/products/:id", middleware.Auth(userService), productHandler.Delete)

	router.Get("/product-addons", middleware.Auth(userService), productAddonHandler.GetAll)
	router.Get("/product-addons/:id", middleware.Auth(userService), productAddonHandler.Get)
	router.Post("/product-addons", middleware.Auth(userService), productAddonHandler.Create)
	router.Put("/product-addons/:id", middleware.Auth(userService), productAddonHandler.Update)
	router.Delete("/product-addons/:id", middleware.Auth(userService), productAddonHandler.Delete)

	router.Get("/product-addon-groups", middleware.Auth(userService), productAddonGroupHandler.GetAll)
	router.Get("/product-addon-groups/:id", middleware.Auth(userService), productAddonGroupHandler.Get)
	router.Post("/product-addon-groups", middleware.Auth(userService), productAddonGroupHandler.Create)
	router.Put("/product-addon-groups/:id", middleware.Auth(userService), productAddonGroupHandler.Update)
	router.Delete("/product-addon-groups/:id", middleware.Auth(userService), productAddonGroupHandler.Delete)

	router.Post("/cart/calculate", cartHandler.Calculate)
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	err := router.Listen(":" + port)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
