import { http } from "@/utils/http";

export type Perm = {
  createdAt: string,
  enabled: number,
  method: string,
  path: string
}


export type PermSyncReq = {
  data: Perm[];
};
export const sendPermission = (data?: PermSyncReq) => {
  return http.request("post", "/operations/System/Perm/CreateMany", { data });
};
/* 将权限管理数据发送到数据库 */

// type Result = {
//   success: boolean;
//   data: Array<any>;
// }
// export const getSubmenu = (data?: number) => {
//   return http.request("get", "/operations/System/Menu/GetChildrenMenus",  {
//     : data
//   });
// }

export const getSubmenu = (data: number) => {
  return http.request("get", `/operations/System/Menu/GetChildrenMenus?pid=${data}`);
}

export const getPerm = (data: number) => {
  return http.request("get", `/operations/System/Menu/GetMenuPerms?menuId=${data}`);
}