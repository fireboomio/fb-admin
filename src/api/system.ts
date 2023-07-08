import { http } from "@/utils/http";

export type Perm = {
  id: number;
  createdAt: string;
  enabled: number;
  method: string;
  path: string;
};

export type PermSyncReq = {
  data: Perm[];
};
export const sendPermission = (data?: PermSyncReq) => {
  return http.request("post", "/operations/System/Perm/CreateMany", { data });
};

export const getSubmenu = (data: number) => {
  return http.request(
    "get",
    `/operations/System/Menu/GetChildrenMenus?pid=${data}`
  );
};

export const getPerm = (data: number) => {
  return http.request(
    "get",
    `/operations/System/Menu/GetMenuPerms?menuId=${data}`
  );
};

export const getPagePerm = (data: number) => {
  return http.request("get", `/operations/System/Perm/GetMany?menuId=${data}`);
};

export const getMenuRoles = (data: string) => {
  return http.request(
    "get",
    `/operations/System/Role/GetMenuRoles?title=${data}`
  );
};

export const getMenuPerms = (data: string) => {
  return http.request("get", `/operations/System/Perm/GetMany?name=${data}`);
};

// 获取动态路由api
export const getDynamicRoute = () => {
  return http.request<[]>("get", `/proxy/asyncRoutes/route`);
};

// 根据角色获取权限列表
export const getRolePerms = (data: string[]) => {
  return http.request<string[]>("post", `/proxy/userPerm/perm`, {
    data
  });
};

export const getBindAPI = () => {
  return http.request("get", "/operations/System/Perm/GetBindPerms");
};
