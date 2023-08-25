package DeleteMany

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

type (
	syncDeleteRolesIn  generated.System__Role__SyncDeleteRoleInput
	syncDeleteRolesOut generated.System__Role__SyncDeleteRoleResponseData

	getRoleByIdIn  generated.System__Role__GetRoleByIdInput
	getRoleByIdOut generated.System__Role__GetRoleByIdResponseData
)

func MutatingPostResolve(hook *base.HookRequest, body generated.System__Role__DeleteManyBody) (res generated.System__Role__DeleteManyBody, err error) {
	hook.Logger().Info("MutatingPostResolve")

	hook.Logger().Info("MutatingPostResolve")

	ids := body.Input.Ids

	code := make([]string, len(ids))

	for _, id := range ids {
		res, err := plugins.ExecuteInternalRequestQueries[getRoleByIdIn, getRoleByIdOut](hook.InternalClient, generated.System__Role__GetRoleById, getRoleByIdIn{id})
		if err != nil {
			return nil, err
		}

		code = append(code, res.Data.Code)
	}

	_, err = plugins.ExecuteInternalRequestMutations[syncDeleteRolesIn, syncDeleteRolesOut](hook.InternalClient, generated.System__Role__SyncRole, syncDeleteRolesIn{code})
	if err != nil {
		return nil, err
	}

	return body, nil
}
