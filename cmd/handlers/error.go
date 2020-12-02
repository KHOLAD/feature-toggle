package handlers

import (
	"net/http"

	"github.com/KHOLAD/feature-toggle-api/models"
	"github.com/labstack/echo"
)

// CustomHTTPErrorHandler - custom error handler
func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		key  = "ServerError"
		msg  string
	)

	if he, ok := err.(*models.HTTPError); ok {
		code = he.Code
		key = he.Key
		msg = he.Message
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, models.NewHTTPError(code, key, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
