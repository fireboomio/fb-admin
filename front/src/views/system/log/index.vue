<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="IP: " prop="ip">
          <el-input v-model="logInfoSearch.ip" placeholder="请输入Ip" clearable style="width: 200px" @keyup.enter="search" />
        </el-form-item>
        <el-form-item label="Method: " prop="method">
          <el-input v-model="logInfoSearch.method" placeholder="请输入Method" clearable style="width: 200px"
            @keyup.enter="search" />
        </el-form-item>
        <el-form-item label="path: " prop="path">
          <el-input v-model="logInfoSearch.path" placeholder="请输入path" clearable style="width: 200px"
            @keyup.enter="search" />
        </el-form-item>
        <el-form-item label="id: " prop="id">
          <el-input v-model="logInfoSearch.id" placeholder="请输入id" clearable style="width: 200px" @keyup.enter="search" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="search">
            <Icon icon="ep:search" />搜索
          </el-button>
          <el-button @click="resetQuery">
            <Icon icon="ep:refresh" />
            重置
          </el-button>
        </el-form-item>
        <el-form-item label="开启日志收集">
          <el-switch v-model="isOpen" inline-prompt />
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <template #header>
        <el-button type="danger" :disabled="ids.length === 0" @click="handleDelete()">
          <Icon icon="ep:delete" />删除
        </el-button>
      </template>
      <el-table v-loading="loading" :data="logList" highlight-current-row border @selection-change="handleSelectionChange"
        @sort-change="changeTableSort">
        <el-table-column type="selection" width="55" />
        <el-table-column label="ID" align="center" prop="id" width="100" :sortable="'custom'" />
        <el-table-column label="用户名" align="center" prop="userName" width="100" />
        <el-table-column label="方法" align="center" prop="method" width="100" />
        <el-table-column label="ip地址" align="center" prop="ip" width="200" />
        <el-table-column label="路径" align="center" prop="path" width="300" />
        <el-table-column label="操作时间" align="center" prop="updatedAt" width="200" :sortable="'custom'" />
        <el-table-column label="用户信息" align="center" prop="ua" width="400" />
        <el-table-column label="操作" align="center" min-width="200">
          <template #default="scope">
            <el-button type="primary" size="small" link @click="deleteLogById(scope.row.id)">
              <Icon icon="ep:delete" />删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination :total="total" v-model:current-page="queryParams.pageNum" :page-size="queryParams.pageSize"
        @current-change="handleQuery" />
    </el-card>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "UserManage"
});

import { User } from "../types";
import api, { convertPageQuery } from "@/api";
import { ref, reactive, onMounted, onBeforeMount, watch } from "vue";
import { ElForm, ElMessage, ElPagination, FormRules, ElMessageBox } from "element-plus";
import { Icon } from '@iconify/vue';
import { LogType } from "@/store/modules/types";
import { getAllLogNumber, GetLog, GetIsOpen, ChangeOpen, deleteLog, getLikeLog, deleteOneLog } from "@/api/system";

const queryFormRef = ref(ElForm); // 查询表单
const currentProp = ref();
const currentOrder = ref();
const operateTime = ref();
const loading = ref(false);
const total = ref(0);
const bindVisible = ref(false);
const bindingUser = ref<User>();
const isOpen = ref<boolean>();
const ids = ref<number[]>([]);
const dialog = reactive<DialogOption>({
  visible: false
});
const queryParams = reactive<PageQuery>({
  pageNum: 1,
  pageSize: 10,
  name: "1"
});
const logInfoSearch = reactive({
  ip: "",
  method: "",
  path: "",
  id: "",
})
const logList = ref<LogType[]>([]);
const logListBySearch = ref<LogType[]>([]);
/**
 * 获取日志数量
 */
async function getAllNumber() {
  await getAllLogNumber().then(res => {
    total.value = res.data.data.logNum._count.id;
  })
}
function formatTimestampToISO8601(timestamp) {
  const date = new Date(timestamp);
  return date.toISOString();
}
/**
 *  监听排序
 */
