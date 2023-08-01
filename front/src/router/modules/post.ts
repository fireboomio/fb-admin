export default {
  path: "/post",
  meta: {
    title: "文章管理",
    rank: 2
  },
  children: [
    {
      path: "/post/index",
      name: "PostManage",
      component: () => import("@/views/post/index.vue"),
      meta: {
        title: "文章管理"
      }
    }
  ]
};
