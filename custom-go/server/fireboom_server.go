package server

import (
	"custom-go/auth"

	"custom-go/generated"

	hooks_System__Role__ConnectOneMenu "custom-go/hooks/System/Role/ConnectOneMenu"

	hooks_System__Role__DisconnectOneMenu "custom-go/hooks/System/Role/DisconnectOneMenu"

	"custom-go/customize"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
	_ "custom-go/proxys/asyncRoutes"
	_ "custom-go/proxys/userPerm"
)

func init() {
	types.WdgHooksAndServerConfig = types.WunderGraphHooksAndServerConfig{
		Hooks: types.HooksConfiguration{
			Global: plugins.GlobalConfiguration{
				HttpTransport: plugins.HttpTransportHooks{},
				WsTransport:   plugins.WsTransportHooks{},
			},

			Authentication: plugins.AuthenticationConfiguration{
				MutatingPostAuthentication: auth.MutatingPostAuthentication,
			},

			Queries: base.OperationHooks{},

			Mutations: base.OperationHooks{
				"System/Role/ConnectOneMenu": {
					PostResolve: plugins.ConvertBodyFunc[generated.System__Role__ConnectOneMenuInternalInput, generated.System__Role__ConnectOneMenuResponseData](hooks_System__Role__ConnectOneMenu.PostResolve),
				},
				"System/Role/DisconnectOneMenu": {
					PostResolve: plugins.ConvertBodyFunc[generated.System__Role__DisconnectOneMenuInternalInput, generated.System__Role__DisconnectOneMenuResponseData](hooks_System__Role__DisconnectOneMenu.PostResolve),
				},
			},

			Subscriptions: base.OperationHooks{},

			Uploads: map[string]plugins.UploadHooks{},
		},
		GraphqlServers: []plugins.GraphQLServerConfig{
			{
				ApiNamespace:          "statistics",
				ServerName:            "statistics",
				EnableGraphQLEndpoint: true,
				Schema:                customize.Statistics_schema,
			},
		},
	}
}
