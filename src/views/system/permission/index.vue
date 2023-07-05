<script setup lang="ts">
import api, { convertPageQuery } from "@/api";
import { API } from "../types";
import { ref, onMounted, reactive } from "vue";
import { ElTable } from 'element-plus'
import { Perm, sendPermission, PermSyncReq } from '@/api/system'
defineOptions({
  name: "PermManage"
});

const loading = ref(false);
const dataSource = ref<API[]>();

const total = ref(0);

/**
 * 查询
 */
const queryParams = reactive<PageQuery>({
  pageNum: 1,
  pageSize: 10,
  name: ""
});

async function handleQuery() {
  loading.value = true;
  const { error, data } = await api.query({
    operationName: "System/Operation/GetMany",
    input: convertPageQuery(queryParams, { containsFields: ["name"] })
  });
  if (!error) {
    dataSource.value = data!.data!;

  }
  loading.value = false;
}

onMounted(() => {
  handleQuery();
});
let multipleSelection = []
const handleSelectionChange = (val) => {
  multipleSelection = val
}
let dataPush = []
const pushAllSelect = () => {
  multipleSelection.forEach((item) => {
    const obj: Perm = {
      createdAt: (new Date(item.createTime)).toISOString(),
      enabled: item.enabled ? 1 : 0,
      method: item.method,
      path: item.title,
    }
    dataPush.push(obj)
  })
  console.log(dataPush);
  // 发送请求到url:  /operations/System/Perm/CreateMany
  const perm: PermSyncReq = { data: dataPush }
  sendPermission(perm).then((res) => {
    console.log(res);
  })
}
</script>

<template>
  <div class="app-container">
    <el-card shadow="never">
      <el-table ref="dataTableRef" v-loading="loading" :data="dataSource" @selection-change="handleSelectionChange"
        highlight-current-row border>
        <el-table-column label="勾选" type="selection" />
        <el-table-column label="请求路径" prop="title" />
        <el-table-column label="Method" prop="method" width="150" />
        <el-table-column label="请求类型" prop="operationType" />
        <el-table-column label="是否实时" prop="liveQuery" width="100" />
        <el-table-column label="是否启用" prop="enabled" width="100" />
      </el-table>
      <div style="margin-top: 20px">
        <el-button @click="clearAll">清除所有选择</el-button>
        <el-button @click="pushAllSelect">勾选项同步到数据库</el-button>

      </div>
      <el-pagination v-if="total > 0" :total="total" :page-count="queryParams.pageNum"
        :page-size="queryParams.pageSize" />
    </el-card>
  </div>
</template>
