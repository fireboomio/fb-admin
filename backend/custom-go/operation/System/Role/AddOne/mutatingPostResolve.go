package AddOne

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

type (
	syncRolesIn  generated.System__Role__SyncRoleInput
	syncRolesOut generated.System__Role__SyncRoleResponseData
)

func MutatingPostResolve(hook *base.HookRequest, body generated.System__Role__AddOneBody) (res generated.System__Role__AddOneBody, err error) {
	hook.Logger().Info("MutatingPostResolve")
	code := body.Response.Data.Data.Code
	object := &generated.System_Role_input_object{Code: code}
	data := make([]*generated.System_Role_input_object, 0)
	data = append(data, object)
	result, err := plugins.ExecuteInternalRequestMutations[syncRolesIn, syncRolesOut](hook.InternalClient, generated.System__Role__SyncRole, syncRolesIn{data})
	if err != nil {
		return nil, err
	}

	hook.Logger().Info(result.Data)
	return body, nil
}
