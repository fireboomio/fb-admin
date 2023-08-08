package server

import (
	"custom-go/global"
	"github.com/joho/godotenv"

	"custom-go/auth"

	"custom-go/customize"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
	_ "custom-go/proxys"
)

const nodeEnvFilepath = "../.env"

func init() {
	_ = godotenv.Overload(nodeEnvFilepath)

	types.WdgHooksAndServerConfig = types.WunderGraphHooksAndServerConfig{
		Hooks: types.HooksConfiguration{
			Global: plugins.GlobalConfiguration{
				HttpTransport: plugins.HttpTransportHooks{
					BeforeOriginRequest: global.BeforeOriginRequest,

					OnOriginResponse: global.OnOriginResponse,
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
