# 飞布介绍

## 是什么？

飞布是下一代 API 开发平台，灵活开放、多语言兼容、简单易学，对标 Firebase，但无供应商锁定。 它帮助你构建生产级 WEB API，但无需花时间重复 coding。

产品愿景：极致开发体验，`飞速布署`应用！

[前往官网->](https://www.fireboom.io/)

## 为什么？

飞布主要面向后端开发者，核心目标是快速构建 API，将开发者从重复工作中解放出来，专注更有价值的业务逻辑。

首先，业务型 Web 应用 80% 由样板代码组成，例如增删改查，权限管理，用户管理，消息或者通知。一次又一次的建立这些功能，不仅乏味，而且减少了我们集中在软件与竞争对手不同之处的时间。

- 增删改查：绝大多数偏业务型项目，都是增删改查，复杂点的包括关联查询等
- 验证鉴权：所有生产型项目都需要身份验证和身份鉴权，且实现该功能需要耗费大量人力
- 文件存储：绝大数应用都需要文件存储，用来存储用户头像等，实现文件上传和管理也较为繁琐

其次，除了重复性工作，后端开发者往往还要实现非功能需求，这些需求不仅消耗大量精力，而且有一定的技术门槛。

- N+1 缓存：避免关联查询时重复查询数据的问题，提高应用性能
- 实时推送：对于 IM 聊天等应用，需要实现实时推送功能（传统方式需要使用 websocket 等技术）

最后，当前市场上存在诸多 API 开发框架，但这些框架都基于某种特定编程语言实现，需要开发者掌握特定编程语言才能上手使用。

## 怎么做？

针对上述问题，飞布采用完全不同的思路。飞布采用声明式开发方式，它以 API 为中心，将所有数据抽象为 API，包括 REST API，GraphQL API ，数据库甚至消息队列等，通过统一协议 GraphQL 把他们聚合为“超图”，同时通过可视化界面，从“超图”中选择子集 Operation 作为函数签名，并将其编译为 REST-API。开发者通过界面配置，即可开启某 API 的的缓存或实时推送功能。此外，飞布基于 HTTP 协议实现了 HOOKS 机制，方便开发者采用任何喜欢的语言实现自定义逻辑。

# 运行

如果你是通过命令行初始化的项目，那么就直接执行下面的命令行启动

```shell
./fireboom dev
```

如果项目根目录中没有`fireboom`二进制文件，那么使用下面的命令行下载

```shell
curl -fsSL https://www.fireboom.io/update.sh | bash
```

然后再执行上面的启动脚本。

启动成功日志：

```sh
Web server started on http://localhost:9123
```

打开控制面板

[http://localhost:9123](http://localhost:9123)

## 更新

```shell
# 同时更新命令行和前端资源
curl -fsSL https://www.fireboom.io/update.sh | bash
```

## 钩子服务

下面以`ts`语言钩子为例，如果是`golang`语言则将`custom-ts`改为`custom-go`

```shell
# 依赖安装
./custom-ts/scripts/install.sh

# 开发时运行
./custom-ts/scripts/run-dev.sh

# 生产时运行
# 先打包（如果没有该文件则跳过 build）
./custom-ts/scripts/run-build.sh
./custom-ts/scripts/run-prod.sh
```

# 文档

更多文档请访问我们的[文档中心](https://ansons-organization.gitbook.io/)