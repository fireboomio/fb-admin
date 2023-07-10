import { defineStore } from "pinia";
import { store } from "@/store";
import { userType } from "./types";
import { routerArrays } from "@/layout/types";
import { router, resetRouter } from "@/router";
import { storageSession } from "@pureadmin/utils";
import { LoginResult, getLogin, refreshTokenApi } from "@/api/user";
import { RefreshTokenResult } from "@/api/user";
import { useMultiTagsStoreHook } from "@/store/modules/multiTags";
import {
  type DataInfo,
  setToken,
  removeToken,
  sessionKey,
  formatToken
} from "@/utils/auth";
import axios from "axios";
import { ElMessage } from "element-plus";
// 在本地存储用户信息
export const useUserStore = defineStore({
  id: "pure-user",
  state: (): userType => ({
    // 用户名
    username:
      storageSession().getItem<DataInfo<number>>(sessionKey)?.username ?? "",
    // 页面级别权限
    roles: storageSession().getItem<DataInfo<number>>(sessionKey)?.roles ?? [],
    permissions: [],
    avatar:
      storageSession().getItem<DataInfo<number>>(sessionKey)?.avatar ?? "", // 头像
    // 判断登录页面显示哪个组件（password：登录（默认）、sms：手机登录)
    loginType: "password"
  }),
  actions: {
    /** 存储用户名 */
    SET_USERNAME(username: string) {
      this.username = username;
    },
    /** 存储角色 */
    SET_ROLES(roles: Array<string>) {
      this.roles = roles;
    },
    /** 存储权限 */
    SET_PERMISSIONS(permissions: Array<string>) {
      this.permissions = permissions;
    },
    /** 存储用户头像 */
    SET_USERAVATAR(avatar: string) {
      this.avatar = avatar;
    },
    /** 存储登录页面显示哪个组件 */
    SET_LOGINTYPE(value: string) {
      this.loginType = value;
    },
    /** 登入 */
    async loginByUsername(data) {
      const phone = data.username;
      const name = data.username;
      return new Promise<LoginResult>((resolve, reject) => {
        getLogin(data)
          .then(async res => {
            // 获取到token
            if (res.data.data.success) {
              // 获取用户信息
              const userInfo = await axios.get("/operations/Casdoor/GetUser", {
                headers: {
                  Authorization: formatToken(res.data.data.data.access_token)
                },
                params: {
                  phone,
                  name
                }
              });
              const dataInfo: DataInfo<number> = {
                accessToken: userInfo.data.data.token.data.access_token,
                expires: userInfo.data.data.token.data.expires_in,
                refreshToken: userInfo.data.data.token.data.refresh_token,
                username: userInfo.data.data.token.data.username,
                roles: userInfo.data.data.user.roles,
                avatar: userInfo.data.data.user.avatar
              };
              setToken(dataInfo);
              resolve(res);
            } else {
              ElMessage.info(res.data.data.msg);
            }
          })
          .catch(error => {
            reject(error);
          });
      });
    },
    /** 前端登出（不调用接口） */
    logOut() {
      this.username = "";
      this.roles = [];
      this.permissions = [];
      removeToken();
      useMultiTagsStoreHook().handleTags("equal", [...routerArrays]);
      resetRouter();
      router.push("/login");
    },
    /** 刷新`token` */
    async handRefreshToken(data) {
      return new Promise<RefreshTokenResult>((resolve, reject) => {
        refreshTokenApi(data)
          .then(res => {
            if (res) {
              setToken(res.data.data);
              resolve(res);
            }
          })
          .catch(error => {
            reject(error);
          });
      });
    }
  }
});

export function useUserStoreHook() {
  return useUserStore(store);
}
