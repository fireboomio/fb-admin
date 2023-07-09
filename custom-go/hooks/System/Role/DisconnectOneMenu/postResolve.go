package DisconnectOneMenu

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

// PostResolve 删除菜单和角色关联关系时触发额外操作: 解绑角色和菜单下的所有api
func PostResolve(hook *base.HookRequest, body generated.System__Role__DisconnectOneMenuBody) (res generated.System__Role__DisconnectOneMenuBody, err error) {
	hook.Logger().Info("PostResolve")
	// 1.获取菜单的 api id集合
	perms, err := plugins.ExecuteInternalRequestQueries[getMenuPermsIn, getMenuPermsOut](hook.InternalClient, generated.System__Menu__GetMenuPerms, getMenuPermsIn{
		MenuId: body.Input.Menu_id,
	})
	if err != nil {
		return nil, err
	}
	// 取出 id 集合
	var menuApis = make([]int64, 0)
	for _, perm := range perms.Data {
		permVal, err := strconv.ParseInt(perm.Id, 10, 64)
		if err != nil {
			hook.Logger().Errorf("execute get menu roles failed: %v", err)
			continue
		}
		menuApis = append(menuApis, permVal)
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

	// 对 roleApis 和 menuApis作差集
	ids := difference(roleApis.Data, menuApis)

	// 4.将 api 与 role进行解绑
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

func difference(arr1, arr2 []int64) []int64 {
	// 将arr1转换成map
	set := make(map[int64]bool)
	for _, num := range arr1 {
		set[num] = true
	}

	// 遍历arr2，删除在arr1中出现过的元素
	for _, num := range arr2 {
		delete(set, num)
	}

	// 将map的key提取到数组中，得到差集结果
	var result []int64
	for num := range set {
		result = append(result, num)
	}

	return result
}
