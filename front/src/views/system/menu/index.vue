<script setup lang="ts">
import IconSelect from "@/components/IconSelect/index.vue";
import { Menu } from "../types";
import api from "@/api";
import { merge } from "@/utils";
import { ref, reactive, nextTick } from "vue";
import { onMounted } from "vue";
import { ElForm, ElMessage, ElMessageBox } from "element-plus";
import { GetApiList, getBindAPI, getMenuPerms } from "@/api/system";
import { Icon } from '@iconify/vue';
import { handleTree, is } from "@pureadmin/utils";
import { parseTime } from "@/utils/date";
defineOptions({
  name: "MenuManage"
});

const menuFormRef = ref(ElForm);

const loading = ref(false);
const isExpandAll = ref(false);
const refreshTable = ref(true);
let apiId = ref()
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
const menuData = ref([]);
const menuTree = [{ id: 0, label: "主目录", api_id: "", createTime: "", icon: "", is_bottom: 0, level: 0, menu_typr: "M", parentId: null, path: "/", perms: "", sort: 1, children: [] }];
const formData = reactive({
  apiId: 0,
  label: "",
  level: 1,
  path: "",
  icon: "",
  sort: 1,
  parentId: 0,
  menuType: "M",
  perms: "",
});
const editingId = ref<number>();
const apiList = ref([]); // 存放api信息
const rules = reactive({
  parentId: [{ required: true, trigger: "blur" }],
  label: [{ required: true, message: "请输入菜单名称", trigger: "blur" }],
  path: [{ required: false, message: "请输入菜单路径", trigger: "blur" }]
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
    let datasource = data!.data!;
    const menu_type = ["目录", "菜单", "按钮"];
    datasource.forEach(item => {
      if (item.menu_type === "M") {
        item.menu_type = menu_type[0];
      } else if (item.menu_type === "C") {
        item.menu_type = menu_type[1];
      } else {
        item.menu_type = menu_type[2];
      }
    })
    menuTree[0].children = handleTree(datasource, "id");
    menuData.value = datasource;
    menuList.value = menuTree[0].children;

  }
  loading.value = false;
}

/**
 * 行点击事件
 * @param row
 */
function onRowClick(row: Menu) {
  selectedRowMenuId.value = row.id;
}

/**
 * 增加
 */
function openDialogAdd(menu?: Menu) {
  formData.perms = "";
  editingId.value = 0;
  dialog.visible = true;
  dialog.title = "新增菜单";
  if (!menu) { // 一级菜单
    formData.level = 1;
    formData.parentId = 0;
  } else {     // 二级菜单
    formData.level = menu.level + 1;
    formData.parentId = menu.id;
  }
}
function openDialogUpdate(menu?: Menu) {
  dialog.visible = true;
  dialog.title = "修改菜单";
  editingId.value = menu.id;
  merge(formData, menu);
  formData.menuType = menu.menu_type;
  if (!formData.parentId) {
    formData.parentId = 0;
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
        // 判断新增的菜单是否重复
        const verify = menuList.value.some((item) => item.label === formData.label)
        if (verify) {
          ElMessage.warning("该菜单已存在, 请勿重复添加");
        } else {
          formData.apiId = apiId.value;
          // 清空apiId
          apiId = ref();
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
      }
      loading.value = false;
    }
  });
}


/**  
 * 查询当前菜单下是否还有子菜单
 */
function judgeHasChildren(id: number) {
  return menuData.value.some(item => {
    return item.parentId === id;
  })
}
/**
 * 删除菜单
 */
function handleDelete(menuId: number) {
  if (judgeHasChildren(menuId)) {
    ElMessage.warning("无法删除父菜单")
  } else {
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
  GetApiList().then(res => {
    apiList.value = res.data.data.data;
  }).catch(err => {
    ElMessage.error("获取信息失败!");
  })
});

/** 展开/折叠操作 */
function toggleExpandAll() {
  refreshTable.value = false;
  isExpandAll.value = !isExpandAll.value;
  nextTick(() => {
    refreshTable.value = true;
  });
}
</script>

