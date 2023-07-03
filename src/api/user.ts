import { http } from "@/utils/http";

export type LoginResult = {
  data: {
    data: {
      success: boolean;
      data: {
        /** `token` */
        access_token: string;
        /** 用于调用刷新`accessToken`的接口时所需的`token` */
        refresh_token: string;
      };
      msg: string;
    };
  };
};

export type UserResult = {
  data: {
    token: {
      success: boolean;
      data: {
        /** `token` */
        accessToken: string;
        /** 用于调用刷新`accessToken`的接口时所需的`token` */
        refreshToken: string;
      };
    };
    user: {
      name: string;
      id: string;
      avatarUrl: string;
      roles: Array<{ code: string }>;
    };
  };
};

export type RefreshTokenResult = {
  success: boolean;
  data: {
    /** `token` */
    accessToken: string;
    /** 用于调用刷新`accessToken`的接口时所需的`token` */
    refreshToken: string;
    /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
    expires: number;
  };
};

/** 登录 */
export const getLogin = (data?: object) => {
  return http.request<LoginResult>("post", "/operations/Casdoor/Login", {
    data
  });
};

/** 刷新token */
export const refreshTokenApi = (data?: object) => {
  return http.request<RefreshTokenResult>("post", "/refreshToken", { data });
};

/** 发送验证码 */
export const sendVerifyCode = (data?: object) => {
  return http.request("post", "/operations/Casdoor/SendCode", { data });
};

