import {
  AxiosError,
  AxiosResponse,
  CustomParamsSerializer,
  InternalAxiosRequestConfig
} from "axios";
import { stringify } from "qs";
import { getToken, formatToken } from "@/utils/auth";
import axios from "axios";
import { ElMessage } from "element-plus";

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
    return JSON.stringify(data.data || {});
  }
];

/** 请求拦截 */
instance.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 在请求发送之前做一些处理，比如添加 token、headers 等
    const token = getToken();
    if (token) {
      config.headers["Authorization"] = formatToken(token.accessToken);
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
