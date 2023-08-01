<script lang="ts">
export default { name: "Dashboard" };
</script>

<script setup lang="ts">
import MonthlySales from "./components/MonthlySales.vue";
import SalesTop10 from "./components/SalesTop10.vue";
import VisitTrending from "./components/VisitTrending.vue";
import SaleTypePercent from "./components/SaleTypePercent.vue";
</script>

<template>
  <div class="dashboard-container">
    <!-- Echarts 图表 -->
    <el-row :gutter="16">
      <el-col :span="16" class="mb-4">
        <MonthlySales />
      </el-col>

      <el-col :span="8" class="mb-4">
        <SalesTop10 />
      </el-col>

      <el-col :span="12" class="mb-4">
        <VisitTrending />
      </el-col>
      <el-col :span="12" class="mb-4">
        <SaleTypePercent />
      </el-col>
    </el-row>
  </div>
</template>

<style lang="scss" scoped>
.dashboard-container {
  padding: 24px;
  position: relative;
  .user-avatar {
    height: 40px;
    width: 40px;
    border-radius: 50%;
  }

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
    z-index: 99;
  }

  .data-box {
    font-weight: bold;
    padding: 20px;
    color: var(--el-text-color-regular);
    background: var(--el-bg-color-overlay);
    box-shadow: var(--el-box-shadow-dark);
    border-color: var(--el-border-color);
    display: flex;
    justify-content: space-between;
  }
  .svg-icon {
    fill: currentColor !important;
  }
}
</style>
