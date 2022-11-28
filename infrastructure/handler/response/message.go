package response

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/luryon/go-ecommerce/model"
	"net/http"
)

const (
	BindFailed = "bind_failed"

	Ok              = "ok"
	RecordCreated   = "record_created"
	RecordUpdated   = "record_updated"
	RecordDeleted   = "record_deleted"
	UnexpectedError = "unexpected_error"
	AuthError       = "authorization_error"
)

type API struct{}

func New() API {
	return API{}
}

func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{Data: data,
		Messages: model.Responses{{
			Code:    Ok,
			Message: "¡Listo!",
		}}}
}

func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{Data: data,
		Messages: model.Responses{{
			Code:    RecordCreated,
			Message: "¡Listo!",
		}}}
}

func (a API) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{Data: data,
		Messages: model.Responses{{
			Code:    RecordUpdated,
			Message: "¡Listo!",
		}}}
}

func (a API) Deleted(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{Data: data,
		Messages: model.Responses{{
			Code:    RecordDeleted,
			Message: "¡Listo!",
		}}}
}

func (a API) BindFailed(err error) error {
	e := model.NewError()
	e.Err = err
	e.Code = BindFailed
	e.StatusHttp = http.StatusBadRequest
	e.Who = "c.Bind()"

	log.Warnf("%s", e.Error())
	return &e
}

func (a API) Error(c echo.Context, who string, err error) *model.Error {
	e := model.NewError()
	e.Err = err
	e.APIMessage = "¡Uy! metimos la pata, disculpanos lo solucionaremos"
	e.Code = UnexpectedError
	e.StatusHttp = http.StatusInternalServerError
	e.Who = who

	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		log.Errorf("cannot get/parse uuid from userID")
	}
	e.UserID = userID.String()

	log.Errorf("%s", e.Error())
	return &e
}
