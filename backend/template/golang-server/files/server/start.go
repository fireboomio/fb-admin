package server

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
	"custom-go/pkg/utils"
	"custom-go/pkg/wgpb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := startServer(); err != nil {
		os.Exit(1)
	}
}

func configureWunderGraphServer() *echo.Echo {
	// 初始化 Echo 实例
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// 配置日志中间件
	e.Use(middleware.Logger())

	// 配置 CORS 中间件
	corsCfg := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
	e.Use(middleware.CORSWithConfig(corsCfg))

	plugins.RegisterProxyHooks(e)
	plugins.RegisterGlobalHooks(e, types.WdgHooksAndServerConfig.Hooks.Global)
	plugins.RegisterAuthHooks(e, types.WdgHooksAndServerConfig.Hooks.Authentication)
	plugins.RegisterUploadsHooks(e, types.WdgHooksAndServerConfig.Hooks.Uploads)

	var internalQueries, internalMutations base.OperationDefinitions
	nodeUrl := utils.GetConfigurationVal(types.WdgGraphConfig.Api.NodeOptions.NodeUrl)
	queryOperations := filterOperationsHooks(types.WdgGraphConfig.Api.Operations, wgpb.OperationType_QUERY)
	if queryLen := len(queryOperations); queryLen > 0 {
		internalQueries = plugins.BuildInternalRequest(e.Logger, nodeUrl, queryOperations)
		plugins.RegisterOperationsHooks(e, queryOperations, types.WdgHooksAndServerConfig.Hooks.Queries)
		e.Logger.Debugf(`Registered (%d) query operations`, queryLen)
	}
	mutationOperations := filterOperationsHooks(types.WdgGraphConfig.Api.Operations, wgpb.OperationType_MUTATION)
	if mutationLen := len(mutationOperations); mutationLen > 0 {
		internalMutations = plugins.BuildInternalRequest(e.Logger, nodeUrl, mutationOperations)
		plugins.RegisterOperationsHooks(e, mutationOperations, types.WdgHooksAndServerConfig.Hooks.Mutations)
		e.Logger.Debugf(`Registered (%d) mutation operations`, mutationLen)
	}
	subscriptionOperations := filterOperationsHooks(types.WdgGraphConfig.Api.Operations, wgpb.OperationType_SUBSCRIPTION)
	if subscriptionLen := len(subscriptionOperations); subscriptionLen > 0 {
		plugins.RegisterOperationsHooks(e, subscriptionOperations, types.WdgHooksAndServerConfig.Hooks.Subscriptions)
		e.Logger.Debugf(`Registered (%d) subscription operations`, subscriptionLen)
	}

	plugins.BuildDefaultInternalClient(internalQueries, internalMutations)
	for _, registeredHook := range base.GetRegisteredHookArr() {
		go registeredHook(e.Logger)
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodGet {
				return next(c)
			}

			var body base.BaseRequestBody
			err := utils.CopyAndBindRequestBody(c.Request(), &body)
			if err != nil {
				return err
			}

			if body.Wg.ClientRequest == nil {
				body.Wg.ClientRequest = &base.ClientRequest{
					Method:     c.Request().Method,
					RequestURI: c.Request().RequestURI,
					Headers:    map[string]string{},
				}
			} else {
				for name, value := range body.Wg.ClientRequest.Headers {
					c.Request().Header.Set(name, value)
				}
			}
			reqId := c.Request().Header.Get("x-request-id")
			internalClient := base.InternalClientFactoryCall(map[string]string{"x-request-id": reqId}, body.Wg.ClientRequest, body.Wg.User)
			internalClient.Queries = internalQueries
			internalClient.Mutations = internalMutations
			brc := &base.BaseRequestContext{
				Context:        c,
				User:           body.Wg.User,
				InternalClient: internalClient,
			}
			return next(brc)
		}
	})

	for _, gqlServer := range types.WdgHooksAndServerConfig.GraphqlServers {
		plugins.RegisterGraphql(e, gqlServer)
	}

	// 健康检查
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	return e
}

func startServer() error {
	// 配置服务器
	wdgServer := configureWunderGraphServer()

	// 启动服务器
	go func() {
		serverListen := types.WdgGraphConfig.Api.ServerOptions.Listen
		address := utils.GetConfigurationVal(serverListen.Host) + ":" + utils.GetConfigurationVal(serverListen.Port)
		if err := wdgServer.Start(address); err != nil {
			panic(err)
		}
	}()

	// 等待终止信号
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// 优雅地关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := wdgServer.Shutdown(ctx); err != nil {
		panic(err)
	}

	return nil
}

func filterOperationsHooks(operations []*types.OperationStruct, operationType wgpb.OperationType) (result []string) {
	for _, operation := range operations {
		if operation.OperationType == operationType && operation.Path != "" {
			result = append(result, operation.Path)
		}
	}
	return
}
