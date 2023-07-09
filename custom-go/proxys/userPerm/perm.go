package userPerm

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"encoding/json"
	"net/http"
)

func init() {
	plugins.AddProxyHook(perms, nil)
}

type (
	getRolePermsIn  generated.System__Perm__GetRolePermsInput
	getRolePermsOut generated.System__Perm__GetRolePermsResponseData
)

// 根据前端数据格式需求自定义 proxy 钩子返回用户的权限
func perms(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	var perms []string
	var roles []string

	err := json.Unmarshal(body.Request.OriginBody, &roles)
	if err != nil {
		return nil, err
	}

	// 查询用户所属角色（可能是多个角色）的权限集合
	for _, role := range roles {
		perm, err := plugins.ExecuteInternalRequestQueries[getRolePermsIn, getRolePermsOut](hook.InternalClient, generated.System__Perm__GetRolePerms, getRolePermsIn{
			Code: role,
		})
		if err != nil {
			return nil, err
		}

		perms = append(perms, perm.Data...)
	}

	// 去重
	perms = RemoveDuplicate(perms)

	bytes, err := json.Marshal(perms)
	if err != nil {
		return nil, err
	}
	body.Response = &base.ClientResponse{StatusCode: http.StatusOK}
	body.Response.OriginBody = bytes
	return body.Response, nil
}

func RemoveDuplicate(strs []string) []string {
	// 创建一个空的 map[string]bool，用来存储出现过的字符串
	strMap := make(map[string]bool)
	// 创建一个空的切片，用来存储去重后的字符串
	result := []string{}

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
