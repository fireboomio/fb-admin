package proxy

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"encoding/json"
	"net/http"
	"sync"
)

func init() {
	plugins.RegisterProxyHook(login, nil)
}

type (
	loginI        generated.Casdoor__LoginInput
	loginO        generated.Casdoor__LoginResponseData
	getUserI      generated.System__User__GetOneInput
	getUserO      generated.System__User__GetOneResponseData
	getRolePermsI generated.System__Perm__GetRolePermsInput
	getRolePermsO generated.System__Perm__GetRolePermsResponseData
	loginRes      struct {
		//Casdoor__LoginResponseData_casdoor_apiLogin
		generated.Casdoor__LoginResponseData_data
		Roles    []string `json:"roles"`
		Perms    []string `json:"perms"`
		Avatar   string   `json:"avatar"`
		UserName string   `json:"username"`
	}
)

func login(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	var loginReq loginI
	loginRes := new(loginRes)

	err := json.Unmarshal(body.Request.OriginBody, &loginReq)
	if err != nil {
		return nil, err
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	var cancel bool

	// 登录获取 token
	go func() {
		defer wg.Done()
		result, _ := plugins.ExecuteInternalRequestMutations[loginI, loginO](hook.InternalClient, generated.Casdoor__Login, loginReq)
		if !result.Data.Success {
			hook.Logger().Errorf(result.Data.Msg)
			cancel = true
		}
		loginRes.Casdoor__LoginResponseData_data = result.Data
	}()

	// 获取roles
	go func() {
		defer wg.Done()
		for !cancel {
			if loginRes.Success {
				userinfo, err := plugins.ExecuteInternalRequestQueries[getUserI, getUserO](hook.InternalClient, generated.System__User__GetOne, getUserI{
					Name:  loginReq.Username,
					Phone: loginReq.Phone,
				})
				if err != nil {
					hook.Logger().Errorf(err.Error())
				}

				loginRes.Roles = userinfo.Data.Roles
				loginRes.Avatar = userinfo.Data.Avatar
				loginRes.UserName = userinfo.Data.Name
				break
			}
		}
	}()

	// 获取perms
	go func() {
		defer wg.Done()
		for !cancel {
			if len(loginRes.Roles) > 0 {
				perms, err := plugins.ExecuteInternalRequestQueries[getRolePermsI, getRolePermsO](hook.InternalClient, generated.System__Perm__GetRolePerms, getRolePermsI{
					Code: loginRes.Roles,
				})
				if err != nil {
					hook.Logger().Errorf(err.Error())
					break
				}

				loginRes.Perms = perms.Data[0].Join.Data

				break
			}
		}
		// 去重
		loginRes.Perms = RemoveNull(loginRes.Perms)
	}()

	// 等待协程执行完毕，返回数据
	wg.Wait()

	statusCode := http.StatusOK

	if !loginRes.Success {
		loginRes.Success = false
		statusCode = http.StatusBadRequest
	}

	bytes, err := json.Marshal(loginRes)
	if err != nil {
		return nil, err
	}

	body.Response = &base.ClientResponse{StatusCode: statusCode}

	body.Response.OriginBody = bytes
	return body.Response, err
}

func RemoveNull(strs []string) []string {
	// 创建一个空的切片，用来存储去重后的字符串
	var result []string

	// 遍历原始字符串切片，把每个元素存入 map 中，元素作为 key，value 始终为 true
	// 当遇到重复元素时，对应的 value 值会被覆盖为 false
	for _, str := range strs {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func RemoveDuplicate(strs []string) []string {
	// 创建一个空的 map[string]bool，用来存储出现过的字符串
	strMap := make(map[string]bool)
	// 创建一个空的切片，用来存储去重后的字符串
	var result []string

	// 遍历原始字符串切片，把每个元素存入 map 中，元素作为 key，value 始终为 true
	// 当遇到重复元素时，对应的 value 值会被覆盖为 false
	for _, str := range strs {
		if !strMap[str] {
			strMap[str] = true
			result = append(result, str)
		}
	}
	return result
}
