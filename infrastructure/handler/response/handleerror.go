package response

import (
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/model"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	//Error personalizado
	e, ok := err.(*model.Error)
	if ok {
		_ = c.JSON(getReponseError(e))
		return
	}

	//check echo error
	if echoErr, ok := err.(*echo.HTTPError); ok {
		msg, ok := echoErr.Message.(string)
		if !ok {
			msg = "¡Ups! Algo inesperado ocurrio"
		}

		_ = c.JSON(echoErr.Code, model.MessageResponse{
			Errors: model.Responses{
				{Code: UnexpectedError, Message: msg},
			},
		})

		return
	}
}

func getReponseError(err *model.Error) (int, model.MessageResponse) {
	outputStatus := 0
	outputResponse := model.MessageResponse{}
	if !err.HasCode() {
		err.Code = UnexpectedError
	}

	if err.HasData() {
		outputResponse.Data = err.Data
	}

	if !err.HasStatus() {
		err.StatusHttp = http.StatusInternalServerError
	}

	outputStatus = err.StatusHttp
	outputResponse.Errors = model.Responses{model.Response{
		Code:    err.Code,
		Message: err.APIMessage,
	}}

	return outputStatus, outputResponse
}
