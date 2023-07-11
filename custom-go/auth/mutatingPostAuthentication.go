package auth

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

type (
	getUserRolesIn  generated.Casdoor__GetRolesByIdInput
	getUserRolesOut generated.Casdoor__GetRolesByIdResponseData
)

func MutatingPostAuthentication(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
	// 查询用户角色
	id := hook.User.UserId
	roles, err := plugins.ExecuteInternalRequestQueries[getUserRolesIn, getUserRolesOut](hook.InternalClient, generated.Casdoor__GetRolesById, getUserRolesIn{UserId: id})
	if err != nil {
		return nil, err
	}
	hook.User.Roles = roles.Data
	return &plugins.AuthenticationResponse{User: hook.User, Status: "ok"}, nil
}
