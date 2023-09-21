package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"simple-casdoor/object"
	"simple-casdoor/util"
	"time"
)

// AddUser ...
//
//	@Title			AddUser
//	@Tag			User API
//	@Description	add user
//	@Param			name			body		string			true	"名称"
//	@Param			password		body		string			true	"密码"
//	@Param			passwordType	body		string			false	"密码类型"
//	@Param			phone			body		string			true	"电话号码"
//	@Param			countryCode		body		string			false	"国际区号（默认CN）"
//	@Success		200				{object}	object.Response	成功
//	@router			/add-user [post]
func AddUser(c echo.Context) (err error) {
	var user object.AdminUser

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: err.Error(),
		})
	}

	if user.CountryCode == "" {
		user.CountryCode = "CN"
	}

	if user.PasswordType == "" {
		user.PasswordType = "md5"
	}
	user.PasswordSalt = util.RandomString(12)

	user.Password = util.GenMd5(user.PasswordSalt, user.Password)
	msg := checkUsername(user.Name)
	if msg != "" {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: msg,
		})
	}

	user.CreatedAt = time.Now()
	affected, err := object.AddUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, object.Response{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, object.Response{
		Msg: fmt.Sprintf("affected:%d ", affected),
	})
}

func UpdateUser(c echo.Context) (err error) {
	var user object.AdminUser
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, object.Response{

			Msg: err.Error(),
		})
	}
	if user.CountryCode == "" {
		user.CountryCode = "CN"
	}

	if user.PasswordType == "" {
		user.PasswordType = "md5"
	}
	user.PasswordSalt = util.RandomString(12)
	user.Password = util.GenMd5(user.PasswordSalt, user.Password)

	msg := checkUsername(user.Name)
	if msg != "" {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: msg,
		})
	}

	affected, err := object.UpdateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, object.Response{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, object.Response{
		Msg: fmt.Sprintf("affected:%d ", affected),
	})

}

func checkUsername(username string) string {
	if username == "" {
		return "check:Empty username."
	} else if len(username) > 39 {
		return "check:Username is too long (maximum is 39 characters)."
	}

	exclude, _ := regexp.Compile("^[\u0021-\u007E]+$")
	if !exclude.MatchString(username) {
		return ""
	}
	return ""
}
