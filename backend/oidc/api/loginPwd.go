package api

import (
	"simple-casdoor/object"
)

const PWD = "password"

func init() {
	// 通过密码登录
	authActionMap[PWD] = &authAction{
		login: func(authForm *AuthForm) (user *object.AdminUser, err error) {
			user, err = object.CheckUserPassword(authForm.Username, authForm.Password)
			if err != nil {
				return nil, err
			}

			return
		},
	}

}
