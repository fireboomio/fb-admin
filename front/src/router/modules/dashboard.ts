export default {
  path: "/dashboard",
  meta: {
    title: "仪表盘",
    icon: "pieChart"
  },
  children: [
    {
      path: "/dashboard/index",
      name: "Dashboard",
      component: () => import("@/views/dashboard/index.vue"),
      meta: {
        title: "仪表盘"
      }
    }
  ]
};
