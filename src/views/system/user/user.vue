<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="名称" prop="name">
          <el-input
            v-model="queryParams.name"
            placeholder="用户名"
            clearable
            style="width: 200px"
            @keyup.enter="handleQuery"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery"
            ><i-ep-search />搜索</el-button
          >
          <el-button @click="resetQuery">
            <i-ep-refresh />
            重置</el-button
          >
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <el-table v-loading="loading" :data="userList">
        <el-table-column
          key="id"
          label="编号"
          align="center"
          prop="id"
          width="300"
        />
        <el-table-column key="name" label="用户名" align="center" prop="name" />
        <el-table-column label="头像" width="60" align="center" prop="avatar">
          <template #default="scope">
            <el-image
              :src="scope.row.avatar"
              :preview-src-list="[scope.row.avatar]"
            />
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          align="center"
          prop="createdAt"
          width="180"
          :formatter="
            (row, col, v) => (v ? new Date(v).toLocaleDateString() : '')
          "
        />
        <el-table-column label="操作" fixed="right" width="220">
          <template #default="scope">
            <el-button
              type="primary"
              size="small"
              link
              @click="bindUser(scope.row)"
              >绑定角色</el-button
            >
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-if="total > 0"
        :total="total"
        :page-count="queryParams.pageNum"
        :page-size="queryParams.pageSize"
      />
    </el-card>
    <role-bind v-model="bindVisible" :user="bindingUser" />
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "UserList"
});

import { User } from "../types";
import api, { convertPageQuery } from "@/api";
import { ref, reactive, onMounted } from "vue";
import { ElForm } from "element-plus";
import RoleBind from /* @vite-ignore */ "./role.bind.vue";

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
    total.value = data!.total!;
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
