package api

import (
	"fmt"
	"simple-casdoor/object"
	"simple-casdoor/util"
)

const SMS = "sms"

func init() {
	// 通过手机短信验证码登录
	authActionMap[SMS] = &authAction{
		login: func(authForm *AuthForm) (user *object.AdminUser, err error) {
			// 转成E.164格式的电话号码
			dest, ok := util.GetE164Number(authForm.Phone, authForm.CountryCode)
			if !ok {
				return nil, fmt.Errorf("verification:Phone number is invalid in your region %s", authForm.CountryCode)
			}
			// 获取最新一条未被使用的验证码进行验证
			checkResult := object.CheckSignInCode(dest, authForm.Code)

			if len(checkResult) != 0 {
				return nil, fmt.Errorf(checkResult)
			}

			// disable the verification code
			go func() {
				err = object.DisableVerificationCode(dest)
			}()

			// get login user
			user, err = object.GetUserByPhone(authForm.Phone)
			return
		},
	}
}
