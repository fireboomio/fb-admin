package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-casdoor/object"
	"simple-casdoor/util"
)

type VerificationForm struct {
	Dest        string `form:"dest"`
	CountryCode string `form:"countryCode"`
}

// SendVerificationCode ...
//
//	@Title			SendVerificationCode
//	@Tag			Verification API
//	@Description	发送验证码
//	@Param			dest		body		string			true	"发送手机号"
//	@Param			countryCode	body		string			false	"国际区号（默认CN）"
//	@Success		200			{object}	object.Response	成功
//	@router			/send-verification-code [post]
func SendVerificationCode(c echo.Context) (err error) {
	var vForm VerificationForm
	if err := c.Bind(&vForm); err != nil {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: err.Error(),
		})
	}
	if vForm.CountryCode == "" {
		vForm.CountryCode = "CN"
	}

	// 通过号码获取用户
	var user *object.AdminUser
	if user, err := object.GetUserByPhone(vForm.Dest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else if user == nil {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: fmt.Sprintf("verification:the user does not exist, please sign up first"),
		})
	}

	// 获取短信提供商
	provider, err := object.GetProvider("fireboom/provider_sms")
	if err != nil {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: fmt.Sprintf("verification:Phone number is invalid in your region %s", vForm.CountryCode),
		})

	}

	if phone, ok := util.GetE164Number(vForm.Dest, vForm.CountryCode); !ok {
		return c.JSON(http.StatusBadRequest, object.Response{
			Msg: fmt.Sprintf("verification:Phone number is invalid in your region %s", vForm.CountryCode),
		})
	} else {
		remoteAddr := util.GetIPFromRequest(c.Request())
		err := object.SendVerificationCodeToPhone(user, provider, remoteAddr, phone)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, object.Response{
				Msg: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, object.Response{
			Msg: "ok",
		})
	}
}
