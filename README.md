# 介绍
fireBoom-admin是一款基于飞布和[Pure-admin](https://yiming_chang.gitee.io/pure-admin-doc/)的中后台管理系统，使用Vue3、TypeScript、Vite、Element-Plus、Pinia、Echarts等主流技术栈进行开发。
## 飞布后端
https://github.com/Echoidf/fireboon-admin-server

# 一、快速开始
## 拉取代码
**从GitHub上拉取**
`git clone [https://github.com/Echoidf/fireboom-admin-web.git](https://github.com/Echoidf/fireboom-admin-web.git)`
## 本地开发

1. 安装依赖

`pnpm install`
`npm install -g pnpm`

2. 启动平台

`pnpm dev`

3. 打包项目

`pnpm build`

```
├── .github  # GitHub 配置文件
├── .husky   # 代码提交前校验配置文件
├── .vscode  # IDE 工具推荐配置文件
│   │   ├── extensions.json       # 一键安装平台推荐的 vscode 插件
│   │   ├── settings.json         # 设置扩展程序或 vscode 编辑器的一些属性
│   │   ├── vue3.0.code-snippets  # vue3.0 代码片段
│   │   └── vue3.2.code-snippets  # vue3.2 代码片段
│   │   └── vue3.3.code-snippets  # vue3.3 代码片段
├── build    # 构建工具
│   │   ├── cdn.ts       # 打包时采用 cdn 模式
│   │   ├── compress.ts  # 打包时启用 gzip 压缩或 brotli 压缩
│   │   ├── index.ts     # 导出环境变量、跨域代理函数
│   │   ├── info.ts      # 输出打包信息（大小、用时）
│   │   ├── optimize.ts  # vite 依赖预构建配置项
│   │   ├── plugins.ts   # vite 相关插件存放处
├── locales   # 国际化文件存放处
│   │   ├── en.yaml      # 英文配置
│   │   ├── zh-CN.yaml   # 中文配置
├── node_modules  # 模块依赖
├── public  # 静态资源
│   │   ├── favicon.ico        # favicon
│   │   ├── serverConfig.json  # 全局配置文件（打包后修改也可生效）
├── src
│   ├── api         # 接口请求统一管理
│   ├── assets      # 字体、图片等静态资源
│   ├── components  # 自定义通用组件
│   ├── config      # 获取平台动态全局配置
│   ├── directives  # 自定义指令
│   ├── layout      # 主要页面布局
│   ├── plugins     # 处理一些库或插件，导出更方便的 api
│   ├── router      # 路由配置
│   ├── store       # pinia 状态管理
│   ├── style       # 全局样式
│   │   ├── dark.scss          # 暗黑模式样式适配文件
│   │   ├── element-plus.scss  # 全局覆盖 element-plus 样式文件
│   │   ├── reset.scss         # 全局重置样式文件
│   │   ├── sidebar.scss       # layout 布局样式文件
│   │   ├── tailwind.css       # tailwindcss 自定义样式配置文件
│   │   ├── ...
│   ├── utils  # 全局工具方法
│   │   ├── http      # 封装 axios 文件
│   │   ├── progress  # 封装 nprogress
│   │   ├── auth.ts   # 处理用户信息和 token 相关
│   │   ├── ...
│   ├── views              # 存放编写业务代码页面
│   ├── App.vue            # 入口页面
│   ├── main.ts            # 入口文件
│   └── mockProdServer.ts  # mock 服务相关
├── types  # 全局 TS 类型配置
├── .dockerignore        # 排除不需要上传到 docker 服务端的文件或目录。
├── .editorconfig        # 编辑器读取文件格式及样式
├── .env                 # 全局环境变量配置
├── .env.development     # 开发环境变量配置
├── .env.production      # 生产环境变量配置
├── .env.staging         # 预发布环境变量配置
├── .eslintignore        # eslint 语法检查忽略文件
├── .eslintrc.js         # eslint 语法检查配置
├── .gitignore           # git 提交忽略文件
├── .markdownlint.json   # markdown 格式检查配置
├── .npmrc               # npm 配置文件
├── .prettierrc.js       # prettier 插件配置
├── .stylelintignore     # stylelint 插件检查忽略文件
├── commitlint.config.js # git 提交前检查配置
├── Dockerfile           # 构建docker镜像
├── index.html           # html 主入口
├── LICENSE              # 证书
├── package.json         # 依赖包管理以及命令配置
├── pnpm-lock.yaml       # 依赖包版本锁定文件
├── postcss.config.js    # postcss 插件配置
├── README.md            # README
├── stylelint.config.js  # stylelint 配置
├── tailwind.config.js   # tailwindcss 配置
├── tsconfig.json        # typescript 配置
└── vite.config.ts       # vite 配置
```
# 管理系统登录方式
两种登录方式可供选择：

1）密码登录

2）手机验证码登录

**登录流程：**

初次登录的时候，前端调用后端的登录接口，发送验证信息，后端收到请求并进行验证，验证成功，就给前端返回一个token和用户信息，前端拿到这些信息，将其存储到Pinia中，从Pinia中取出token值放入到浏览器Cookies中，取出用户信息存储到SessionStorage中。

# 状态持久化——Pinia
Pinia是Vue的存储库，允许跨组件/跨页面共享状态数据，与 Vuex 相比，Pinia 提供了一个更简单的 API，具有更少的规范，提供了 Composition-API 风格的 API，最重要的是，在与 TypeScript 一起使用时具有可靠的类型推断支持。因此，本项目使用Pinia来替代Vuex。
详情请参考文档[Pinia中文文档](https://pinia.web3doc.top/introduction.html)
# 二、权限设置
## 菜单权限
通过返回路由的`roles`配置项，查询用户的动态权限，并将其存储在Pinia中
![image.png](https://cdn.nlark.com/yuque/0/2023/png/29078601/1688980708164-1290fca5-2a2c-47a8-9a44-f1f7f5bb9cc7.png#averageHue=%23232120&clientId=u80b370e8-7558-4&from=paste&height=248&id=u057f572d&originHeight=310&originWidth=644&originalType=binary&ratio=1.25&rotation=0&showTitle=false&size=35176&status=done&style=none&taskId=u538ff28e-6b27-4c21-94a1-1cff12c99d0&title=&width=515.2)
## 按钮权限
通过hasAuth函数，进行权限组件的封装
```typescript
/** 是否有按钮级别的权限 */
function hasAuth(value: string | Array<string>): boolean {
  if (!value) return false;
  const auths = getAuths();
  if (!auths) return false;
  const isAuths = isString(value)
    ? auths.includes(value)
    : isIncludeAllChildren(value, auths);
  return isAuths ? true : false;
}

```
配置自定义指令`auth`以及全局组件`Auth`
```typescript
export const auth: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const { value } = binding;
    if (value) {
      !hasAuth(value) && el.parentNode?.removeChild(el);
    } else {
      throw new Error(
        "[Directive: auth]: need auths! Like v-auth=\"['btn.add','btn.edit']\""
      );
    }
  }
};
```
```typescript
// 全局注册按钮级别权限组件
import { Auth } from "@/components/ReAuth";
app.component("Auth", Auth);
```
给按钮加上权限
例如：

1. 使用Auth（较为推荐）
```vue
<Auth value="/Post/GetList">
	<el-button type="primary" @click="handleQuery()">
  	<Icon icon="ep:search" />搜索
  </el-button>
</Auth>
```

2. 使用v-auth（一般推荐）
```vue
<el-button v-auth="'/Post/GetList'" type="primary" @click="handleQuery()">
  	<Icon icon="ep:search" />搜索
</el-button>
```

3. 使用v-if，**需要引入**hasAuth方法（不推荐）

**引入hasAuth**
```vue
import { hasAuth } from "@/router/utils";
```
**使用v-if**
```vue
<el-button v-if="hasAuth('/Post/GetList')" type="primary" @click="handleQuery()">
  	<Icon icon="ep:search" />搜索
</el-button>
```
# 三、http请求
对axios进行二次封装，包括设置相应的请求拦截器与响应拦截器。
## 拦截器
### 请求拦截器
请求拦截器的作用是在请求发送之前进行一些操作，例如在每个请求体里加上token，统一做了处理。
```typescript
private httpInterceptorsRequest(): void {
  PureHttp.axiosInstance.interceptors.request.use(
    async (config: PureHttpRequestConfig): Promise<any> => {
      // 开启进度条动画
      // 优先判断post/get等方法是否传入回调，否则执行初始化设置等回调
      // token过期刷新
      // ......
    }
  );
}
```
### 响应拦截器
响应拦截器的作用是在接收到响应后进行一些操作，例如在服务器返回登录状态失效，需要重新登录的时候，跳转到登录页。
```typescript
private httpInterceptorsResponse(): void {
  const instance = PureHttp.axiosInstance;
  instance.interceptors.response.use(
    (response: PureHttpResponse) => {
      // 关闭进度条动画
      // 判断post/get等方法是否传入回调
      // ......
    }
  );
}
```
## 请求api

1. 发送`get`请求
```typescript
import { http } from "@/utils/http";

// params传参
export const textRequest = (params?: object) => {
  return http.request("get", "/xxx", { params });
};

// url拼接传参
export const textRequest = (params?: object) => {
  return http.request("get", "/xxx?message=" + params);
};

```

2. 发送`post`请求
```typescript
import { http } from "@/utils/http";

// params传参：会直接转换成url?这种形式
export const textRequest = (params?: object) => {
  return http.request("post", "/xxx", { params });
};

// data传参：放在请求体中
export const textRequest = (data?: object) => {
  return http.request("post", "/xxx", { data });
};

```

3. 发送`delete`请求
```typescript
import { http } from "@/utils/http";

// params传参
export const textRequest = (params?: object) => {
  return http.request("delete", "/xxx", { params });
};

// data传参
export const textRequest = (data?: object) => {
  return http.request("delete", "/xxx", { data });
};
```

4. 发送`put`请求
```typescript
import { http } from "@/utils/http";

// params传参
export const textRequest = (params?: object) => {
  return http.request("put", "/xxx", { params });
};

// data传参
export const textRequest = (data?: object) => {
  return http.request("put", "/xxx", { data });
};
```

