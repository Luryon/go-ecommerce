package login

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/domain/login"
	"github.com/luryon/go-ecommerce/domain/user"
	storageUser "github.com/luryon/go-ecommerce/infrastructure/postgres/user"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCaseUser := user.New(storageUser.New(dbPool))
	useCase := login.New(useCaseUser)
	return newHandler(useCase)
}

func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/login")

	route.POST("", h.Login)
}
