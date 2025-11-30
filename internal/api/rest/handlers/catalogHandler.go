package handlers

import (
	"strconv"

	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
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
		Repo: repository.NewCatalogRepository(rh.DB),
		Auth: rh.Auth,
		Config: rh.Config,
	}
	handler := &CatalogHandler{Svc: catalogService}

	// Public endpoints
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProductByID)

	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryByID)

	// Private endpoints
	sellerRoutes := app.Group("/seller", handler.Svc.Auth.AuthorizeSeller)

	sellerRoutes.Post("/categories", handler.CreateCategory)
	sellerRoutes.Patch("/categories/:id", handler.EditCategory)
	sellerRoutes.Delete("/categories/:id", handler.DeleteCategory)

	sellerRoutes.Post("/products", handler.CreateProduct)
	sellerRoutes.Get("/products", handler.GetSellerProducts)
	sellerRoutes.Get("/products/:id", handler.GetProductByID)
	sellerRoutes.Put("/products/:id", handler.EditProduct)
	sellerRoutes.Patch("/products/:id", handler.UpdateStock) // update stock
	sellerRoutes.Delete("/products/:id", handler.DeleteProduct)
}


func (h *CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	categories, err := h.Svc.GetCategories()

	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}

	return rest.OkResponse(ctx, "get categories success", categories)
}


func (h *CatalogHandler) GetCategoryByID(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	category, err := h.Svc.GetCategoryByID(uint(id))
	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}

	return rest.OkResponse(ctx, "get category by id success", category)
}


func (h *CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {

	createCategory := dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(&createCategory); err != nil {
		return rest.BadRequest(ctx, "invalid request body")
	}

	if err := h.Svc.CreateCategory(createCategory); err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return rest.OkResponse(ctx, "category created successfully", nil)
}


func (h *CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	createCategory := dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(&createCategory); err != nil {
		return rest.BadRequest(ctx, "invalid request body to update category")
	}

	category, err := h.Svc.EditCategory(uint(id), createCategory)
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return rest.OkResponse(ctx, "edit category success", category)
}


func (h *CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := h.Svc.DeleteCategory(uint(id)); err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return rest.OkResponse(ctx, "delete category success", nil)
}


func (h *CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {

	createProdReq := dto.CreateProduct{}

	if err := ctx.BodyParser(&createProdReq); err != nil {
		return rest.BadRequest(ctx, "invalid request body to create product")
	}

	user := h.Svc.Auth.GetCurrentUser(ctx)

	if err := h.Svc.CreateProduct(user.ID, createProdReq); err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return rest.OkResponse(ctx, "create product success", nil)
}


func (h *CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	products, err := h.Svc.GetProducts()

	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}

	return rest.OkResponse(ctx, "get products success", products)
}


func (h *CatalogHandler) GetProductByID(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	product, err := h.Svc.GetProductByID(uint(id))

	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}

	return rest.OkResponse(ctx, "get product by id success", product)
}


func (h *CatalogHandler) GetSellerProducts(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)

	products, err := h.Svc.GetSellerProducts(user.ID)

	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}

	return rest.OkResponse(ctx, "get products success", products)
}


func (h *CatalogHandler) EditProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	createProdReq := dto.CreateProduct{}

	if err := ctx.BodyParser(&createProdReq); err != nil {
		return rest.BadRequest(ctx, "invalid request body to create product")
	}

	product, err := h.Svc.EditProduct(uint(id), createProdReq)

	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return rest.OkResponse(ctx, "edit product success", product)
}


func (h *CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	updateStock := dto.UpdateStock{}

	if err := ctx.BodyParser(&updateStock); err != nil {
		return rest.BadRequest(ctx, "invalid request body to create product")
	}

	product, err := h.Svc.UpdateStock(uint(id), updateStock)

	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return rest.OkResponse(ctx, "update stock success", product)
}


func (h *CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := h.Svc.DeleteProduct(uint(id)); err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return rest.OkResponse(ctx, "delete product success", nil)
}