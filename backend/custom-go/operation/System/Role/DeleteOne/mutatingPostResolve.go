package DeleteOne

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

type (
	syncDeleteRolesIn  generated.System__Role__SyncDeleteRoleInput
	syncDeleteRolesOut generated.System__Role__SyncDeleteRoleResponseData
)

func MutatingPostResolve(hook *base.HookRequest, body generated.System__Role__DeleteOneBody) (res generated.System__Role__DeleteOneBody, err error) {
	hook.Logger().Info("MutatingPostResolve")

	code := body.Input.Code

	_, err = plugins.ExecuteInternalRequestMutations[syncDeleteRolesIn, syncDeleteRolesOut](hook.InternalClient, generated.System__Role__SyncRole, syncDeleteRolesIn{[]string{code}})
	if err != nil {
		return nil, err
	}

	hook.Logger().Info("delete role: ", code)
	return body, nil
}
