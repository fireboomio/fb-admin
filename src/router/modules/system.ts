export default {
  path: "/system",
  meta: {
    title: "系统管理",
    rank: 3
  },
  children: [
    {
      path: "/system/menu/index",
      name: "MenuManage",
      component: () => import("@/views/system/menu/index.vue"),
      meta: {
        title: "菜单管理",
        roles: ["admin"]
      }
    },
    {
      path: "/system/user/index",
      name: "UserManage",
      component: () => import("@/views/system/user/index.vue"),
      meta: {
        title: "用户管理",
        roles: ["admin"]
      }
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
    // {
    //   path: "/system/perm",
    //   name: "PermManage",
    //   component: () => import("@/views/system/permission/index.vue"),
    //   meta: {
    //     title: "权限管理",
    //     roles: ["admin"],
    //   }
    // }
  ]
};
