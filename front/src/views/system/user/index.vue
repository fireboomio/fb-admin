<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="用户名" prop="name">
          <el-input v-model="searchUserData.name" placeholder="请输入用户名:" clearable style="width: 200px"
            @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item label="电话" prop="name">
          <el-input v-model="searchUserData.phone" placeholder="请输入电话:" clearable style="width: 200px"
            @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item label="用户名" prop="name">
          <el-input v-model="searchUserData.id" placeholder="请输入用户ID:" clearable style="width: 200px"
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
      <template #header>
        <el-button type="success" @click="openDialog()">
          <Icon icon="ep:plus" />新增
        </el-button>
      </template>
      <el-table v-loading="loading" :data="userList" highlight-current-row border>
        <el-table-column key="id" label="编号" align="center" prop="id" width="100" />
        <el-table-column key="name" label="用户名" align="center" prop="name" width="200" />
        <el-table-column label="头像" width="80" align="center" prop="avatar">
          <template #default="scope">
            <Avatar :username="scope.row.name" :src="scope.row.avatar" background-color="#ccc" color="#fff"
              style=" vertical-align: middle;" :inline="true">
            </Avatar>
          </template>
        </el-table-column>
        <el-table-column label="角色" width="260" align="center" prop="roles">
          <template #default="scope">
            <!-- <el-button type="success" round v-if="scope.row.roles.indexOf('admin')">admin</el-button> -->
            <el-button type="primary" round v-for="item in scope.row.roles">{{ item }}</el-button>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" align="center" prop="createdAt" width="200" :formatter="(row, col, v) => (v ? new Date(v).toLocaleDateString() : '')
          " />
        <el-table-column label="操作" fixed="right" min-width="200" align="center">
          <template #default="scope">
            <el-button type="primary" size="small" link @click="bindUser(scope.row)">绑定角色</el-button>
            <el-button type="primary" size="small" link @click="handleDelete(scope.row.id)">
              <Icon icon="ep:delete" />删除
            </el-button>
            <el-button size="small" link type="primary" @click.stop="editUser(scope.row.user_id)">
              <Icon icon="ep:edit" />修改
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination :total="total" v-model:current-page="queryParams.pageNum" v-model:page-size="queryParams.pageSize"
        @current-change="handleQuery" />
    </el-card>
    <role-bind v-model="bindVisible" :user="bindingUser" />
    <!-- 新增用户与修改用户信息 -->
    <el-dialog :title="dialog.title" v-model="dialog.visible" width="500px" @close="closeDialog">
      <el-form ref="roleFormRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="国 籍" prop="countruCode">
          <el-select placeholder="请选择国籍" filterable style="width: 77.9%" v-model="formData.countryCode">
            <el-option-group v-for="group in country" :key="group.label" :label="group.label">
              <el-option v-for="item in group.options" :key="item.value" :label="item.label" :value="item.label">
                <span style="float: left">{{ item.label }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
              </el-option>
            </el-option-group>
          </el-select>
        </el-form-item>
        <el-form-item label="用户名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="formData.password" placeholder="请输入密码" show-password />
        </el-form-item>
        <el-form-item label="确认密码" prop="passwordType">
          <el-input type="password" v-model="formData.passwordType" placeholder="确认密码" show-password />
        </el-form-item>
        <el-form-item label="电话" prop="phone">
          <el-input v-model="formData.phone" placeholder="请输入电话号码" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="addUser" :disabled="isDisabled">确 定</el-button>
          <el-button @click="closeDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "UserManage"
});

import { User } from "../types";
import api, { convertPageQuery } from "@/api";
import { ref, reactive, onMounted } from "vue";
import { ElForm, ElMessage, ElPagination, FormRules, ElMessageBox } from "element-plus";
import RoleBind from /* @vite-ignore */ "./role.bind.vue";
import { Icon } from '@iconify/vue';
import { createUser, deleteUserById, getAllUserNumber, getRoleUsers, getUserLike, UserInfo } from "@/api/user";
import { country } from "./utils/country"
import { useRouter } from 'vue-router';

