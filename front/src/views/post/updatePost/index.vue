<template>
  <div class="app-container">
    <el-card>
      <template #header>
        编辑文章信息
      </template>
      <el-form ref="dataFormRef" :model="formData" :rules="rules" label-width="150px" v-loading="loading">
        <el-form-item label="标题" prop="title">
          <el-input v-model="formData.title" placeholder="请输入文章标题" style="width: 200px" />
        </el-form-item>
        <el-form-item label="文章类别" prop="Category">
          <el-dropdown trigger="click" @command="searchPostByCategory">
            <el-button type="primary">
              {{ currentPostCategory }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item v-for="cate in category" :command="cate">{{
                  cate.name }}</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-form-item>
        <!-- 后端接收文件的URL -->
        <el-form-item label="封面" prop="poster">
          <el-upload class="avatar-uploader" action="#" :show-file-list="false" accept="image/jpg,image/jpeg,image/png"
            :before-upload="beforeAvatarUpload" :on-change="changeImage" :auto-upload="false">
            <img v-if="imageUrl" :src="imageUrl" class="avatar" title="点击重新上传" />
            <el-icon v-else class="avatar-uploader-icon">
              <plus />
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input type="textarea" :rows="5" v-model="formData.content" placeholder="请输入文章内容" />
        </el-form-item>
        <el-form-item label="发布时间" prop="publishedAt">
          <el-date-picker v-model="formData.publishedAt" type="date" placeholder="请选择发布时间" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmitUpdate">修 改</el-button>
          <el-button @click="close">取 消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import api from "@/api";
import { message } from '@/utils/message';
import { Post__CreateOneInput } from "@/api/models";
import type { Arrayable } from "element-plus/es/utils";
import router from "@/router";
import { getCategory, updatePost } from "@/api/post";
import { FormItemRule, ElForm, ElMessage, ElMessageBox, ElPagination, ElUpload, ElIcon } from "element-plus";
import { formatToken, getToken } from '@/utils/auth';
import axios from 'axios';
import { useTags } from "@/layout/hooks/useTag";
import { useMultiTagsStoreHook } from "@/store/modules/multiTags";


const rules = reactive<
  Partial<Record<keyof Post__CreateOneInput, Arrayable<FormItemRule>>>
>({
  title: [{ required: true, message: "请输入文章标题", trigger: "blur" }],
  Category: [{ required: true, message: "选择文章类别", trigger: "blur" }],
  poster: [{ required: false, message: "请上传文章封面", trigger: "blur" }],
  content: [{ required: true, message: "请输入文章内容", trigger: "blur" }]
});
const loading = ref(false);
const currentPostCategory = ref("文章类别"); // 文章类别初始值
const route = useRoute();
const imageUrl = ref('');
const File = ref();
const dataFormRef = ref(ElForm);
const category = ref([]);  //存储所有文章类别
const { multiTags } = useTags();
const formData = reactive({
  cate: "",
  id: "",
  title: "",
  poster: "",
  publishedAt: "",
  author: "",
  content: "",
  Category: "",
  userId: "",
})

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

async function handleQuery() {
  const { error, data } = await api.query({
    operationName: "Post/GetOne",
    input: { id: Number(route.query?.id) }
  });
  if (!error) {
    formData.id = String(data.data.id);
    formData.title = data.data.title;
    formData.publishedAt = data.data.published_at;
    formData.author = data.data.author;
    formData.Category = data.data.case_category.name;
    formData.cate = data.data.case_category.id;
    formData.content = data.data.content;
    formData.poster = data.data.poster;
    currentPostCategory.value = formData.Category;
    formData.author = data.data.admin_user.name;
    formData.userId = data.data.admin_user.id;
    if (data.data.poster) {
      imageUrl.value = data.data.poster;
    }
  }
  // 获取所有文章种类
  await getCategory().then(res => {
    category.value = res.data.data.data.map(item => {
      return {
        id: item.id,
        name: item.name,
      }
    })
    category.value = category.value.filter(item => item.name)
  })
}
/**
 * 下拉菜单选择回调函数
 */
function searchPostByCategory(cate) {
  currentPostCategory.value = cate.name;
  formData.cate = cate.id;
  formData.Category = cate.name;
}
/**
 * 提交修改文章
 */
async function onSubmitUpdate() {
  loading.value = true;
  const data = getToken();
  if (File.value) {
    await axios.post("/s3/tengxunyun/upload?directory=/admin", File.value, {
      headers: {
        "Content-Type": "multipart/form-data",
        "Authorization": formatToken(data.accessToken)
      }
    }).then(res => {
      formData.poster = "https://test-1314985928.cos.ap-nanjing.myqcloud.com/" + res.data[0].key;
    }).catch(err => {
      ElMessage.error(err)
    })
  }
  dataFormRef.value.validate(async (isValid: boolean) => {
    if (isValid) {
      const submitData = {
        id: 0,
        title: "",
        poster: "",
        content: "",
        author: "",
        cate: "",
        userId: ""
      };
      for (const key in submitData) {
        if (formData[key]) {
          submitData[key] = formData[key];
        }
      }
      submitData.id = Number(submitData.id);
      await updatePost(submitData).then(res => {
        ElMessage.success("修改成功");
        close();
      }).catch(error => {
        ElMessage.error("修改失败");
      })
    }
  })
  loading.value = false;
}
onMounted(() => {
  handleQuery();
})


/**
 * 撤退到文章管理页面
 */
function close() {
  // 关闭当前标签页
  const valueIndex: number = multiTags.value.findIndex((item: any) => {
    return item.path === "/post/updatePost";
  })
  const spliceRoute = (startIndex: number, length: number): void => {
    useMultiTagsStoreHook().handleTags("splice", "", { startIndex, length });
  }
  spliceRoute(valueIndex, 1);
  // 跳转到上一页
  router.replace({
    name: "PostManage",
  });
}

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