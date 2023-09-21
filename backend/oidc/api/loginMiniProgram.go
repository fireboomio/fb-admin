package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"simple-casdoor/object"
)

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

const MINIPROGRAM = "MiniProgram"

func init() {
	//通过小程序登录
	authActionMap[MINIPROGRAM] = &authAction{
		login: func(authForm *AuthForm) (user *object.AdminUser, err error) {
			login, err := WXLogin(authForm.AppId, authForm.Secret, authForm.MiniProGramCode)
			if err != nil {
				return nil, err
			}
			var adminUser = &object.AdminUser{
				Name:   "WxUser",
				WxResp: login,
			}

			_, _ = object.AddUser(adminUser)
			user, err = object.GetUserByWxResp(login)
			if err != nil {
				return nil, err
			}

			return
		},
	}

}

// WXLogin 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(appId, secret, code string) (string, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, appId, secret, code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return "", err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return "", errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	encode := GetMD5Encode(wxResp.OpenId + wxResp.SessionKey)
	return encode, nil
}

// 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
