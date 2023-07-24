package server

import (
	"custom-go/global"

	"custom-go/auth"

	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"

	"custom-go/customize"
	_ "custom-go/proxys"
)

func init() {
	types.WdgHooksAndServerConfig = types.WunderGraphHooksAndServerConfig{
		Hooks: types.HooksConfiguration{
			Global: plugins.GlobalConfiguration{
				HttpTransport: plugins.HttpTransportHooks{
					BeforeOriginRequest: global.BeforeOriginRequest,
				},
				WsTransport: plugins.WsTransportHooks{},
			},

			Authentication: plugins.AuthenticationConfiguration{
				MutatingPostAuthentication: auth.MutatingPostAuthentication,
			},

			Queries: base.OperationHooks{},

			Mutations: base.OperationHooks{},

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
