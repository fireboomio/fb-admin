<script lang="ts">
export default {
  name: "roleBindApi"
};
</script>
<template>
  <!-- 角色绑定API弹窗 -->
  <el-dialog title="绑定权限" :model-value="modelValue" width="600px" append-to-body
    @update:model-value="$emit('update:modelValue', $event)" @close="close">
    <el-table v-loading="loading" :data="apis" ref="tableRef" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="路径" prop="title" />
      <el-table-column label="Method" align="center" prop="method" />
      <el-table-column label="实时查询" align="center" prop="liveQuery" />
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
import type { Role, API } from "../types";
import { useUserStoreHook } from "@/store/modules/user";
import { getRolePerms } from "@/api/system";
const store = useUserStoreHook()
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  role: {
    type: Object as PropType<Role>
  },
  roles: {
    type: Array as PropType<Role[]>,
    required: true
  }
});
const tableRef = ref<InstanceType<typeof ElTable>>();
const emit = defineEmits(["update:modelValue"]);
const loading = ref(false);
const apis = ref<API[]>([]);
const selections = ref<number[]>([]);
let originApis: number[] = [];

function handleSelectionChange(_selections: API[]) {
  selections.value = _selections.map(item => +item.id!);
}

function close() {
  emit("update:modelValue", false);
  loading.value = false;
  selections.value = [];
  originApis = [];
}

async function onSubmit() {
  loading.value = true;
  const { error } = await api.mutate({
    operationName: "System/Role/BindRoleApis",
    input: {
      allRoles: props.roles.map(item => item.code!),
      apis: selections.value,
      roleCode: props.role!.code!
    }
  });
  if (!error) {
    const roles = store.roles;
    getRolePerms(roles).then(res => {
      store.SET_PERMISSIONS(res);
    });
    ElMessage.success("绑定成功");
    close();
  } else {
    ElMessage.error("绑定失败");
    loading.value = false;
  }
}

watchEffect(async () => {
  if (props.modelValue && props.role?.code) {
    loading.value = true;
    const res1 = await api.query({
      operationName: "System/Operation/GetMany"
    });
    if (!res1.error) {
      apis.value = res1.data!.data!;
      const { error, data } = await api.query({
        operationName: "System/Role/GetRoleBindApis",
        input: {
          code: props.role.code
        }
      });
      if (!error) {
        originApis = data!.data?.map(item => item.id!) ?? [];
        selections.value = originApis;
        const selectedApis = apis.value.filter(item =>
          originApis.includes(+item.id!)
        );
        for (const api of selectedApis) {
          tableRef.value!.toggleRowSelection(api, true);
        }
      }
    }
    loading.value = false;
  }
});
</script>
