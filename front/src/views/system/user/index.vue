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
        <el-form-item label="用户ID" prop="name">
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
        <el-form-item>
          <el-dropdown trigger="click" @command="searchUserByRole">
            <el-button type="primary">
              {{ currentSelectRole }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item v-for="role in roleList" :key="role.id" :command="role">{{
                  role.remark }}</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
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
    <role-bind v-model="bindVisible" :user="bindingUser" @update:modelValue="getBindSuccess" />
    <!-- 新增用户 -->
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
import { createUser, deleteUserById, getAllUserNumber, getUserLike, UserInfo } from "@/api/user";
import { getRoleUsers } from "@/api/system";
import { country } from "./utils/country"
import { useRouter, useRoute } from 'vue-router';
import { Role } from "../types";
import { Md5 } from 'ts-md5';

const salt = "qwertyuiop";

const route = useRoute();
const queryFormRef = ref(ElForm); // 查询表单
const roleFormRef = ref(ElForm);
const loading = ref(false);
const total = ref(0);
const bindVisible = ref(false);
const bindingUser = ref<User>();
const isDisabled = ref(false);
const router = useRouter();
const userInfoById = ref([]);  // 用来接收角色对应用户的数组
const roleList = ref<Role[]>([]);  // 用来接收所有的角色信息在下拉菜单处展示
const currentSelectRole = ref("选择角色");   // 当前选择的角色名称
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
  password: [{ required: true, message: "请输入密码", trigger: ["blur"] }, {
    validator: function (rule, value, callback) {
      if (!/^[a-zA-Z]\w{5,17}$/.test(value)) {
        callback(new Error("密码格式应以字母开头，长度在6~18之间，只能包含字母，数字和下划线"));
      } else {
        if (formData.passwordType !== "") {
          roleFormRef.value.validateField('passwordType');

        }
        callback();
      }
    }
  }],
  passwordType: [{ required: true, message: "请再次输入密码", trigger: ["change"] }, {
    validator: function (rule, value, callback) {
      if (value !== formData.password) {
        callback(new Error("两次密码不一致!"));
      } else {
        callback();
      }
    }
  }],
  phone: [{ required: true, message: "请输入您的电话号码", trigger: "blur" }, {
    validator: function (rule, value, callback) {
      if (! /^\d{11}$/.test(value)) {
        callback(new Error("电话号码格式错误!"));
      } else {
        callback();
      }
    }
  }],
});
function getBindSuccess(val) {
  if (!val) {
    handleQuery();
  }
}
/** 
 * 打开对话框 
 */
function openDialog(id?: number) {
  dialog.visible = true;
  dialog.title = "新增用户";
}

function resetformData() {
  formData.countryCode = ""
  formData.name = ""
  formData.password = ""
  formData.passwordType = ""
  formData.phone = ""
  roleFormRef.value.resetFields();
}
/**  
 * 关闭对话框
*/
function closeDialog() {
  resetformData();
  dialog.visible = false;
}
/**
 * 根据角色查用户
 */
async function searchUserByRole(role: Role) {
  currentSelectRole.value = role.remark;
  await getRoleUsers(role.code).then(res => {
    userInfoById.value = res.data.data.data[0].role.main_findManyadmin_role_user.map(item => {
      return {
        id: item._join.user[0].id,
        name: item._join.user[0].name,
        avatar: item._join.user[0].avatar,
        createdAt: item._join.user[0].created_at,
      };
    })
    userInfoById.value.forEach(item => {
      item.roles = [res.data.data.data[0].code];
    })
    total.value = userInfoById.value.length;
  })
  userList.value = userInfoById.value.slice((queryParams.pageNum - 1) * queryParams.pageSize, queryParams.pageNum * queryParams.pageSize);
}
/**
 * 查询
 */
async function handleQuery() {
  loading.value = true;
  const { error, data } = await api.query({
    operationName: "System/Role/GetList",
    input: convertPageQuery(queryParams, { containsFields: ["code", "remark"] })
  });
  if (!error) {
    roleList.value = data!.data!;
  }
  if (route.query?.code) {    // 判断是否由角色管理跳转，若是则只展示相应角色的数据给用户
    currentSelectRole.value = String(route.query?.code);

    await getRoleUsers(route.query?.code).then(res => {
      const userInfo = res.data.data.data[0].role.main_findManyadmin_role_user.filter(item => {
        return item._join.user.length;
      })
      userInfoById.value = userInfo.map(item => {
        return {
          id: item._join.user[0].id,
          name: item._join.user[0].name,
          avatar: item._join.user[0].avatar,
          createdAt: item._join.user[0].created_at,
        };
      })

      userInfoById.value.forEach(item => {
        item.roles = [res.data.data.data[0].code];
      })
      total.value = userInfoById.value.length;
    })
    userList.value = userInfoById.value.slice((queryParams.pageNum - 1) * queryParams.pageSize, queryParams.pageNum * queryParams.pageSize);
  } else if (currentSelectRole.value !== "选择角色") {
    userList.value = userInfoById.value.slice((queryParams.pageNum - 1) * queryParams.pageSize, queryParams.pageNum * queryParams.pageSize);
  } else {
    if (searchUserData.name || searchUserData.phone || searchUserData.id) {   //查询数据非空
      const params = {};
      for (let key in searchUserData) {
        if (searchUserData[key])
          params[key] = searchUserData[key];
      }
      getUserLike(params).then(res => {
        if (res.data.data.data1) {
          const accute = res.data.data.data1;
          searchUser.value = res.data.data.data.filter(item => {
            return item.id !== accute.id;
          });
          searchUser.value.unshift(accute);
          if (!searchUserData.name && !searchUserData.phone) {
            searchUser.value.length = 0;
            searchUser.value.unshift(accute);
          }
        } else {
          searchUser.value = res.data.data.data;
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
          item.roles = item._join.main_findManyadmin_role_user.map(item1 => {
            return item1._join.main_findManyadmin_role[0].code;
          })
        })
      }
      await getAllUserNumber().then(res => {
        total.value = res.data.data.data._count.id;
      })
    }
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
      const md5: any = new Md5();
      md5.appendAsciiStr(formData.password[0] + salt[5] + formData.password.slice(1, formData.password.length));
      formData.password = md5.end();
      window.setTimeout(() => {
        isDisabled.value = false;
      }, 1000);
      createUser(formData).then(res => {
        ElMessage.success("成功添加一个新用户");
        closeDialog();
        clearFormData();
        handleQuery();
        resetformData();
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
