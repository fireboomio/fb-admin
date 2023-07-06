<!--  线 + 柱混合图 -->
<template>
  <el-card>
    <template #header>月销售额</template>
    <div ref="chartElRef" :class="className" class="h-96" />
  </el-card>
</template>

<script setup lang="ts">
import * as echarts from "echarts";
import api from "@/api";
import { ref, onMounted } from "vue";

defineProps({
  className: {
    type: String,
    default: ""
  }
});

const chartElRef = ref<HTMLDivElement>();
const chartRef = ref<echarts.ECharts>();

onMounted(async () => {
  // 图表初始化
  chartRef.value = echarts.init(chartElRef.value!);

  // 大小自适应
  window.addEventListener("resize", () => {
    chartRef.value?.resize();
  });
  const { data, error } = await api.query({
    operationName: "Statistics/MonthlySales"
  });
  if (!error) {
    chartRef.value.setOption({
      tooltip: {
        trigger: "item"
      },
      grid: {
        left: 80,
        top: 20,
        bottom: 24,
        right: 20
      },
      xAxis: {
        type: "category",
        data: data!.data!.map(item => item.months) ?? []
      },
      yAxis: {
        type: "value"
      },
      series: [
        {
          data: data!.data!.map(item => item.totalSales) ?? [],
          type: "bar"
        }
      ]
    });
  }
});
</script>
