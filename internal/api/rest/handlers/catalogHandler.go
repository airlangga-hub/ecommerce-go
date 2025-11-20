package handlers

import (
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
	"github.com/airlangga-hub/ecommerce-go/internal/service"
	"github.com/gofiber/fiber/v2"
)


type CatalogHandler struct {
	Svc *service.CatalogService
}


func SetupCatalogRoutes(rh *rest.HttpHandler) {
	app := rh.App

	catalogService := &service.CatalogService{
		CatalogRepo: repository.NewCatalogRepository(rh.DB),
		Auth: rh.Auth,
		Config: rh.Config,
	}
	handler := &CatalogHandler{Svc: catalogService}

	// Public endpoints
	app.Get("/products")
	app.Get("/products/:id")

	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryByID)

	// Private endpoints
	sellerRoutes := app.Group("/seller", handler.Svc.Auth.Authorize)

	sellerRoutes.Post("/categories", handler.CreateCategories)
	sellerRoutes.Patch("/categories/:id", handler.EditCategory)
	sellerRoutes.Delete("/categories/:id", handler.DeleteCategory)

	sellerRoutes.Post("/products", handler.CreateProducts)
	sellerRoutes.Get("/products", handler.GetProducts)
	sellerRoutes.Get("/products/:id", handler.GetProduct)
	sellerRoutes.Put("/products/:id", handler.EditProduct)
	sellerRoutes.Patch("/products/:id", handler.UpdateStock) // update stock
	sellerRoutes.Delete("/products/:id", handler.DeleteProduct)
}


func (h *CatalogHandler) GetCategories(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "get categories", nil)
}


func (h *CatalogHandler) GetCategoryByID(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "get category by id", nil)
}


func (h *CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	// user := h.GetCurrentUser(ctx)



	return rest.OkResponse(ctx, "create categories", nil)
}


func (h *CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "edit category", nil)
}


func (h *CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "delete category", nil)
}


func (h *CatalogHandler) CreateProducts(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "create products", nil)
}


func (h *CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "get products", nil)
}


func (h *CatalogHandler) GetProduct(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "get products", nil)
}


func (h *CatalogHandler) EditProduct(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "get products", nil)
}


func (h *CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "get products", nil)
}


func (h *CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	return rest.OkResponse(ctx, "get products", nil)
}