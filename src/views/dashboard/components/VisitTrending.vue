<!--  线 + 柱混合图 -->
<template>
  <el-card>
    <template #header>日访问量趋势</template>
    <div ref="chartElRef" :class="className" class="h-60" />
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
    operationName: "Statistics/VisitTrending"
  });
  if (!error) {
    chartRef.value.setOption({
      tooltip: {
        trigger: "axis"
      },
      grid: {
        left: 40,
        top: 20,
        bottom: 20,
        right: 20
      },
      xAxis: {
        type: "category",
        data: data!.data!.map(item => item.days) ?? []
      },
      yAxis: {
        type: "value"
      },
      series: [
        {
          data: data!.data!.map(item => item.count) ?? [],
          type: "line",
          smooth: true,
          areaStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: "rgba(58,77,233,0.8)"
              },
              {
                offset: 1,
                color: "rgba(58,77,233,0.3)"
              }
            ])
          }
        }
      ]
    });
  }
});
</script>
