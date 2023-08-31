<template>
  <div class="app-container">
    <el-card>
      <template #header>
        修改用户信息
      </template>
      <el-form ref="userFormRef" label-position="top" label-width="100px" :model="formData" style="max-width: 400px"
        :rules="rules">
        <el-form-item label="用户名" prop="name">
          <el-input v-model="formData.name" style="width: 230px" />
        </el-form-item>
        <el-form-item label="头 像">
          <el-upload class="avatar-uploader" action="#" :show-file-list="false" accept="image/jpg,image/jpeg,image/png"
            :before-upload="beforeAvatarUpload" :on-change="changeImage" :auto-upload="false">
            <img v-if="imageUrl" :src="imageUrl" class="avatar" title="点击重新上传" />
            <el-icon v-else class="avatar-uploader-icon">
              <plus />
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="手机号码" prop="phone">
          <el-input v-model="formData.phone" style="width: 300px" />
        </el-form-item>
        <el-form-item label="密 码" prop="password">
          <el-input v-model="formData.password" type="password" style="width: 300px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitUpdateUser">修 改</el-button>
          <el-button @click="close">取 消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import { updateUser, getUserByUserId } from "@/api/user";
import { ElForm, ElMessage } from 'element-plus';
import { useTags } from "@/layout/hooks/useTag";
import { useMultiTagsStoreHook } from "@/store/modules/multiTags";
import router from "@/router";
import { message } from '@/utils/message';
import { FormRules } from 'element-plus';
import { formatToken, getToken } from '@/utils/auth';
import axios from 'axios';
import { Md5 } from 'ts-md5';

const salt = "qwertyuiop";
const rules = reactive<FormRules>({
  name: [{ required: true, message: "用户名不能为空!" }, {
    validator: function (rule, value, callback) {
      if (!/^[a-zA-Z0-9_-]{3,30}$/.test(value)) {
        callback(new Error("用户名由3-30个字符、支持中英文、数字、”_“或减号"));
      } else {
        callback();
      }
    }
  }],
  password: [{ required: false }, {
    validator: function (rule, value, callback) {
      if (!/^[a-zA-Z]\w{5,17}$/.test(value)) {
        callback(new Error("密码由以字母开头,长度在6~18之间,只能包含字母、数字和下划线"));
      } else {
        callback();
      }
    }
  }],
  phone: [{ required: true, message: "电话号码不能为空" }, {
    validator: function (rule, value, callback) {
      if (!/^1[3-9]\d{9}$/.test(value)) {
        callback(new Error("电话号码非法"));
      } else {
        callback();
      }
    }
  }]
});
const File = ref();
const imageUrl = ref('');
const { multiTags } = useTags();
const userFormRef = ref(ElForm);
const formData = reactive({
  id: 0,
  avatar: "",
  name: "",
  password: "",
  phone: "",
  userId: "",
})
defineOptions({
  name: "UserUpdateManage"
});
const route = useRoute();
async function submitUpdateUser() {
  userFormRef.value.validate(async (valid: any) => {
    if (valid) {
      if (File.value) {
        const data = getToken();
        await axios.post("/s3/tengxunyun/upload?directory=/admin", File.value, {
          headers: {
            "Content-Type": "multipart/form-data",
            "Authorization": formatToken(data.accessToken)
          }
        }).then(res => {
          formData.avatar = "https://test-1314985928.cos.ap-nanjing.myqcloud.com/" + res.data[0].key;
        }).catch(err => {
          ElMessage.error(err)
        })
      }
      const submitData = {
        id: 0,
        avatar: "",
        name: "",
        password: "",
        phone: "",
        userId: "",
      };
      for (const key in submitData) {
        if (formData[key]) {
          submitData[key] = formData[key];
        }
      }
      const md5: any = new Md5();
      md5.appendAsciiStr(submitData.password[0] + salt[5] + submitData.password.slice(1, submitData.password.length));
      submitData.password = md5.end();
      updateUser(submitData).then(res => {
        ElMessage.success("修改成功");
        close();
      }).catch(error => {
        ElMessage.error("修改失败");
      })
    }
  })
}

function beforeAvatarUpload(rawFile) {
  const isIMAGE = rawFile.type === 'image/jpeg' || 'image/jpg' || 'image/png';
  if (!isIMAGE) {
    message("上传文件只能是图片格式", { type: "error" });
    return false;
  }
  if (rawFile.size / 1024 / 1024 > 2) {
    message("上传文件大小不能超过2MB", { type: "error" });
    return false;
  }
  return true;
}
/**
 * 上传文件
 * @param file 
 */
function changeImage(file) {
  File.value = file;
  imageUrl.value = URL.createObjectURL(file.raw!)
}

/**
 * 关闭当前修改页并返回用户管理界面
 */
function close() {
  const valueIndex: number = multiTags.value.findIndex((item: any) => {
    return item.path === "/user/updateUser";
  })
  const spliceRoute = (startIndex: number, length: number): void => {
    useMultiTagsStoreHook().handleTags("splice", "", { startIndex, length });
  }
  spliceRoute(valueIndex, 1);
  // 跳转到上一页
  router.replace({
    name: "UserManage",
  });
}
onMounted(() => {
  getUserByUserId(route.query?.userId).then(res => {
    formData.name = res.data.data.data[0].name;
    formData.id = res.data.data.data[0].id;
    formData.phone = res.data.data.data[0].phone;
    formData.userId = res.data.data.data[0].user_id;
    formData.avatar = res.data.data.data[0].avatar;
    if (formData.avatar) {
      imageUrl.value = formData.avatar;
    }
  });
});
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>