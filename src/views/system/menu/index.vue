<script lang="ts">
export default {
  name: "cmenu"
};
</script>

<script setup lang="ts">
import IconSelect from "@/components/IconSelect/index.vue";
import { Menu } from "../types";
import api from "@/api";
import { merge } from "@/utils";
import { ref, reactive } from "vue";
import { onMounted } from "vue";
import { ElForm, ElMessage, ElMessageBox } from "element-plus";
import { getSubmenu, getPerm } from "@/api/system";

defineOptions({
  name: "MenuList"
});

const menuFormRef = ref(ElForm);

const loading = ref(false);
const dialog = reactive<DialogOption>({
  visible: false
});
const dialogSubmenu = reactive<DialogOption>({
  visible: false
});

const dialogPerm = reactive<DialogOption>({
  visible: false
});
const menuList = ref<Menu[]>([]);
const tableSubmenu = ref([]); // 存放子菜单数据
const tableSubPerm = ref([]); // 存放子权限数据
const formData = reactive<Required<Omit<Menu, "id" | "parentId">>>({
  label: "",
  level: 1,
  path: "",
  icon: "",
  sort: 1
});
const editingId = ref<number>();

const rules = reactive({
  label: [{ required: true, message: "请输入菜单名称", trigger: "blur" }],
  path: [{ required: true, message: "请输入菜单路径", trigger: "blur" }]
});

// 选择表格的行菜单ID
const selectedRowMenuId = ref<number | undefined>();

/**
 * 查询
 */
async function handleQuery() {
  // 重置父组件
  loading.value = true;
  const { error, data } = await api.query({
    operationName: "System/Menu/GetMany"
  });
  if (!error) {
    menuList.value = data!.data!;
    console.log("menuList1", menuList.value);
  }
  // menuList.value.pop()
  // console.log('menuList2', menuList.value);
  // console.log('data.data' + data.data);
  loading.value = false;
}

/**
 * 行点击事件
 *
 * @param row
 */
function onRowClick(row: Menu) {
  selectedRowMenuId.value = row.id;
}

/**
 * 打开表单弹窗
 *
 */
function openDialog(menu?: Menu) {
  dialog.visible = true;
  if (menu) {
    dialog.title = "编辑菜单";
    editingId.value = menu.id;
    merge(formData, menu);
  } else {
    dialog.title = "新增菜单";
  }
}

/**
 * 菜单提交
 */
function submitForm() {
  menuFormRef.value.validate(async (isValid: boolean) => {
    if (isValid) {
      if (editingId.value) {
        const { error } = await api.mutate({
          operationName: "System/Menu/UpdateOne",
          input: {
            id: editingId.value,
            ...formData
          }
        });
        if (!error) {
          ElMessage.success("修改成功");
          closeDialog();
          handleQuery();
        }
      } else {
        const { error } = await api.mutate({
          operationName: "System/Menu/CreateOne",
          input: {
            ...formData
          }
        });
        if (!error) {
          ElMessage.success("新增成功");
          closeDialog();
          handleQuery();
        }
      }
      loading.value = false;
    }
  });
}

/**
 * 删除菜单
 */
function handleDelete(menuId: number) {
  if (!menuId) {
    ElMessage.warning("请勾选删除项");
    return false;
  }

  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(async () => {
      const { error } = await api.mutate({
        operationName: "System/Menu/DeleteOne",
        input: {
          id: menuId
        }
      });
      if (!error) {
        ElMessage.success("删除成功");
        handleQuery();
      }
    })
    .catch(() => ElMessage.info("已取消删除"));
}

/**
 * 关闭弹窗
 */
function closeDialog() {
  dialog.visible = false;
  resetForm();
}
function closeSubmenuDialog() {
  dialogSubmenu.visible = false;
}
function closePermDialog() {
  dialogPerm.visible = false;
}
/**
 * 重置表单
 */
function resetForm() {
  menuFormRef.value.resetFields();
  menuFormRef.value.clearValidate();
  formData.icon = "";
  formData.label = "";
  formData.level = 1;
  formData.path = "";
  formData.sort = 1;
}

onMounted(() => {
  handleQuery();
});

/**
 * 查看子菜单
 */