<template>
  <div class="app-container">
    <el-card shadow="never">
      <template #header>
        <Auth value="system:menu:add">
          <el-button type="success" @click="openDialogAdd()">
            <template #icon>
              <Icon icon="ep:plus" />
            </template>
            新增</el-button>
        </Auth>
        <el-button type="info" @click="toggleExpandAll">
          <Icon icon="ep:sort" />展开/折叠
        </el-button>
      </template>


      <el-table v-if="refreshTable" v-loading="loading" :data="menuList" row-key="id" :default-expand-all="isExpandAll"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }">
        <el-table-column prop="label" label="菜单名称" :show-overflow-tooltip="true" width="160"></el-table-column>
        <el-table-column prop="menu_type" label="菜单类型" align="center" :show-overflow-tooltip="true"
          width="160"></el-table-column>
        <el-table-column prop="icon" label="图标" align="center" width="100"></el-table-column>
        <el-table-column prop="sort" label="排序" width="60" align="center"></el-table-column>
        <!-- <el-table-column prop="level" label="层级" width="60" align="center"></el-table-column> -->
        <el-table-column prop="perms" label="权限标识" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column prop="path" label="组件路径" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="创建时间" align="center" prop="create_time">
          <template #default="scope">
            <span v-if="scope.row.create_time">{{ parseTime(scope.row.create_time, "{y}-{m}-{d} {h}:{i}:{s}") }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="300">
          <template #default="scope">
            <Auth value="system:menu:edit">
              <el-button size="small" link type="primary" icon="el-icon-edit" @click.stop="openDialogUpdate(scope.row)">
                <Icon icon="ep:edit" />修改
              </el-button>
            </Auth>
            <Auth value="system:menu:add">
              <el-button size="small" link type="primary" icon="el-icon-plus" @click.stop="openDialogAdd(scope.row)">
                <Icon icon="ep:plus" />新增
              </el-button>
            </Auth>
            <Auth value="system:menu:edit">
              <el-button size="small" link type="primary" icon="el-icon-delete" @click="handleDelete(scope.row.id)">
                <Icon icon="ep:delete" />删除
              </el-button>
            </Auth>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!--@@@@@ openDialog @@@@@-->
    <el-dialog :title="dialog.title" v-model="dialog.visible" @close="closeDialog" destroy-on-close appendToBody
      width="750px">
      <el-form ref="menuFormRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="上级菜单">
          <el-tree-select v-model="formData.parentId" value-key="id" :data="menuTree" check-strictly
            :render-after-expand="false" />
        </el-form-item>
        <el-form-item label="菜单名称" prop="label">
          <el-input v-model="formData.label" placeholder="请输入菜单名称" />
        </el-form-item>

        <el-col :span="24">
          <el-form-item label="菜单类型" prop="menuType">
            <el-radio-group v-model="formData.menuType">
              <el-radio label="M">目录</el-radio>
              <el-radio label="C">菜单</el-radio>
              <el-radio label="F">按钮</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>
        <el-col :span="24" v-if="formData.menuType !== 'F'">
          <el-form-item label="图标" prop="icon">
            <!-- 图标选择器 -->
            <icon-select v-model="formData.icon" />
          </el-form-item>
        </el-col>
        <el-col v-if="formData.menuType !== 'F'">
          <el-form-item label="组件路径" prop="path">
            <el-input v-model="formData.path" placeholder="/system  (目录以/开头)" />
          </el-form-item>
        </el-col>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="formData.sort" style="width: 100px" controls-position="right" :min="0" />
        </el-form-item>
        <el-col v-if="formData.menuType !== 'M'">
          <el-form-item label="权限标识" prop="perms">
            <el-input v-model="formData.perms" placeholder="请输入权限标识" />
          </el-form-item>
        </el-col>
        <el-col v-if="formData.menuType === 'F'">
          <el-form-item label="绑定飞布Api" prop="api_id">
            <el-select v-model="apiId" placeholder="请选择飞布Api接口">
              <el-option v-for="item in apiList" :key="item.id" :label="item.path" :value="item.id" />
            </el-select>
          </el-form-item>
        </el-col>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <Auth value="system:menu:edit">
            <el-button type="primary" @click="submitForm">确 定</el-button>
          </Auth>
          <Auth value="system:menu:edit">
            <el-button @click="closeDialog">取 消</el-button>
          </Auth>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
