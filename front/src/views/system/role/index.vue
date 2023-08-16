<script setup lang="ts">
import api, { convertPageQuery } from "@/api";
import { merge } from "@/utils";
import {
  ElForm,
  ElMessage,
  ElMessageBox,
  ElTree,
  FormRules,
  ElPagination
} from "element-plus";
import { Role } from "../types";
import RoleBindApi from "./api.bind.vue";
import RoleBindMenu from "./menu.bind.vue";
import { ref, reactive, onMounted, nextTick } from "vue";
import { Icon } from '@iconify/vue';
import { hasAuth, getAuths } from "@/router/utils";
import { useUserStoreHook } from "@/store/modules/user";
import { deleteRoles, getRolePerms, getRoleUsers, roleMenuTreeselect, updateRolePermAdd, updateRolePermRemove } from "@/api/system";
import { sessionKey } from "@/utils/auth";
import { storageSession } from "@pureadmin/utils";
import router from "@/router";

defineOptions({
  name: "RoleManage"
});
const store = useUserStoreHook();
const queryFormRef = ref(ElForm);
const roleFormRef = ref(ElForm);
const menuRef = ref(ElTree);
const menuOptions = ref([]);
const loading = ref(false);
const ids = ref<number[]>([]);
const total = ref(0);
const userInfoById = ref([]);
const queryParams = reactive<PageQuery>({
  pageNum: 1,
  pageSize: 10,
  code: "",
});
const roleList = ref<Role[]>([]);
const dialog = reactive<DialogOption>({
  visible: false
});
let dialogEnterLeave = ""
let checkedKeys = [];
let newCheckedKeys = [];
let allKeys = [];
const formData = reactive({
  code: "",
  remark: "",
  menuCheckStrictly: false,
})
const editingId = ref<number>();
const rules = reactive<FormRules>({
  code: [{ required: true, message: "请输入角色编码", trigger: ['blur'] },
  {
    validator: function (rule, value, callback) {
      if (! /^[a-zA-Z][a-zA-Z0-9_]*$/.test(value)) {
        callback(new Error('角色只能以字母开头, 并且只能包含字母、数字、_'));
      } else {
        callback();
      }
    }
  }]
});
const apiBindVisible = ref(false);
const menuBindVisible = ref(false);
const bindingRole = ref<Role>();
const menuExpand = ref(false);
const menuNodeAll = ref(false);
const menu = ref(null)
defineExpose({
  menu,
});

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

/**
 * checkbox change事件
 */
function handleSelectionChange(selection: any) {
  ids.value = selection.map((item: any) => item.id);
}

/**
 * 打开角色表单弹窗
 */
function openDialog(role?: Role) {
  dialogEnterLeave = "dialog-from-right";
  dialog.visible = true;
  if (role) {
    dialog.title = "修改角色";
    merge(formData, role);
    editingId.value = role.id;
    getRoleMenuTreeselect(role.id);
  } else {
    editingId.value = 0;
    roleMenuTreeselect(1).then(response => {
      menuOptions.value = response.data.menus;
      checkedKeys = [];
    })
    dialog.title = "新增角色";
  }
}

/**
 * 角色表单提交
 */
