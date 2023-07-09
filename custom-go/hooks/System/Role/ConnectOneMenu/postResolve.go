package ConnectOneMenu

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"strconv"
)

type (
	getMenuPermsIn  generated.System__Menu__GetMenuPermsInput
	getMenuPermsOut generated.System__Menu__GetMenuPermsResponseData
	getRolesIn      generated.System__Role__GetManyInput
	getRolesOut     generated.System__Role__GetManyResponseData
	getOneRoleIn    generated.System__Role__GetOneInput
	getOneRoleOut   generated.System__Role__GetOneResponseData
	bindRoleAPisIn  generated.System__Role__BindRoleApisInternalInput
	bindRoleAPisOut generated.System__Role__BindRoleApisResponseData
	getRoleBindIn   generated.System__Role__GetRoleBindApiIdsInput
	getRoleBindOut  generated.System__Role__GetRoleBindApiIdsResponseData
)

// PostResolve 创建菜单和角色关联关系时触发额外操作: 绑定角色和菜单下的所有api
func PostResolve(hook *base.HookRequest, body generated.System__Role__ConnectOneMenuBody) (res generated.System__Role__ConnectOneMenuBody, err error) {
	//hook.Logger().Info("PostResolve")
	//
	//wg := &sync.WaitGroup{}
	//wg.Add(5)
	//
	//menuApis := make(chan []int64)
	//roleApis := make(chan []int64)
	//allRoles := make(chan []string)
	//roleCode := make(chan string)
	//
	//defer func() {
	//	// 关闭 channel
	//	close(roleApis)
	//	close(menuApis)
	//	close(allRoles)
	//	close(roleCode)
	//}()
	//
	//// 1.获取菜单的 api id集合
	//go func() {
	//	defer wg.Done()
	//	perms, err := plugins.ExecuteInternalRequestQueries[getMenuPermsIn, getMenuPermsOut](hook.InternalClient, generated.System__Menu__GetMenuPerms, getMenuPermsIn{
	//		MenuId: body.Input.Menu_id,
	//	})
	//	if err != nil {
	//		hook.Logger().Errorf("execute get menu roles failed: %v", err)
	//	}
	//	// 取出 id 集合
	//	var ids = make([]int64, 0)
	//	for _, perm := range perms.Data {
	//		permVal, err := strconv.ParseInt(perm.Id, 10, 64)
	//		if err != nil {
	//			hook.Logger().Errorf("execute get menu roles failed: %v", err)
	//			continue
	//		}
	//		ids = append(ids, permVal)
	//	}
	//	menuApis <- ids
	//}()
	//
	//// 2.获取当前绑定的role code
	//go func() {
	//	defer wg.Done()
	//	role, err := plugins.ExecuteInternalRequestQueries[getOneRoleIn, getOneRoleOut](hook.InternalClient, generated.System__Role__GetOne, getOneRoleIn{
	//		Id: body.Input.Role_id,
	//	})
	//	if err != nil {
	//		hook.Logger().Errorf("execute get role code failed: %v", err)
	//	}
	//	roleCode <- role.Data.Code
	//}()
	//
	//// 3.获取 all roles
	//go func() {
	//	defer wg.Done()
	//	roles, err := plugins.ExecuteInternalRequestQueries[getRolesIn, getRolesOut](hook.InternalClient, generated.System__Role__GetMany, getRolesIn{})
	//	if err != nil {
	//		hook.Logger().Errorf("execute get all roles failed: %v", err)
	//	}
	//	// 取出roles code集合
	//	var roleCodes = make([]string, 0)
	//	for _, role := range roles.Data {
	//		roleCodes = append(roleCodes, role.Code)
	//	}
	//	allRoles <- roleCodes
	//}()
	//
	//// 4.获取当前角色已绑定的api ids
	//go func() {
	//	defer wg.Done()
	//	ids, err := plugins.ExecuteInternalRequestQueries[getRoleBindIn, getRoleBindOut](hook.InternalClient, generated.System__Role__GetRoleBindApiIds, getRoleBindIn{
	//		Code: <-roleCode,
	//	})
	//	if err != nil {
	//		hook.Logger().Errorf("execute get role bind apis failed: %v", err)
	//	}
	//	roleApis <- ids.Data
	//}()
	//
	//// 5.接收协程
	//go func() {
	//	defer wg.Done()
	//	apiIds := append(<-menuApis, <-roleApis...)
	//
	//	// 将 api 与 role进行绑定
	//	_, err = plugins.ExecuteInternalRequestQueries[bindRoleAPisIn, bindRoleAPisOut](hook.InternalClient, generated.System__Role__BindRoleApis, bindRoleAPisIn{
	//		AllRoles: <-allRoles,
	//		RoleCode: <-roleCode,
	//		Apis:     apiIds,
	//	})
	//	if err != nil {
	//		hook.Logger().Errorf("execute bind role apis failed: %v", err)
	//	}
	//}()
	//
	//wg.Wait()
	//
	//return body, nil
	hook.Logger().Info("PostResolve")
	// 1.获取菜单的 api id集合
	perms, err := plugins.ExecuteInternalRequestQueries[getMenuPermsIn, getMenuPermsOut](hook.InternalClient, generated.System__Menu__GetMenuPerms, getMenuPermsIn{
		MenuId: body.Input.Menu_id,
	})
	if err != nil {
		return nil, err
	}
	// 取出 id 集合
	var ids = make([]int64, 0)
	for _, perm := range perms.Data {
		permVal, err := strconv.ParseInt(perm.Id, 10, 64)
		if err != nil {
			hook.Logger().Errorf("execute get menu roles failed: %v", err)
			continue
		}
		ids = append(ids, permVal)
	}

	// 2.获取 all roles
	roles, err := plugins.ExecuteInternalRequestQueries[getRolesIn, getRolesOut](hook.InternalClient, generated.System__Role__GetMany, getRolesIn{})
	if err != nil {
		return nil, err
	}
	// 取出roles code集合
	var roleCodes = make([]string, 0)
	for _, role := range roles.Data {
		roleCodes = append(roleCodes, role.Code)
	}

	// 3.获取当前绑定的role code
	role, err := plugins.ExecuteInternalRequestQueries[getOneRoleIn, getOneRoleOut](hook.InternalClient, generated.System__Role__GetOne, getOneRoleIn{
		Id: body.Input.Role_id,
	})
	if err != nil {
		return nil, err
	}

	// 4.获取当前角色已绑定的api ids
	roleApis, err := plugins.ExecuteInternalRequestQueries[getRoleBindIn, getRoleBindOut](hook.InternalClient, generated.System__Role__GetRoleBindApiIds, getRoleBindIn{
		Code: role.Data.Code,
	})
	if err != nil {
		hook.Logger().Errorf("execute get role bind apis failed: %v", err)
	}
	ids = append(ids, roleApis.Data...)

	// 5.将 api 与 role进行绑定
	_, err = plugins.ExecuteInternalRequestMutations[bindRoleAPisIn, bindRoleAPisOut](hook.InternalClient, generated.System__Role__BindRoleApis, bindRoleAPisIn{
		AllRoles: roleCodes,
		RoleCode: role.Data.Code,
		Apis:     ids,
	})
	if err != nil {
		return nil, err
	}
	return body, nil
}
