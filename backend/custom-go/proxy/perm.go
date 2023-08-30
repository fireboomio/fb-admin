package proxy

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/wgpb"
	"encoding/json"
	"net/http"
)

func init() {
	plugins.RegisterProxyHook(perms, wgpb.OperationType_MUTATION)
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
	perm, err := plugins.ExecuteInternalRequestQueries[getRolePermsIn, getRolePermsOut](hook.InternalClient, generated.System__Perm__GetRolePerms, getRolePermsIn{Code: roles})
	if err != nil {
		return nil, err
	}

	perms = perm.Data[0].Join.Data

	// 去重
	perms = RemoveNull(perms)

	bytes, err := json.Marshal(perms)
	if err != nil {
		return nil, err
	}
	body.Response = &base.ClientResponse{StatusCode: http.StatusOK}
	body.Response.OriginBody = bytes
	return body.Response, nil
}