async function handleSubmit(code) {
  loading.value = true;
  roleFormRef.value.validate(async (valid: any) => {
    if (valid) {
      if (editingId.value) {
        // 判断是权限是增加、减少、没有改变
        const menuIdsAdd = [];
        const menuIdsRemove = [];
        newCheckedKeys.forEach(item => {
          if (!checkedKeys.includes(item)) {
            menuIdsAdd.push(item)
          }
        })
        checkedKeys.forEach(item => {
          if (!newCheckedKeys.includes(item)) {
            menuIdsRemove.push(item)
          }
        })
        // 发送增加权限的请求
        if (menuIdsAdd.length) {
          await updateRolePermAdd(code, editingId.value, menuIdsAdd);
        }
        // 发送减少权限的请求
        if (menuIdsRemove.length) {
          await updateRolePermRemove(code, editingId.value, menuIdsRemove);
        }
        if (!menuIdsAdd.length && !menuIdsRemove.length) {
          ElMessage.warning("您没有做任何修改");
          closeDialog();
        } else {
          // 修改本地以及store中的权限perm
          getRolePerms(store.roles).then(res => {
            store.SET_PERMISSIONS(res.data);
            ElMessage.success("修改成功");
            closeDialog();
            resetQuery();
          })
        }
      } else {
        // 角色重复性校验
        const verify = roleList.value.some(item => item.code === formData.code);
        if (verify) {
          ElMessage.warning("该角色已存在,请勿重复添加");
        } else {
          const { data, error } = await api.mutate({
            operationName: "System/Role/AddOne",
            input: {
              ...formData
            }
          });
          // 判断是权限是增加、减少、没有改变
          const menuIdsAdd = [];
          const menuIdsRemove = [];
          newCheckedKeys.forEach(item => {
            if (!checkedKeys.includes(item)) {
              menuIdsAdd.push(item)
            }
          })
          checkedKeys.forEach(item => {
            if (!newCheckedKeys.includes(item)) {
              menuIdsRemove.push(item)
            }
          })
          // 发送增加权限的请求
          if (menuIdsAdd.length) {
            await updateRolePermAdd(code, data.data.id, menuIdsAdd);
          }
          // 发送减少权限的请求
          if (menuIdsRemove.length) {
            await updateRolePermRemove(code, data.data.id, menuIdsRemove);
          }
          if (!error) {
            ElMessage.success("新增成功");
            closeDialog();
            resetQuery();
          }
        }
      }
    }
  });
  loading.value = false;
}

/**
 * 关闭弹窗
 */
function closeDialog() {
  dialogEnterLeave = "dialog-leave";
  dialog.visible = false;
  resetForm();
}

/**
 * 重置表单
 */
function resetForm() {
  roleFormRef.value.resetFields();
  roleFormRef.value.clearValidate();

  formData.code = "";
  formData.remark = "";
}

/**
 * 删除
 */
function handleDelete(roleId?: number) {
  const roleIds = roleId ? [roleId] : ids.value;
  if (!roleIds) {
    ElMessage.warning("请勾选删除项");
    return;
  }
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    loading.value = true;
    deleteRoles(roleIds).then(res => {
      if (res) {
        ElMessage.success("删除成功");
        resetQuery();
      }
    })
    loading.value = false;
  });
}


/**
 * 根据角色id查看用户信息
 */
function getRoleUser(code: string) {
  // 打开用户管理的页面，并展示当前角色下的用户
  router.push({
    name: "UserManage",
    query: {
      code
    }
  });
}

onMounted(() => {
  handleQuery();
});

/** 根据角色ID查询菜单树结构 */
function getRoleMenuTreeselect(roleId?: number) {
  roleMenuTreeselect(roleId).then(response => {
    menuOptions.value = response.data.menus;
    checkedKeys = response.data.checkedKeys;
    menuNodeAll.value = checkedKeys.length === response.data.allKeys.length;
    allKeys = response.data.allKeys;
    nextTick(() => {
      response.data.checkedKeys.forEach(item => {
        menu.value.setChecked(item, true);
      });
    })
    return response;
  });
}

// 树权限（展开/折叠）
function handleCheckedTreeExpand(value) {
  let treeList = menuOptions.value;
  for (let i = 0; i < treeList.length; ++i) {
    menu.value.store.nodesMap[treeList[i].id].expanded = value;
  }
}

// 树权限（全选/全不选）
function handleCheckedTreeNodeAll(value) {
  menu.value.setCheckedNodes(value ? menuOptions.value : [])
  if (value) {
    newCheckedKeys = allKeys;
  } else {
    newCheckedKeys = []
  }
}
// 树权限（父子联动）
function handleCheckedTreeConnect(event) {
  formData.menuCheckStrictly = event;
}
function currentChecked(nodeObj, SelectedObj) {
  newCheckedKeys = SelectedObj.checkedKeys;
  menuNodeAll.value = newCheckedKeys.length === allKeys.length;
}
</script>

