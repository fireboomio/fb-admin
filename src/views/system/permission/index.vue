<script setup lang="ts">
import api from "@/api";
import { ref, onMounted } from "vue";
import { ElTable } from 'element-plus'
import { Perm, sendPermission, PermSyncReq } from '@/api/system'
import axios from "axios";
const dataTableRef = ref<InstanceType<typeof ElTable>>();

defineOptions({
  name: "PermManage"
});

const loading = ref(false);
let dataSource = [];

const total = ref(0);
const selections = ref<string[]>([]);
async function handleQuery() {
  loading.value = true;
  const { error, data } = await api.query({
    operationName: "System/Operation/GetMany",
  });
  if (!error) {
    dataSource = data!.data!;
    total.value = dataSource.length
    initTable();
    axios.get("/operations/System/Perm/GetBindPerms").then(res => {
      const originPerms = res.data.data.data.map(item => item.path) ?? []
      selections.value = originPerms
      const selectedPerms = dataSource.filter(item => {
        return originPerms.includes(item.title)
      })
      for (const perm of selectedPerms) {
        dataTableRef.value!.toggleRowSelection(perm, true)
      }
    })
  }

  loading.value = false;
}

onMounted(() => {
  handleQuery();

});

let multipleSelection = []
const handleSelectionChange = (val) => {
  multipleSelection = val
  selections.value = val.map(item => {
    item.title
  })
}
let dataPush = []
const pushAllSelect = () => {
  multipleSelection.forEach((item) => {
    const obj: Perm = {
      id: item.id,
      createdAt: (new Date(item.createTime)).toISOString(),
      enabled: item.enabled ? 1 : 0,
      method: item.method,
      path: item.title,
    }
    dataPush.push(obj)
  })

  // 发送请求到api:  /operations/System/Perm/CreateMany
  const perm: PermSyncReq = { data: dataPush }
  sendPermission(perm).then((res) => {
    console.log(res);
  })
}

const page = ref<number>(1);
const size = ref<number>(10);

const initTable = () => {
  return dataSource.slice(
    (page.value - 1) * size.value,
    page.value * size.value
  );

}
const getRowKey = (val) => {
  return val.title
}
</script>

<template>
  <div class="app-container">
    <el-card shadow="never">
      <el-table ref="dataTableRef" v-loading="loading" :data="initTable()" @selection-change="handleSelectionChange"
        :row-key="getRowKey" highlight-current-row border>
        <el-table-column label="勾选" type="selection" :reserve-selection="true" />
        <el-table-column label="请求路径" prop="title" width="250" />
        <el-table-column label="Method" prop="method" width="150" />
        <el-table-column label="请求类型" prop="operationType" />
        <el-table-column label="是否实时" prop="liveQuery" width="100" />
        <el-table-column label="是否启用" prop="enabled" width="100" />
      </el-table>
      <!-- 实现分页功能 -->
      <el-pagination small background layout="prev, pager, next" v-model:current-page="page" :total="total" :size="10"
        class="mt-4" />
      <div style="margin-top: 20px">
        <el-button @click="pushAllSelect">勾选项同步到数据库</el-button>
      </div>

    </el-card>
  </div>
</template>
