package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/infrastructure/handler/invoice"
	"github.com/luryon/go-ecommerce/infrastructure/handler/login"
	"github.com/luryon/go-ecommerce/infrastructure/handler/paypal"
	"github.com/luryon/go-ecommerce/infrastructure/handler/product"
	"github.com/luryon/go-ecommerce/infrastructure/handler/purchaseorder"
	"github.com/luryon/go-ecommerce/infrastructure/handler/user"
	"net/http"
	"time"
)

func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	health(e)

	// A

	// B

	// C

	// I
	invoice.NewRouter(e, dbPool)

	// L
	login.NewRouter(e, dbPool)

	// P
	paypal.NewRouter(e, dbPool)
	product.NewRouter(e, dbPool)
	purchaseorder.NewRouter(e, dbPool)

	// U
	user.NewRouter(e, dbPool)
}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello World!",
				"serviec_name": "",
			})
	})
}
