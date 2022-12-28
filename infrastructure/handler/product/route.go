package product

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/domain/product"
	"github.com/luryon/go-ecommerce/infrastructure/handler/middle"
	productStorage "github.com/luryon/go-ecommerce/infrastructure/postgres/product"
)

func NewRouter(e *echo.Echo, dbpool *pgxpool.Pool) {
	h := buildHandler(dbpool)

	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(e, h)
}

func buildHandler(dbpool *pgxpool.Pool) handler {
	useCase := product.New(productStorage.New(dbpool))

	return handler{useCase: useCase}
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/admin/products", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}

func publicRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/public/products", middlewares...)

	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}
