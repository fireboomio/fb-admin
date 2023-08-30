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

export type UserInfo = {
  countryCode: string;
  name: string;
  password: string;
  passwordType: string;
  phone: string;
}
export type RefreshTokenResult = {
  data: {
    accessToken: string;
    expireIn: number;
    refreshToken: string;
  };
  msg: string;
  success: boolean;
};

/** 登录 */
export const getLogin = (data?: object) => {
  return axios.post<LoginResult>("operations/proxy/login", {
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


/** 根据id删除用户 */
export const deleteUserById = (data: object) => {
  return axios.post("/operations/System/User/DeleteOne", { data });
};

/** 新增用户 */
export const createUser = (data: UserInfo) => {
  return axios.post("/operations/System/User/CreateOne", { data });
}

/** 更新用户信息 */
export const updateUser = (data: object) => {
  return axios.post("/operations/System/User/UpdateOne", { data });
}

/**
 *  获取用户总数
 */
export const getAllUserNumber = () => {
  return axios.get<any>("/operations/System/User/GetAllList");
}
/**
 * 用户的模糊查询
 */
export const getUserLike = (data: object) => {
  return axios.get<any>("/operations/System/User/GetLikeUser", {
    params: data,
  });
}

/**
 * 根据userId来查询用户信息
 */
export const getUserByUserId = (userId) => {
  return axios.get<any>(`/operations/System/User/GetUserByUserId?userId=${userId}`)
}
/**
 * 根据用户名查询角色信息
 */
export const getUserInfoByName = (equals: string) => {
  return axios.get<any>("/operations/System/User/GetUserInfo", {
    params: {
      equals
    }
  })
}