package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-casdoor/object"
)

type Userinfo struct {
	Sub string `json:"sub"`
}

func GetJwks(c echo.Context) (err error) {
	jwks, err := object.GetJsonWebKeySet()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, jwks)
}

// GetUserInfo ...
//
//	@Title			GetUserInfo
//	@Tag			UserInfo API
//	@Description	Get User Info
//	@Success		200		{object}	object.Userinfo	成功
//	@router			/userinfo [get]
func GetUserInfo(c echo.Context) (err error) {
	user := c.Get("user").(*object.AdminUser)
	userinfo := &Userinfo{
		Sub: user.UserId,
	}
	return c.JSON(200, userinfo)
}

func GetOidcDiscovery(c echo.Context) (err error) {
	host := c.Request().Host
	return c.JSON(200, object.GetOidcDiscovery(host))
}