function changeTableSort(column) {
  const fieldName = column.prop;
  const sortingType = column.order;
  currentOrder.value = sortingType;
  currentProp.value = fieldName;
  // 如果是时间，则需转成时间戳进行比较
  if (fieldName === "updatedAt") {
    logList.value.map(item => {
      item.updatedAt = Date.parse(item.updatedAt);
    })
  }
  if (sortingType === "descending") {
    logList.value = logList.value.sort((a, b) => b[fieldName] - a[fieldName]);
  } else if (sortingType === "ascending") {
    logList.value = logList.value.sort((a, b) => a[fieldName] - b[fieldName]);
  }
  if (fieldName === "updatedAt") {
    logList.value.map(item => {
      item.updatedAt = formatTimestampToISO8601(item.updatedAt);
    })
  }
}
/**
 * 查询
 */
async function handleQuery() {

  loading.value = true;
  if (logInfoSearch.id || logInfoSearch.ip || logInfoSearch.method || logInfoSearch.path) {    // 模糊查询的数据
    total.value = logListBySearch.value.length;
    logList.value = logListBySearch.value.slice((queryParams.pageNum - 1) * queryParams.pageSize, queryParams.pageNum * queryParams.pageSize);
  } else {
    const skip = queryParams.pageSize * (queryParams.pageNum - 1);
    await GetLog(skip, queryParams.pageSize).then(res => {
      logList.value = res.data.data.data;
      logListBySearch.value.length = 0;
    }).catch(err => {
      ElMessage.error(`获取日志信息失败,错误信息:${err}`)
    })
    getAllNumber();
  }
  loading.value = false;
}

/**
 * 行checkbox change事件
 */
function handleSelectionChange(selection: any) {
  ids.value = selection.map((item: any) => item.id);
  console.log(ids.value);

}

async function search() {
  loading.value = true;
  const searchData = {};
  if (logInfoSearch.id) {
    searchData['id'] = logInfoSearch.id;
  }
  if (logInfoSearch.ip) {
    searchData['ip'] = logInfoSearch.ip;
  }
  if (logInfoSearch.method) {
    searchData['method'] = logInfoSearch.method;
  }
  if (logInfoSearch.path) {
    searchData['path'] = logInfoSearch.path;
  }
  if (JSON.stringify(searchData) == "{}") {
    queryParams.pageNum = 1;
    handleQuery();
  } else {
    await getLikeLog(searchData).then(res => {
      ElMessage.success("查询成功");
      if (res.data.data.main_findUniqueadmin_apilog) {   // 若精准查询成功，则将精准查询数据放在第一条
        const accurate = res.data.data.main_findUniqueadmin_apilog;
        logListBySearch.value = res.data.data.main_findManyadmin_apilog.filter(item => {
          return item.id !== accurate.id;
        });
        logListBySearch.value.unshift(accurate);
      } else {
        logListBySearch.value = res.data.data.main_findManyadmin_apilog;
      }
      handleQuery();
    }).catch(err => {
      ElMessage.error("查询失败");
    })
  }
  loading.value = false;
}
/**
 * 获取当前是否开启收集日志信息
 */
async function getIsOpen() {
  await GetIsOpen().then(res => {
    isOpen.value = res.data.data.main_findFirstdic.isOpen;
  })
}
/**
 * 重置查询
 */
async function resetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  logInfoSearch.ip = "";
  logInfoSearch.method = "";
  logInfoSearch.path = "";
  logInfoSearch.id = "";
  await handleQuery();
  if (currentOrder.value && currentProp.value) {
    changeTableSort({ "order": currentOrder.value, "prop": currentProp.value });
  }
}
/**
 * 删除多个log
 */
function handleDelete() {
  const idList = ids.value;
  if (!idList.length) {
    ElMessage.warning("请勾选删除项");
    return;
  }
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    console.log(idList);
    console.log({
      equals: idList,
    });
    deleteLog(idList).then(res => {
      ElMessage.success("删除成功");
      resetQuery();
    })
  });
}

/**
 * 根据日志id数组删除日志
 */
function deleteLogById(id) {
  ElMessageBox.confirm(
    '确定要删除该日志嘛?',
    'Warning',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    deleteOneLog(id).then(res => {
      ElMessage.success("删除成功");
      resetQuery()
    })
  }).catch(() => {
    ElMessage.info("已取消");
  });
}

onMounted(() => {
  getAllNumber();
  handleQuery();                    // 初始化用户列表数据
  getIsOpen();
});

/**
 * 监视isOpen
 */
watch(isOpen, (newValue) => {
  const set = {
    "set": isOpen.value,
  }
  ChangeOpen(set).then(res => {
    if (newValue) {
      ElMessage.success("已开启收集日志");
    } else {
      ElMessage.success("已关闭收集日志");
    }
  })
})
</script>
