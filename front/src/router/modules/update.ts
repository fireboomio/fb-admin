export default {
  path: "/update",
  redirect: "/welocome",
  meta: {
    title: "编辑",
    showLink: false,
    rank: 9
  },
  children: [
    {
      path: "/user/updateUser",
      name: "UserUpdateManage",
      component: () => import("@/views/system/user/updateUser/index.vue"),
      meta: {
        title: "用户编辑",
        roles: ["admin"],
      },
    },
    {
      path: "/post/updatePost",
      name: "PostUpdate",
      component: () => import("@/views/post/updatePost/index.vue"),
      meta: {
        title: "文章编辑"
      }
    }

  ]
} as RouteConfigsTable;
