<script setup lang="ts">
import api from "@/api";
import { API } from "../types";
import { ref, onMounted } from "vue";

defineOptions({
  name: "PermissionManage"
});

const loading = ref(false);
const dataSource = ref<API[]>();

/**
 * 查询
 */
async function handleQuery() {
  loading.value = true;
  const { error, data } = await api.query({
    operationName: "System/Operation/GetMany"
  });
  if (!error) {
    dataSource.value = data!.data!;
  }
  loading.value = false;
}

onMounted(() => {
  handleQuery();
});
</script>

<template>
  <div class="app-container">
    <el-card shadow="never">
      <el-table
        ref="dataTableRef"
        v-loading="loading"
        :data="dataSource"
        highlight-current-row
        border
      >
        <el-table-column label="请求路径" prop="title" />
        <el-table-column label="Method" prop="method" width="150" />
        <el-table-column label="请求类型" prop="operationType" />
        <el-table-column label="是否实时" prop="liveQuery" width="100" />
        <el-table-column label="是否启用" prop="enabled" width="100" />
      </el-table>
    </el-card>
  </div>
</template>
