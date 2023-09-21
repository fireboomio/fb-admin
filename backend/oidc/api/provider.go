package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-casdoor/object"
)

// UpdateProvider
//
//	@Title			UpdateProvider
//	@Tag			Provider API
//	@Description	update provider
//	@Param			clientId		body		string			true	"clientId"
//	@Param			clientSecret	body		string			true	"clientSecret"
//	@Param			signName		body		string			true	"签名"
//	@Param			templateCode	body		string			true	"模板代码"
//	@Success		200				{object}	object.Response	成功
//	@router			/update-provider [post]
func UpdateProvider(c echo.Context) (err error) {
	var provider object.AdminProvider

	if err := c.Bind(&provider); err != nil {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: err.Error(),
		})
	}

	affected, err := object.UpdateProvider(&provider)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, object.Response{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, object.Response{
		Msg: fmt.Sprintf("affected:%d ", affected),
	})
}
