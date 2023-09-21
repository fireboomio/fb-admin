package object

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type AdminUser struct {
	Name      string    `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedAt time.Time `xorm:"varchar(100) index" json:"created_at"`

	UserId       string `xorm:"varchar(255)" json:"userId"`
	Password     string `xorm:"varchar(100)" json:"password"`
	PasswordSalt string `xorm:"varchar(100)" json:"passwordSalt,omitempty"`
	PasswordType string `xorm:"varchar(100)" json:"passwordType"`
	Phone        string `xorm:"varchar(20) index" json:"phone,omitempty"`
	CountryCode  string `xorm:"varchar(6)" json:"countryCode"`
	WxResp       string `xorm:"varchar(100)" json:"wx_resp,omitempty"`
}

type Userinfo struct {
	Name  string `json:"preferred_username,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func AddUser(user *AdminUser) (int64, error) {
	var err error
	if user.UserId == "" {
		user.UserId = uuid.NewString()
	}

	affected, err := adapter.Engine.Insert(user)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func UpdateUser(user *AdminUser) (int64, error) {
	var err error
	if user.UserId == "" {
		return 0, errors.New("user not exist")
	}

	affected, err := adapter.Engine.Where("user_id=?", user.UserId).Update(user)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func GetUserByPhone(phone string) (*AdminUser, error) {
	if phone == "" {
		return nil, nil
	}

	user := AdminUser{Phone: phone}
	existed, err := adapter.Engine.Get(&user)
	if err != nil {
		return nil, err
	}

	if existed {
		return &user, nil
	} else {
		return nil, nil
	}
}

func GetUserByWxResp(WxResp string) (*AdminUser, error) {
	if WxResp == "" {
		return nil, nil
	}

	user := AdminUser{WxResp: WxResp}
	existed, err := adapter.Engine.Get(&user)
	if err != nil {
		return nil, err
	}

	if existed {
		return &user, nil
	} else {
		return nil, nil
	}

}
