<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="名称" prop="name">
          <el-input v-model="queryParams.name" placeholder="用户名" clearable style="width: 200px"
            @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item>
          <Auth value="system:user:query">
            <el-button type="primary" @click="handleQuery">
              <Icon icon="ep:search" />搜索
            </el-button>
          </Auth>
          <Auth value="system:user:query">
            <el-button @click="resetQuery">
              <Icon icon="ep:refresh" />
              重置
            </el-button>
          </Auth>

        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <el-table v-loading="loading" :data="userList" highlight-current-row border>
        <el-table-column key="id" label="编号" align="center" prop="id" width="100" />
        <el-table-column key="name" label="用户名" align="center" prop="name" width="200" />
        <el-table-column label="头像" width="80" align="center" prop="avatar">
          <template #default="scope">
            <!-- <el-image :src="scope.row.avatar" :preview-src-list="[scope.row.avatar]" /> -->
            <Avatar :username="scope.row.name" :src="scope.row.avatar" background-color="#ccc" color="#fff"
              style=" vertical-align: middle;" :inline="true">
            </Avatar>
          </template>

        </el-table-column>
        <el-table-column label="创建时间" align="center" prop="createdAt" width="180" :formatter="(row, col, v) => (v ? new Date(v).toLocaleDateString() : '')
          " />
        <el-table-column label="操作" fixed="right" min-width="220">
          <template #default="scope">

            <el-button type="primary" size="small" link @click="bindUser(scope.row)">绑定角色</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination :total="total" :page-count="queryParams.pageNum" :page-size="queryParams.pageSize"
        @pagination="handleQuery" />
    </el-card>
    <role-bind v-model="bindVisible" :user="bindingUser" />
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "UserManage"
});

import { User } from "../types";
import api, { convertPageQuery } from "@/api";
import { ref, reactive, onMounted } from "vue";
import { ElForm, ElPagination } from "element-plus";
import RoleBind from /* @vite-ignore */ "./role.bind.vue";
import { Icon } from '@iconify/vue';

const queryFormRef = ref(ElForm); // 查询表单
const loading = ref(false);
const total = ref(0);
const bindVisible = ref(false);
const bindingUser = ref<User>();

const queryParams = reactive<PageQuery>({
  pageNum: 1,
  pageSize: 10,
  name: ""
});
const userList = ref<User[]>([]);

/**
 * 查询
 */
async function handleQuery() {
  loading.value = true;
  const { error, data } = await api.query({
    operationName: "System/User/GetList",
    input: convertPageQuery(queryParams, { containsFields: ["name"] })
  });
  if (!error) {
    userList.value = data!.data!;
    total.value = data!.data!.length;
  }
  loading.value = false;
}

/**
 * 重置查询
 */
function resetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  handleQuery();
}

function bindUser(user: User) {
  bindVisible.value = true;
  bindingUser.value = user;
}

onMounted(() => {
  handleQuery(); // 初始化用户列表数据
});
</script>
