import axios from "@/utils/http";

export type LoginResult = {
  data: {
    accessToken: string;
    expireIn: number;
    refreshToken: string;
  };
  msg: string;
  success: boolean;
  roles: string[];
  perms: string[];
  avatar: string;
  menus: Array<{ level: number; name: string; path: string; title: string }>;
  username: string;
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
      avatar: string;
      roles: Array<{ code: string }>;
    };
  };
};

export type RefreshTokenResult = {
  data: {
    data: {
      /** `token` */
      accessToken: string;
      /** 用于调用刷新`accessToken`的接口时所需的`token` */
      refreshToken: string;
      /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
      expires: number;
    };
  };
};

/** 登录 */
export const getLogin = (data?: object) => {
  return axios.post<LoginResult>("/proxy/login", {
    data
  });
};

/** 刷新token */
export const refreshTokenApi = (data?: object) => {
  return axios.post<RefreshTokenResult>("/operations/Casdoor/RefreshToken", {
    data
  });
};

/** 发送验证码 */
export const sendVerifyCode = (data?: object) => {
  return axios.post("/operations/Casdoor/SendCode", { data });
};
