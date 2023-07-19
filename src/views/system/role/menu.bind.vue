<script lang="ts">
export default {
  name: "roleBindMenu"
};
</script>
<template>
  <!-- 角色绑定菜单弹窗 -->
  <el-dialog title="绑定菜单" :model-value="modelValue" width="600px" append-to-body
    @update:model-value="$emit('update:modelValue', $event)" @close="close">
    <el-table v-loading="loading" :data="menus" ref="tableRef" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="名称" prop="label" />
      <el-table-column label="路径" align="center" prop="path" />
    </el-table>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="close">取消</el-button>
        <el-button type="primary" @click="onSubmit">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import api from "@/api";
import { ElTable, ElMessage } from "element-plus";
import { PropType, ref, watchEffect } from "vue";
import type { Role, Menu } from "../types";

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  role: {
    type: Object as PropType<Role>
  }
});
const tableRef = ref<InstanceType<typeof ElTable>>();
const emit = defineEmits(["update:modelValue"]);
const loading = ref(false);
const menus = ref<Menu[]>([]);
const selections = ref<number[]>([]);
let originMenus: number[] = [];

function handleSelectionChange(_selections: Menu[]) {
  selections.value = _selections.map(item => item.id!);
}

function close() {
  emit("update:modelValue", false);
  loading.value = false;
  selections.value = [];
  originMenus = [];
}

async function onSubmit() {
  loading.value = true;
  // 先移除
  for (const menuId of originMenus) {
    await api.mutate({
      operationName: "System/Role/DisconnectOneMenu",
      input: {
        id: props.role!.id!,
        menuId
      }
    });
  }
  // 再添加
  for (const menuId of selections.value) {
    await api.mutate({
      operationName: "System/Role/ConnectOneMenu",
      input: {
        id: props.role!.id!,
        menuId
      }
    });
  }
  close();
  ElMessage.success("绑定成功");
  loading.value = false;
}

watchEffect(async () => {
  if (props.modelValue && props.role?.code) {
    loading.value = true;
    const res1 = await api.query({
      operationName: "System/Menu/GetMany",
      input: {
        in: [1, 2]
      }
    });
    if (!res1.error) {
      menus.value = res1.data!.data!;
      const { error, data } = await api.query({
        operationName: "System/Role/GetBindMenus",
        input: {
          roleId: props.role.id!
        }
      });
      if (!error) {
        originMenus = data!.data?.map(item => item.id!) ?? [];
        selections.value = originMenus;
        const selectedMenus = menus.value.filter(item =>
          originMenus.includes(item.id!)
        );
        for (const role of selectedMenus) {
          tableRef.value!.toggleRowSelection(role, true);
        }
      }
    }
    loading.value = false;
  }
});
</script>