function viewSubmenu(id: number) {
  dialogSubmenu.title = "子菜单";
  // 向后端请求子菜单数据 /operations/System/Menu/GetChildrenMenus
  getSubmenu(id).then(res => {
    dialogSubmenu.visible = true;
    tableSubmenu.value = res.data.data;
  });
}
/**
 * 查看子权限
 */
function viewPerm(id: number) {
  dialogPerm.title = "子权限";
  getPerm(id).then(res => {
    dialogPerm.visible = true;
    tableSubPerm.value = res.data.data;
  });
}
/**
 * 查看子权限
 */
</script>

<template>
  <div class="app-container">
    <!-- <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="关键字" prop="keywords">
          <el-input
            v-model="queryParams.keywords"
            placeholder="菜单名称"
            clearable
            @keyup.enter="handleQuery"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery"
            ><template #icon><i-ep-search /></template>搜索</el-button
          >
          <el-button @click="resetQuery">
            <template #icon><i-ep-refresh /></template>
            重置</el-button
          >
        </el-form-item>
      </el-form>
    </div> -->
    <el-card shadow="never">
      <template #header>
        <el-button type="success" @click="openDialog()">
          <template #icon><i-ep-plus /></template>
          新增</el-button>
      </template>

      <el-table v-loading="loading" :data="menuList" highlight-current-row
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }" @row-click="onRowClick" row-key="id"
        default-expand-all border>
        <el-table-column label="菜单名称" min-width="200" prop="label" />

        <el-table-column label="菜单路径" align="center" width="150" prop="path" />

        <el-table-column label="排序" align="center" width="100" prop="sort" />

        <el-table-column fixed="right" align="center" label="操作" width="220">
          <template #default="scope">
            <el-button type="primary" link size="small" @click.stop="openDialog(scope.row)">
              <i-ep-edit />编辑
            </el-button>
            <el-button type="primary" link size="small" @click.stop="handleDelete(scope.row.id)"><i-ep-delete />
              删除
            </el-button>
            <el-button type="primary" link size="small" @click.stop="viewSubmenu(scope.row.id)"
              v-if="!scope.row.is_bottom">
              <i-ep-edit />子菜单
            </el-button>
            <el-button type="primary" link size="small" @click.stop="viewPerm(scope.row.id)" v-if="scope.row.is_bottom">
              <i-ep-edit />子权限
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog :title="dialog.title" v-model="dialog.visible" @close="closeDialog" destroy-on-close appendToBody
      width="750px">
      <el-form ref="menuFormRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="菜单名称" prop="label">
          <el-input v-model="formData.label" placeholder="请输入菜单名称" />
        </el-form-item>

        <el-form-item label="路由路径" prop="path">
          <el-input v-model="formData.path" placeholder="/system  (目录以/开头)" />
        </el-form-item>

        <el-form-item label="图标" prop="icon">
          <!-- 图标选择器 -->
          <icon-select v-model="formData.icon" />
        </el-form-item>

        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="formData.sort" style="width: 100px" controls-position="right" :min="0" />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="closeDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 点击子菜单弹出的对话框 -->
    <el-dialog :title="dialogSubmenu.title" v-model="dialogSubmenu.visible" @close="closeSubmenuDialog" destroy-on-close
      appendToBody width="750px">
      <el-table :data="tableSubmenu" style="width: 100%">
        <el-table-column label="菜单名称" width="180" prop="label" />
        <el-table-column label="菜单路径" width="180" prop="path" />
        <el-table-column label="排序" align="center" width="180" prop="sort" />
        <el-table-column fixed="right" align="center" label="操作">
          <template #default="scope">
            <el-button type="primary" link size="small" @click.stop="viewPerm(scope.row.id)">
              子权限
            </el-button>
          </template> </el-table-column>>
      </el-table>
    </el-dialog>

    <!-- 点击子权限弹出的对话框 -->
    <el-dialog :title="dialogPerm.title" v-model="dialogPerm.visible" @close="closePermDialog" destroy-on-close
      appendToBody width="750px">
      <el-table :data="tableSubPerm" style="width: 100%">
        <el-table-column label="创建时间" width="200" prop="createdAt" />
        <el-table-column label="是否启用" align="center" width="180">
          <template #default="scope">
            {{ scope === 0 ? "否" : "是" }}
          </template>
        </el-table-column>
        <el-table-column label="方法" align="center" width="100" prop="method" />
        <el-table-column label="路径" width="240" prop="path" />
      </el-table>
    </el-dialog>
  </div>
</template>
