package proxys

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type (
	unBindMenusI generated.System__Role__UnBindMenusInput
	unBindMenusO generated.System__Role__UnBindMenusResponseData

	unBindRoleApisI generated.System__Role__UnBindRoleApisInput
	unBindRoleApisO generated.System__Role__UnBindRoleApisResponseData
)

func init() {
	plugins.AddProxyHook(unBind, nil)
}

func unBind(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	var req = new(Req)
	err := json.Unmarshal(body.Request.OriginBody, &req)
	if err != nil {
		hook.Logger().Errorf(err.Error())
		return nil, err
	}

	// 删除menu-role关联关系
	go func() {
		_, err = plugins.ExecuteInternalRequestMutations[unBindMenusI, unBindMenusO](hook.InternalClient, generated.System__Role__UnBindMenus, unBindMenusI{
			RoleId:  req.RoleId,
			MenuIds: req.MenuIds,
		})
		if err != nil {
			hook.Logger().Errorf(err.Error())
		}
	}()

	//修改角色权限时--飞布需要同步角色和权限 bind role apis
	wg := &sync.WaitGroup{}
	wg.Add(1)
	var apiIds []int64

	// 1.获取 api ids
	// 根据 menuIds 获取 apiIds
	go func() {
		defer wg.Done()
		apis, err := plugins.ExecuteInternalRequestQueries[GetApisByMenusI, GetApisByMenusO](hook.InternalClient, generated.System__Menu__GetApisByMenus, GetApisByMenusI{
			MenuIds: req.MenuIds,
		})
		if err != nil {
			hook.Logger().Errorf("execute get all roles failed: %s", err.Error())
		}

		for _, i := range apis.Data {
			val, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				hook.Logger().Errorf(err.Error())
			}

			apiIds = append(apiIds, val)
		}
	}()

	// 2. 调用 unbind role apis 接口解绑角色和API
	wg.Wait()
	go func() {
		count, err := plugins.ExecuteInternalRequestMutations[unBindRoleApisI, unBindRoleApisO](hook.InternalClient, generated.System__Role__UnBindRoleApis, unBindRoleApisI{
			Apis:     apiIds,
			RoleCode: req.RoleCode,
		})
		if err != nil {
			hook.Logger().Errorf(err.Error())
		}
		hook.Logger().Print(count)
	}()

	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusBadRequest
	}

	body.Response = &base.ClientResponse{StatusCode: statusCode}
	return body.Response, err
}
