<!--  线 + 柱混合图 -->
<template>
  <el-card>
    <template #header>门店销售额排名</template>
    <div v-for="(item, index) in dataSource" :key="index" className="flex py-1.5 items-center">
      <div class="order" :class="{ highlight: index < 3 }">{{ index + 1 }}</div>
      <div class="ml-4">{{ item.shopName }}</div>
      <div class="ml-auto">{{ Math.round(item.totalSales! * 100) / 100 }}</div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import api from "@/api";

type SalesTop = {
  shopName?: string;
  totalSales?: number;
};
const dataSource = ref<SalesTop[]>([]);

onMounted(async () => {
  const { data, error } = await api.query({
    operationName: "Statistics/SalesTop10"
  });
  if (!error) {
    dataSource.value = data!.data!;
  }
});
</script>

<style scoped>
.order {
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  width: 18px;
  height: 18px;
  background-color: #666;
  color: #fff;
  font-size: 12px;
}

.highlight {
  background-color: #314659;
  color: #fafafa;
}
</style>
