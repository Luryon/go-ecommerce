package purchaseorder

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/domain/purchaseorder"
	"github.com/luryon/go-ecommerce/infrastructure/handler/middle"
	purchaseorderStorage "github.com/luryon/go-ecommerce/infrastructure/postgres/purchaseorder"
)

func NewRouter(e *echo.Echo, dbpool *pgxpool.Pool) {
	h := buildHandler(dbpool)

	authMiddleware := middle.New()

	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbpool *pgxpool.Pool) handler {
	useCase := purchaseorder.New(purchaseorderStorage.New(dbpool))
	return newHandler(useCase)
}

func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/private/purchase-orders", middlewares...)

	route.POST("", h.Create)
}
