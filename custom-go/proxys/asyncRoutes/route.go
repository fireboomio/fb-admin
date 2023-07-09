package asyncRoutes

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"encoding/json"
	"net/http"
)

type DynamicRoute struct {
	Path     string            `json:"path"`
	Meta     Meta              `json:"meta"`
	Children []RouterChildItem `json:"children"`
}

type RouterChildItem struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Meta Meta   `json:"meta"`
}

type Meta struct {
	Title string   `json:"title"`
	Roles []string `json:"roles,omitempty"`
	Auths []string `json:"auths,omitempty"`
}

type (
	getMenusIn      generated.System__GetMenusInput
	getMenusOut     generated.System__GetMenusResponseData
	getMenuRolesIn  generated.System__Role__GetMenuRolesInput
	getMenuRolesOut generated.System__Role__GetMenuRolesResponseData
	getMenuPermsIn  generated.System__Perm__GetManyInput
	getMenuPermsOut generated.System__Perm__GetManyResponseData
)

func init() {
	plugins.AddProxyHook(route, nil)
}

// 根据前端数据格式需求自定义 proxy 钩子返回用户的动态路由
func route(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	var dynamicRoutes []DynamicRoute

	menus, err := plugins.ExecuteInternalRequestQueries[getMenusIn, getMenusOut](hook.InternalClient, generated.System__GetMenus, getMenusIn{})
	if err != nil {
		return nil, err
	}
	for _, menu := range menus.Data {
		if len(menu.Children) == 0 {
			//一级菜单，需要构造一个child
			child := RouterChildItem{
				Path: menu.Path + "/index",
				Name: menu.Name,
				Meta: Meta{
					Title: menu.Label,
				},
			}
			// 获取 menu roles
			roles, err := plugins.ExecuteInternalRequestQueries[getMenuRolesIn, getMenuRolesOut](hook.InternalClient, generated.System__Role__GetMenuRoles, getMenuRolesIn{
				Title: menu.Label,
			})
			if err != nil {
				return nil, err
			}
			child.Meta.Roles = roles.Data

			// 获取 menu perms
			auths, err := plugins.ExecuteInternalRequestQueries[getMenuPermsIn, getMenuPermsOut](hook.InternalClient, generated.System__Perm__GetMany, getMenuPermsIn{
				Name: menu.Name,
			})
			if err != nil {
				return nil, err
			}
			child.Meta.Auths = auths.Data

			dynamicRoutes = append(dynamicRoutes, DynamicRoute{
				Path: menu.Path,
				Meta: Meta{
					Title: menu.Label,
				},
				Children: []RouterChildItem{child},
			})
		} else {
			// 二级菜单
			var routerChildItems []RouterChildItem
			for _, child := range menu.Children {

				// 获取 menu roles
				roles, err := plugins.ExecuteInternalRequestQueries[getMenuRolesIn, getMenuRolesOut](hook.InternalClient, generated.System__Role__GetMenuRoles, getMenuRolesIn{
					Title: child.Label,
				})
				if err != nil {
					return nil, err
				}

				// 获取 menu perms
				auths, err := plugins.ExecuteInternalRequestQueries[getMenuPermsIn, getMenuPermsOut](hook.InternalClient, generated.System__Perm__GetMany, getMenuPermsIn{
					Name: child.Name,
				})
				if err != nil {
					return nil, err
				}

				routerChildItems = append(routerChildItems, RouterChildItem{
					Path: child.Path,
					Name: child.Name,
					Meta: Meta{
						Title: child.Label,
						Roles: roles.Data,
						Auths: auths.Data,
					},
				})
			}

			dynamicRoutes = append(dynamicRoutes, DynamicRoute{
				Path: menu.Path,
				Meta: Meta{
					Title: menu.Label,
				},
				Children: routerChildItems,
			})
		}
	}

	body.Response = &base.ClientResponse{StatusCode: http.StatusOK}
	bytes, err := json.Marshal(dynamicRoutes)
	if err != nil {
		return nil, err
	}

	body.Response.OriginBody = bytes
	return body.Response, nil
}
