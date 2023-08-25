package proxy

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"encoding/json"
	"net/http"
	"sync"
)

func init() {
	plugins.RegisterProxyHook(tree, nil)
}

type MenuTreeResp struct {
	Code        int     `json:"code"`
	Menus       []Menu  `json:"menus"`
	Msg         string  `json:"msg"`
	CheckedKeys []int64 `json:"checkedKeys"`
	AllKeys     []int64 `json:"allKeys"`
}

type Menu struct {
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Children []Menu `json:"children,omitempty"`
}

type MenuItem struct {
	Id       int64      `json:"id"`
	IsBottom int64      `json:"is_bottom"`
	Label    string     `json:"label"`
	Children []MenuItem `json:"-"`
}

type (
	getRoleMenuIdI  generated.System__Role__GetRoleMenuIdInput
	getRoleMenuIdO  generated.System__Role__GetRoleMenuIdResponseData
	getMenuI        generated.System__Menu__GetManyInput
	getMenuO        generated.System__Menu__GetManyResponseData
	getMenuByLevelI generated.System__Menu__GetMenuByLevelOrPidInput
	getMenuByLevelO generated.System__Menu__GetMenuByLevelOrPidResponseData
)

// 根据角色获取树形的菜单权限数据
func tree(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	var role int64

	wg := &sync.WaitGroup{}
	wg.Add(2)

	err := json.Unmarshal(body.Request.OriginBody, &role)
	if err != nil {
		return nil, err
	}

	var resp = *new(MenuTreeResp)

	// 1.获取角色已分配的菜单权限
	go func() {
		defer wg.Done()
		menuIds, err := plugins.ExecuteInternalRequestQueries[getRoleMenuIdI, getRoleMenuIdO](hook.InternalClient, generated.System__Role__GetRoleMenuId, getRoleMenuIdI{
			RoleId: role,
		})
		if err != nil {
			hook.Logger().Errorf("execute get role menu ids failed: %s", err.Error())
		}

		resp.CheckedKeys = menuIds.Data
	}()

	// 2.获取所有菜单
	go func() {
		defer wg.Done()
		res, err := plugins.ExecuteInternalRequestQueries[getMenuI, getMenuO](hook.InternalClient, generated.System__Menu__GetMany, getMenuI{})
		if err != nil {
			hook.Logger().Errorf("execute get role menu ids failed: %s", err.Error())
		}

		var menus = res.Data
		var menuRes = make([]Menu, 0)
		// 备忘录 -- 记录已经遍历过的 menu，避免出现重复数据
		var memo = make(map[int64]int)

		for _, menuItem := range menus {
			resp.AllKeys = append(resp.AllKeys, menuItem.Id)

			menu := Menu{
				Id:    menuItem.Id,
				Label: menuItem.Label,
			}

			// 构建父菜单的children
			if menuItem.Is_bottom == 0 && !ContainsKey(menuItem.Id, memo) {
				children, err := ConstructMenuChildren(hook, menuItem.Id, memo)
				if err != nil {
					hook.Logger().Errorf("execute get role menu ids failed: %s", err.Error())
				}

				menu.Children = children
				menuRes = append(menuRes, menu)
			}
		}

		resp.Menus = menuRes
	}()

	resp.Msg = "操作成功"
	wg.Wait()

	bytes, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	body.Response = &base.ClientResponse{StatusCode: http.StatusOK}
	body.Response.OriginBody = bytes

	return body.Response, nil
}

func ConstructMenuChildren(hook *base.HttpTransportHookRequest, parentId int64, memo map[int64]int) ([]Menu, error) {
	menuItems, err := plugins.ExecuteInternalRequestQueries[getMenuByLevelI, getMenuByLevelO](hook.InternalClient, generated.System__Menu__GetMenuByLevelOrPid, getMenuByLevelI{
		Pid: parentId,
	})
	if err != nil {
		hook.Logger().Errorf("execute get role menu ids failed: %s", err.Error())
		return nil, err
	}

	children := []Menu{}

	for _, menuItem := range menuItems.Data {
		menu := Menu{
			Id:    menuItem.Id,
			Label: menuItem.Label,
		}

		if menuItem.Is_bottom == 0 {
			memo[menuItem.Id] = 1
			grandChildren, _ := ConstructMenuChildren(hook, menuItem.Id, memo)
			if err != nil {
				hook.Logger().Errorf("execute get role menu ids failed: %s", err.Error())
				return nil, err
			}

			menu.Children = grandChildren
		}

		children = append(children, menu)
	}

	return children, nil
}

func ContainsKey(key int64, m map[int64]int) bool {
	_, ok := m[key]
	return ok
}
