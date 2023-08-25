package plugins

import (
	"custom-go/pkg/base"
	"custom-go/pkg/consts"
	"custom-go/pkg/utils"
	"custom-go/pkg/wgpb"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"os"
	"path"
	"path/filepath"
)

type httpProxyHookFunction func(*base.HttpTransportHookRequest, *HttpTransportBody) (*base.ClientResponse, error)

func RegisterProxyHook(hookFunc httpProxyHookFunction, conf ...*HookConfig) {

	callerName := utils.GetCallerName(consts.PROXY)
	apiPrefixPath := "/" + consts.PROXY
	apiPath := path.Join(apiPrefixPath, callerName)

	base.AddEchoRouterFunc(func(e *echo.Echo) {
		e.Logger.Debugf(`Registered hookFunction [%s]`, apiPath)
		e.POST(apiPath, BuildHookFunc(hookFunc))
	})

	base.AddHealthFunc(func(e *echo.Echo, s string, report *base.HealthReport) {
		// 生成 operation 声明文件  proxy/xxx.json
		operation := &wgpb.Operation{
			Name: callerName,
			Path: apiPath,
		}
		if len(conf) > 0 && conf[0] != nil {
			operation.AuthenticationConfig = &wgpb.OperationAuthenticationConfig{AuthRequired: conf[0].AuthRequired}
			operation.AuthorizationConfig = conf[0].AuthorizationConfig
		}

		operationBytes, err := json.Marshal(operation)
		if err != nil {
			e.Logger.Errorf("json marshal failed, err: %v", err.Error())
			return
		}
		err = os.WriteFile(filepath.Join(consts.PROXY, callerName)+consts.JSON_EXT, operationBytes, 0644)
		if err != nil {
			e.Logger.Errorf("write file failed, err: %v", err.Error())
			return
		}

		report.Proxys = append(report.Proxys, callerName)
	})
}