const queryFormRef = ref(ElForm); // 查询表单
const roleFormRef = ref(ElForm);
const loading = ref(false);
const total = ref(0);
const bindVisible = ref(false);
const bindingUser = ref<User>();
const isDisabled = ref(false);
const router = useRouter();
const dialog = reactive<DialogOption>({
  visible: false
});
const queryParams = reactive<PageQuery>({
  pageNum: 1,
  pageSize: 10,
  name: ""
});
const searchUserData = reactive({
  id: "",
  name: "",
  phone: "",
})
const userList = ref([]);
const searchUser = ref([]);
const formData = reactive<UserInfo>({
  countryCode: "",
  name: "",
  password: "",
  passwordType: "",
  phone: "",
});
const rules = reactive<FormRules>({
  countryCode: [{ required: true, message: "请选择您的国籍" }],
  name: [{ required: true, message: "请输入用户名", trigger: ["blur"] }],
  password: [{ required: true, message: "请输入密码", trigger: ["blur"] }],
  passwordType: [{ required: true, message: "请再次输入密码", trigger: ["blur"] }, {
    validator: function (rule, value, callback) {
      if (value !== formData.password) {
        callback(new Error("两次密码不一致!"));
      } else {
        callback();
      }
    }
  }],
  phone: [{ required: true, message: "请输入您的电话号码", trigger: ["blur"] }, {
    validator: function (rule, value, callback) {
      if (! /^\d{11}$/.test(value)) {
        callback(new Error("电话号码格式错误!"));
      } else {
        callback();
      }
    }
  }],
});

/** 
 * 打开对话框 
 */
function openDialog(id?: number) {
  dialog.visible = true;
  dialog.title = "新增用户";
}

/**  
 * 关闭对话框
*/
function closeDialog() {
  dialog.visible = false;
}


/**
 * 查询
 */
async function handleQuery() {
  loading.value = true;
  if (searchUserData.name || searchUserData.phone || searchUserData.id) {   //查询数据非空
    const params = {};
    for (let key in searchUserData) {
      if (searchUserData[key])
        params[key] = searchUserData[key];
    }
    getUserLike(params).then(res => {
      if (res.data.data.main_findUniqueuser) {
        const accute = res.data.data.main_findUniqueuser;
        searchUser.value = res.data.data.main_findManyuser.filter(item => {
          return item.id !== accute.id;
        });
        searchUser.value.unshift(accute);
      } else {
        searchUser.value = res.data.data.main_findManyuser;
      }
      total.value = searchUser.value.length;
      userList.value = searchUser.value.slice((queryParams.pageNum - 1) * queryParams.pageSize, queryParams.pageNum * queryParams.pageSize);
    }).catch(err => {
      ElMessage.error("查询失败");
    })
  } else {
    const { error, data } = await api.query({
      operationName: "System/User/GetList",
      input: convertPageQuery(queryParams, { containsFields: ["name"] })
    });
    if (!error) {
      userList.value = data!.data!;
      userList.value.forEach(item => {
        item.roles = item._join.main_findManyrole_user.map(item1 => {
          return item1._join.main_findManyrole[0].code;
        })
      })
    }
    await getAllUserNumber().then(res => {
      total.value = res.data.data.main_aggregateuser._count.id;
    })
  }
  loading.value = false;
}

/**
 * 重置查询
 */
function resetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  searchUserData.name = "";
  searchUserData.phone = "";
  searchUserData.id = "";
  handleQuery();
}


/**
 * 根据用户ID修改用户
 */
function editUser(userId: string) {
  // 打开新的页面
  router.push({
    name: "UserUpdateManage",
    query: {
      userId
    },
  });
}
/**
 * 根据用户Id删除用户
 */
function handleDelete(Id: number) {
  ElMessageBox.confirm(
    '确定要删除该用户嘛?',
    'Warning',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    // 查看该用户是否为超级管理员,如果是超级管理员，则禁止删除。

    const obj = {
      id: Id
    }
    deleteUserById(obj).then(res => {
      ElMessage.success("删除成功");
      handleQuery();
    }).catch(err => {
      ElMessage.error("删除失败");
      handleQuery();
    });
  }).catch(() => {
    ElMessage.info("已取消");
  });
}

/**
 * 清空表单数据
 */
function clearFormData() {
  formData.countryCode = "";
  formData.name = "";
  formData.password = "";
  formData.passwordType = "";
  formData.phone = "";
}

/**
 * 新增用户
 */
async function addUser() {
  roleFormRef.value.validate((valid: any) => {
    if (valid) {
      isDisabled.value = true;
      window.setTimeout(() => {
        isDisabled.value = false;
      }, 1000);
      createUser(formData).then(res => {
        ElMessage.success("成功添加一个新用户");
        closeDialog();
        clearFormData();
        handleQuery();
      }).catch(err => {
        ElMessage.error("新增失败");
        closeDialog();
      })
    }
  })
}

function bindUser(user: User) {
  bindVisible.value = true;
  bindingUser.value = user;
}

onMounted(() => {
  handleQuery(); // 初始化用户列表数据
});
</script>
