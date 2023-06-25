export default {
  path: "/system",
  meta: {
    title: "系统管理"
  },
  children: [
    {
      path: "/system/user",
      name: "UserList",
      component: () => import("@/views/system/user/user.vue"),
      meta: {
        title: "用户管理"
      }
    },
    {
      path: "/system/role",
      name: "RoleList",
      component: () => import("@/views/system/role/index.vue"),
      meta: {
        title: "角色管理"
      }
    },
    {
      path: "/system/permission",
      name: "PermissionManage",
      component: () => import("@/views/system/permission/index.vue"),
      meta: {
        title: "权限管理"
      }
    },
    {
      path: "/system/menu",
      name: "MenuList",
      component: () => import("@/views/system/menu/index.vue"),
      meta: {
        title: "菜单管理"
      }
    }
  ]
};
