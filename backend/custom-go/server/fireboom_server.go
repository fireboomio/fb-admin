package server

import (
	authentication "custom-go/authentication"
	"github.com/joho/godotenv"

	"custom-go/generated"

	operation_System__Role__AddOne "custom-go/operation/System/Role/AddOne"

	operation_System__Role__DeleteMany "custom-go/operation/System/Role/DeleteMany"

	operation_System__Role__DeleteOne "custom-go/operation/System/Role/DeleteOne"

	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
)

const nodeEnvFilepath = "../.env"

func init() {
	_ = godotenv.Overload(nodeEnvFilepath)

	types.WdgHooksAndServerConfig = types.WunderGraphHooksAndServerConfig{
		Hooks: types.HooksConfiguration{
			Global: plugins.GlobalConfiguration{},
			Authentication: plugins.AuthenticationConfiguration{
				MutatingPostAuthentication: authentication.MutatingPostAuthentication,
			},
			Queries: base.OperationHooks{},
			Mutations: base.OperationHooks{
				"System/Role/AddOne": {
					MutatingPostResolve: plugins.ConvertBodyFunc[generated.System__Role__AddOneInternalInput, generated.System__Role__AddOneResponseData](operation_System__Role__AddOne.MutatingPostResolve),
				},
				"System/Role/DeleteMany": {
					MutatingPostResolve: plugins.ConvertBodyFunc[generated.System__Role__DeleteManyInternalInput, generated.System__Role__DeleteManyResponseData](operation_System__Role__DeleteMany.MutatingPostResolve),
				},
				"System/Role/DeleteOne": {
					MutatingPostResolve: plugins.ConvertBodyFunc[generated.System__Role__DeleteOneInternalInput, generated.System__Role__DeleteOneResponseData](operation_System__Role__DeleteOne.MutatingPostResolve),
				},
			},
			Uploads: map[string]plugins.UploadHooks{},
		},
	}
}