<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item prop="code" label="编码">
          <el-input v-model="queryParams.code" placeholder="角色编码" clearable @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery">
            <Icon icon="ep:search" />搜索
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="resetQuery">
            <Icon icon="ep:refresh" />重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <template #header>
        <Auth value="system:role:add">
          <el-button type="success" @click="openDialog()">
            <Icon icon="ep:plus" />新增
          </el-button>
        </Auth>
        <Auth value="system:role:remove">
          <el-button type="danger" :disabled="ids.length === 0" @click="handleDelete()">
            <Icon icon="ep:delete" />删除
          </el-button>
        </Auth>
      </template>
      <el-table ref="dataTableRef" v-loading="loading" :data="roleList" @selection-change="handleSelectionChange"
        highlight-current-row border>
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="角色编码" prop="code" width="200" />
        <el-table-column label="角色描述" prop="remark" width="300" />
        <el-table-column fixed="right" label="操作" align="center">
          <template #default="scope">
            <Auth value="system:role:edit">
              <el-button type="primary" size="small" link @click="openDialog(scope.row)">
                <Icon icon="ep:edit" />编辑
              </el-button>
            </Auth>
            <Auth value="system:role:remove">
              <el-button type="primary" size="small" link @click="handleDelete(scope.row.id)">
                <Icon icon="ep:delete" />删除
              </el-button>
            </Auth>
            <el-button type="primary" size="small" link @click="getRoleUser(scope.row.code)">
              <Icon icon="ep:search" />查看用户
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-if="total > 0" v-model:total="total" v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize" @pagination="handleQuery" />
    </el-card>

    <!-- 角色表单弹窗 -->
    <el-dialog :title="dialog.title" v-model="dialog.visible" width="500px" @close="closeDialog" :class="dialogEnterLeave"
      :show-close="false" :height="'100vh'">
      <el-form ref="roleFormRef" :model="formData" :rules="rules" label-width="100px" class="dialog-content-left">
        <el-form-item label="角色编码" prop="code">
          <el-input v-model="formData.code" :readonly="!!editingId" placeholder="请输入角色编码" />
        </el-form-item>
        <el-form-item label="角色描述" prop="remark">
          <el-input v-model="formData.remark" placeholder="请输入角色描述" />
        </el-form-item>
        <el-form-item label="菜单权限" v-if="dialog.title !== '新增角色'">
          <el-checkbox v-model="menuExpand" @change="handleCheckedTreeExpand($event)">展开/折叠</el-checkbox>
          <el-checkbox v-model="menuNodeAll" @change="handleCheckedTreeNodeAll($event)">全选/全不选</el-checkbox>
        </el-form-item>
        <el-form-item>
          <el-tree class="tree-border" :data="menuOptions" show-checkbox ref="menu" node-key="id" empty-text="......"
            @check="currentChecked" style="display: block;">
          </el-tree>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit(formData.code)">确 定</el-button>
          <el-button @click="closeDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style>
.dialog-from-right {
  height: 100vh;
  width: 25%;
  margin: 0 auto !important;
  overflow: auto;
  animation: dialogFromRightEnter 0.5s;
  position: fixed;
  right: 0;
  top: 0;
  bottom: 0;
}

.dialog-leave {
  height: 100vh;
  width: 25%;
  margin: 0 auto !important;
  overflow: auto;
  position: fixed;
  right: 0;
  top: 0;
  bottom: 0;
  animation: dialogFromRightLeave 0.5s;
}

@keyframes dialogFromRightEnter {
  from {
    opacity: 0;
    transform: translateX(100%);
  }

  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes dialogFromRightLeave {
  from {
    opacity: 1;
    transform: translateX(0);
  }

  to {
    opacity: 0;
    transform: translateX(100%);
  }
}
</style>
