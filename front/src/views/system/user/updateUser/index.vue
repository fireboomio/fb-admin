<template>
  <div class="app-container">
    <el-card>
      <template #header>
        修改用户信息
      </template>
      <el-form label-position="top" label-width="100px" :model="formData" style="max-width: 460px">
        <el-form-item label="用户名">
          <el-input v-model="formData.name" />
        </el-form-item>
        <el-form-item label="电话">
          <el-input v-model="formData.phone" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="formData.password" type="password" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitUpdateUser">修改</el-button>
        </el-form-item>
      </el-form>
      <template #footer>

      </template>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from 'vue';
import { useRoute } from 'vue-router';
import { updateUser, getUserByUserId } from "@/api/user";
import { ElMessage } from 'element-plus';
const formData = reactive({
  name: "",
  password: "",
  passwordtype: "",
  phone: "",
  userId: "",
})
defineOptions({
  name: "UserUpdateManage"
});
const route = useRoute();
function submitUpdateUser() {
  updateUser(formData).then(res => {
    ElMessage.success("修改成功");
  }).catch(error => {
    ElMessage.error("修改失败");
  })
}
onMounted(() => {
  getUserByUserId(route.query?.userId).then(res => {
    formData.name = res.data.data.main_findManyuser[0].name;
    formData.phone = res.data.data.main_findManyuser[0].phone;
    formData.userId = res.data.data.main_findManyuser[0].user_id;
  })
});
</script>
