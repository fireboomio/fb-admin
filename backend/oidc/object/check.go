package object

import (
	"fmt"
	"simple-casdoor/util"
)

func CheckUserPassword(username string, password string) (user *AdminUser, err error) {
	// 获取user的密码
	user = &AdminUser{}
	existed, err := adapter.Engine.Where("name=?", username).Get(user)
	if err != nil {
		return nil, err
	}

	if !existed {
		return nil, fmt.Errorf("该用户不存在")
	}

	password = util.GenMd5(user.PasswordSalt, password)

	if user.Password == password {
		return user, nil
	}

	return nil, fmt.Errorf("密码错误")
}
