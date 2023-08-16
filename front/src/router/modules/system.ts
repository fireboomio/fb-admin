export default {
  path: "/system",
  meta: {
    icon: "setting",
    title: "系统管理",
    rank: 3
  },
  children: [
    {
      path: "/system/menu",
      name: "MenuManage",
      component: () => import("@/views/system/menu/index.vue"),
      meta: {
        title: "菜单管理",
        roles: ["admin"]
      }
    },
    {
      path: "/system/user",
      name: "UserManage",
      component: () => import("@/views/system/user/index.vue"),
      meta: {
        title: "用户管理",
        roles: ["admin"]
      },
    },
    {
      path: "/system/role",
      name: "RoleManage",
      component: () => import("@/views/system/role/index.vue"),
      meta: {
        title: "角色管理",
        roles: ["admin"]
      }
    },
    {
      path: "/system/log",
      name: "LogManage",
      component: () => import("@/views/system/log/index.vue"),
      meta: {
        title: "日志管理",
        roles: ["admin"],
      },
    },
  ]
};
