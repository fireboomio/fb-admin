export default {
  path: "/post",
  meta: {
    title: "发布文章"
  },
  children: [
    {
      path: "/post/index",
      name: "Post",
      component: () => import("@/views/post/index.vue"),
      meta: {
        title: "发布文章",
        roles: ["admin"]
      }
    }
  ]
};
