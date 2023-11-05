package routes

import (
	"github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/api/handlers"
	"github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/api/middlewares"
	"github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/infra/datastores"
	"github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/infra/identity"
	"github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/use_cases/productsuc"
	"github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/use_cases/usermgmtuc"
	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(app *fiber.App) {

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to My Demo Rest API"))
	})

	grp := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)

	grp.Post("/user", handlers.RegisterHandler(registerUseCase))
}

func InitProtectedRoutes(app *fiber.App) {

	grp := app.Group("/api/v1")

	productsDataStore := datastores.NewProductsDataStore()

	createProductUseCase := productsuc.NewCreateProductUseCase(productsDataStore)
	grp.Post("/products", middlewares.NewRequiresRealmRole("admin"),
		handlers.CreateProductHandler(createProductUseCase))

	getProductsUseCase := productsuc.NewGetProductsUseCase(productsDataStore)
	grp.Get("/products", middlewares.NewRequiresRealmRole("viewer"),
		handlers.GetProductsHandler(getProductsUseCase))
}
