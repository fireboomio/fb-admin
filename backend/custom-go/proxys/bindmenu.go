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

type Req struct {
	RoleCode string  `json:"roleCode"`
	RoleId   int64   `json:"roleId"`
	MenuIds  []int64 `json:"menuIds"`
}

type (
	bindMenusI      generated.System__Role__BindMenusInput
	bindMenusO      generated.System__Role__BindMenusResponseData
	getRolesI       generated.System__Role__GetManyInput
	getRolesO       generated.System__Role__GetManyResponseData
	GetApisByMenusI generated.System__Menu__GetApisByMenusInput
	GetApisByMenusO generated.System__Menu__GetApisByMenusResponseData
	bindRoleApisI   generated.System__Role__BindRoleApisInput
	bindRoleApisO   generated.System__Role__BindRoleApisResponseData
)

func init() {
	plugins.AddProxyHook(bindMenu, nil)
}

func bindMenu(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	var req = new(Req)
	err := json.Unmarshal(body.Request.OriginBody, &req)
	if err != nil {
		hook.Logger().Errorf(err.Error())
		return nil, err
	}

	// 原入参 - roleCode  roleId  menuIds
	// 目标入参 - []{Menu_id int64  Role_id int64}
	var target []*generated.Main_admin_menu_roleCreateManyInput
	for _, menuId := range req.MenuIds {
		target = append(target, &generated.Main_admin_menu_roleCreateManyInput{
			Menu_id: menuId,
			Role_id: req.RoleId,
		})
	}

	// 创建menu-role关联关系
	go func() {
		_, err = plugins.ExecuteInternalRequestMutations[bindMenusI, bindMenusO](hook.InternalClient, generated.System__Role__BindMenus, bindMenusI{
			Data: target,
		})
		if err != nil {
			hook.Logger().Errorf(err.Error())
		}
	}()

	//修改角色权限时--飞布需要同步角色和权限 bind role apis
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var allRoles []string
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

	// 2.查询 all roles
	go func() {
		defer wg.Done()
		roles, err := plugins.ExecuteInternalRequestQueries[getRolesI, getRolesO](hook.InternalClient, generated.System__Role__GetMany, getRolesI{})
		if err != nil {
			hook.Logger().Errorf("execute get all roles failed: %s", err.Error())
		}

		for _, item := range roles.Data {
			allRoles = append(allRoles, item.Code)
		}
	}()

	// 3. 调用 bind role apis 接口同步角色和权限
	wg.Wait()
	go func() {
		_, err = plugins.ExecuteInternalRequestMutations[bindRoleApisI, bindRoleApisO](hook.InternalClient, generated.System__Role__BindRoleApis, bindRoleApisI{
			AllRoles: allRoles,
			Apis:     apiIds,
			RoleCode: req.RoleCode,
		})
		if err != nil {
			hook.Logger().Errorf(err.Error())
		}
	}()

	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusBadRequest
	}

	body.Response = &base.ClientResponse{StatusCode: statusCode}
	return body.Response, err
}
