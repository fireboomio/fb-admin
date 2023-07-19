import axios from "@/utils/http";

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
  return axios.post<any>("/operations/System/Perm/CreateMany", {
    data
  });
};

export const getSubmenu = (data: number) => {
  return axios.get<any>(`/operations/System/Menu/GetChildrenMenus?pid=${data}`);
};

export const getPerm = (data: number) => {
  return axios.get<any>(`/operations/System/Menu/GetMenuPerms?menuId=${data}`);
};

export const getPagePerm = (data: number) => {
  return axios.get<any>(`/operations/System/Perm/GetMany?menuId=${data}`);
};

export const getMenuRoles = (data: string) => {
  return axios.get<any>(`/operations/System/Role/GetMenuRoles?title=${data}`);
};

export const getMenuPerms = (data: string) => {
  return axios.get<any>(`/operations/System/Perm/GetMany?name=${data}`);
};

export const getBindAPI = () => {
  return axios.get<any>("/operations/System/Perm/GetBindPerms");
};

// 根据角色获取权限列表
export const getRolePerms = (data: string[]) => {
  return axios.post<string[]>(`/proxy/perm`, {
    data
  });
};

export const deleteRoles = (ids: number[] | number) => {
  return axios.post("/operations/System/Role/DeleteMany", {
    data: { ids }
  });
};

export const roleMenuTreeselect = (id: number) => {
  return axios.post<any>("/proxy/menuTree", {
    data: id
  });
}

export const updateRolePermAdd = (rolecode: number, roleId: number, menuIds: number[]) => {
  return axios.post<any>("/proxy/bindmenu", {
    data: {
      rolecode,
      roleId,
      menuIds
    }
  })
};

export const updateRolePermRemove = (rolecode: number, roleId: number, menuIds: number[]) => {
  return axios.post<any>("/proxy/unBindMenu", {
    data: {
      rolecode,
      roleId,
      menuIds
    }
  })
};