<template>
  <!-- 绑定角色弹窗 -->
  <el-dialog title="绑定角色" :model-value="modelValue" width="600px" append-to-body
    @update:model-value="$emit('update:modelValue', $event)" @close="close">
    <el-table v-loading="loading" :data="roles" ref="tableRef" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column key="code" label="Code" align="center" prop="code" width="200" />
      <el-table-column key="remark" label="描述" align="center" prop="remark" />
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
import { PropType } from "vue";
import type { Role, User } from "../types";
import { ref, watchEffect } from "vue";

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  user: {
    type: Object as PropType<User>
  }
});
const tableRef = ref<InstanceType<typeof ElTable>>();
const emit = defineEmits(["update:modelValue"]);
const loading = ref(false);
const roles = ref<Role[]>([]);
const selections = ref<number[]>([]);
let originRoles: number[] = [];

function handleSelectionChange(_selections: Role[]) {
  selections.value = _selections.map(item => item.id!);
}

function close() {
  emit("update:modelValue", false);
  loading.value = false;
  selections.value = [];
  originRoles = [];
}

async function onSubmit() {
  loading.value = true;
  // 先移除
  for (const originRoleCode of originRoles) {
    await api.mutate({
      operationName: "System/User/DisconnectRole",
      input: {
        userId: props.user!.id!,
        roleId: originRoleCode
      }
    });
  }
  // 再添加
  for (const id of selections.value) {
    await api.mutate({
      operationName: "System/User/ConnectRole",
      input: {
        userId: props.user!.id!,
        roleId: id
      }
    });
  }
  ElMessage.success("绑定完成");
  close();

}

watchEffect(async () => {
  if (props.modelValue && props.user) {
    loading.value = true;
    const res1 = await api.query({
      operationName: "System/Role/GetMany"
    });
    if (!res1.error) {
      roles.value = res1.data!.data!;
      const { error, data } = await api.query({
        operationName: "System/User/GetUserRole",
        input: {
          userId: props.user!.id!
        }
      });
      if (!error) {
        originRoles = data!.data!.map(item => item.id!) || [];
        selections.value = originRoles;
        const selectedRoles = roles.value.filter(item =>
          originRoles.includes(item.id!)
        );
        for (const role of selectedRoles) {
          tableRef.value!.toggleRowSelection(role, true);
        }
      }
    }
    loading.value = false;
  }
});
</script>
