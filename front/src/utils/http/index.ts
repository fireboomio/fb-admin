import {
  AxiosError,
  AxiosResponse,
  CustomParamsSerializer,
  InternalAxiosRequestConfig
} from "axios";
import { stringify } from "qs";
import { getToken, formatToken, setToken } from "@/utils/auth";
import axios from "axios";
import { ElMessage } from "element-plus";
import { useUserStoreHook } from "@/store/modules/user";
import router from "@/router";
import { storageSession } from "@pureadmin/utils";
const whiteList = ["/login", "/operations/Casdoor/RefreshToken"];
// 相关配置请参考：www.axios-js.com/zh-cn/docs/#axios-request-config-1
const instance = axios.create({
  // 请求超时时间
  timeout: 10000,
  headers: {
    Accept: "application/json, text/plain, */*",
    "Content-Type": "application/json",
    "X-Requested-With": "XMLHttpRequest"
  },
  // 数组格式参数序列化（https://github.com/axios/axios/issues/5142）
  paramsSerializer: {
    serialize: stringify as unknown as CustomParamsSerializer
  }
});

// 自定义默认的请求数据处理方式
instance.defaults.transformRequest = [
  (data: any) => {
    // 对请求数据进行处理，去掉外层的data
    if (!data) {
      return;
    }
    return JSON.stringify(data.data || {});
  }
];

/** 请求拦截 */
instance.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 在请求发送之前做一些处理，比如添加 token、headers 等
    if (whiteList.includes(config.url)) {
      return config;
    }
    const data = getToken();
    if (data) {
      const now = new Date().getTime();
      const expired = ((parseInt(data.expires)) - (now / 1000) <= 0); // 是否过期
      if (expired) {
        // token过期刷新
        // 获取refreshToken的值
        const refreshToken = storageSession().getItem("user-info")['refreshToken'];
        useUserStoreHook().handRefreshToken({ refreshToken })
          .then(res => {
            const token = res.data.data.data.data.accessToken;
            config.headers["Authorization"] = formatToken(token);
          }).catch(err => {
            // refreshToken已失效，跳转登录界面
            useUserStoreHook().logOut();
          })
      } else {
        config.headers["Authorization"] = formatToken(data.accessToken);
      }
    }
    return config;
  },
  (error: AxiosError) => {
    // 对请求错误做处理
    return Promise.reject(error);
  }
);

instance.interceptors.response.use(
  (response: AxiosResponse) => {
    // 对响应数据做处理
    return response;
  },
  (error: AxiosError) => {
    // 对响应错误做处理
    if (error.response.status === 401) {
      ElMessage.error("权限未认证");
    } else {
      ElMessage.error(error.message);
    }
    return Promise.reject(error);
  }
);

export default instance;
