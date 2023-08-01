<script lang="ts">
export default {
  name: "post",
  components: { WangEditor }
};
</script>

<script setup lang="ts">
import api, { convertPageQuery } from "@/api";
import { Post__CreateOneInput, Post__GetListResponse } from "@/api/models";
import { merge } from "@/utils";
import WangEditor from "@/components/WangEditor/index.vue";
import { FormItemRule, ElForm, ElMessage, ElMessageBox, ElPagination, ElUpload, ElIcon } from "element-plus";
import type { Arrayable } from "element-plus/es/utils";
import { ref, reactive, onMounted } from "vue";
import { Icon } from '@iconify/vue';
import { createOne } from "@/api/post";
import { useUserStoreHook } from "@/store/modules/user";
import { message } from '@/utils/message';
import axios from 'axios';
import { formatToken, getToken } from '@/utils/auth';
const queryFormRef = ref(ElForm);
const dataFormRef = ref(ElForm);
const hackReset = ref(false)
const loading = ref(false);
const ids = ref<number[]>([]);
const total = ref(0);
const imageUrl = ref('');
const fileList = ref([]);
const published_at = ref("");
const queryParams = reactive<PageQuery>({
  title: "",
  pageNum: 1,
  pageSize: 10
});

const dataSource = ref<
  NonNullable<NonNullable<Post__GetListResponse["data"]>["data"]>
>([]);

const dialog = reactive<DialogOption>({
  visible: false
});

const formData = reactive({
  title: "",
  poster: "",
  content: "",
  publishedAt: "",
});
const editingId = ref<number>();

const rules = reactive<
  Partial<Record<keyof Post__CreateOneInput, Arrayable<FormItemRule>>>
>({
  title: [{ required: true, message: "请输入文章标题", trigger: "blur" }],
  poster: [{ required: false, message: "请上传文章封面", trigger: "blur" }],
  content: [{ required: true, message: "请输入文章内容", trigger: "blur" }]
});

/**
 * 查询
 */
async function handleQuery() {
  loading.value = true;
  const { error, data } = await api.query({
    operationName: "Post/GetList",
    input: convertPageQuery(queryParams, { containsFields: ["title"] })
  });
  if (!error) {
    dataSource.value = data!.data!;
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
 * 行checkbox change事件
 */
function handleSelectionChange(selection: any) {
  ids.value = selection.map((item: any) => item.id);
}

/**
*  时间格式转换
*/
function reverse(dateStr: string) {
  // 创建一个新的Date对象
  const date = new Date(dateStr);
  date.setDate(date.getDate() + 1);
  // 定义月份的简写名称数组
  const monthNames = [
    "Jan", "Feb", "Mar", "Apr", "May", "Jun",
    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
  ];

  // 获取日期中的月份、日期、年份以及时区偏移
  const month = monthNames[date.getMonth()];
  const day = date.getDate().toString().padStart(2, '0');
  const year = date.getFullYear();
  const time = date.toTimeString().split(' ')[0];
  // const timezoneOffset = date.toString().match(/\((.*)\)/)[1];

  // 构建最终格式的字符串
  const formattedDate = `Mon ${month} ${day} ${year} ${time} GMT+0800 (GMT+08:00)`;
  return formattedDate;
}
/**
 * 新增/编辑
 *
 * @param id 数据ID
 */
async function openDialog(post) {
  if (post) {
    editingId.value = post.id;
    if (editingId.value) {
      dialog.title = "修改文章";
      imageUrl.value = post.poster;
      const { error, data } = await api.query({
        operationName: "Post/GetOne",
        input: { id: editingId.value }
      });
      if (!error) {
        merge(formData, data!.data!);
        formData.publishedAt = data!.data!.published_at;
        published_at.value = data!.data!.published_at;
        if (formData.publishedAt) {
          formData.publishedAt = reverse(formData.publishedAt);
        }
        hackReset.value = true;
        dialog.visible = true;
      } else {
        ElMessage.error("获取数据失败");
      }
    }
  } else {
    hackReset.value = true;
    dialog.visible = true;
    dialog.title = "新增文章";
  }

}

/**
 * 表单提交
 */
async function handleSubmit() {
  loading.value = false;
  const data = getToken();
  if (File) {
    await axios.post("/s3/tengxunyun/upload?directory=/admin", File, {
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
      if (editingId.value) {
        formData.publishedAt = published_at.value;
        const { error } = await api.mutate({
          operationName: "Post/UpdateOne",
          input: {
            id: editingId.value,
            ...formData
          }
        });
        if (!error) {
          ElMessage.success("修改成功");
          closeDialog();
          handleQuery();
        } else {
          ElMessage.error("修改失败");
        }
        loading.value = false;
      } else {
        const username = useUserStoreHook().username;
        createOne(formData.title, username, formData.content, formData.poster, formData.publishedAt).then(res => {
          if (res.data.data.data) {
            ElMessage.success("新增成功");
            closeDialog();
            handleQuery();
          } else {
            ElMessage.error("新增失败");
          }
          loading.value = false;
        });
      }
    }
  });
}

/**
 * 关闭弹窗
 */
function closeDialog() {
  hackReset.value = false;
  dialog.visible = false;
  resetForm();
}

/**
 * 重置表单
 */
function resetForm() {
  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();

  formData.title = "";
  formData.poster = "";
  formData.content = "";
  formData.publishedAt = "";
}

/**
 * 删除数据
 */
function handleDelete(id?: number) {
  const idList = id ? [id] : ids.value;
  if (!idList.length) {
    ElMessage.warning("请勾选删除项");
    return;
  }

  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const { error } = await api.mutate({
      operationName: "Post/DeleteMany",
      input: { ids: idList }
    });
    if (!error) {
      ElMessage.success("删除成功");
      resetQuery();
    }
  });
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
let File = undefined;
function handleAvatarSuccess(res, file) {
  imageUrl.value = URL.createObjectURL(file.raw!)
}
function changeImage(file) {
  File = file;
  imageUrl.value = URL.createObjectURL(file.raw!)
}
function UploadImage(param) {
}

onMounted(() => {
  handleQuery();
});

</script>

<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="标题" prop="title">
          <el-input v-model="queryParams.title" placeholder="标题" clearable @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item>
          <Auth value="post:query">
            <el-button type="primary" @click="handleQuery()">
              <Icon icon="ep:search" />搜索
            </el-button>
          </Auth>
          <el-button @click="resetQuery()">
            <Icon icon="ep:refresh" />重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <template #header>
        <Auth value="post:edit">
          <el-button type="success" @click="openDialog()">
            <Icon icon="ep:plus" />新增
          </el-button>
        </Auth>

        <Auth value="post:remove">
          <el-button type="danger" :disabled="ids.length === 0" @click="handleDelete()">
            <Icon icon="ep:delete" />删除
          </el-button>
        </Auth>

      </template>
      <el-table highlight-current-row :data="dataSource" v-loading="loading" @selection-change="handleSelectionChange"
        border>
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="标题" prop="title" width="200" ellipsis />
        <el-table-column label="封面" prop="poster" width="200">
          <template #default="scope">
            <el-image v-if="scope.row.poster.startsWith('https')" :src="scope.row.poster" fit="fill"
              :preview-src-list="[scope.row.poster]" :preview-teleported="true" :hide-on-click-modal="true" />
            <el-image v-else src="https://pic4.zhimg.com/80/v2-15a8561ac8de1d9c8a8bbbb502d4e873_720w.webp"
              :preview-src-list="['https://pic4.zhimg.com/80/v2-15a8561ac8de1d9c8a8bbbb502d4e873_720w.webp']" fit="fill"
              :preview-teleported="true" :hide-on-click-modal="true" />
          </template>
        </el-table-column>
        <el-table-column label="作者" prop="author" width="200" ellipsis align="center" />
        <el-table-column label="发布时间" prop="published_at" align="center" :formatter="(row, col, v) => (v ? new Date(v).toLocaleDateString() : '')
          " />
        <el-table-column fixed="right" label="操作" align="center" width="220">
          <template #default="scope">
            <Auth value="post:edit">
              <el-button type="primary" link size="small" @click.stop="openDialog(scope.row)">
                <Icon icon="ep:edit" />编辑
              </el-button>
            </Auth>
            <Auth value="post:remove">
              <el-button type="primary" link size="small" @click.stop="handleDelete(scope.row.id)">
                <Icon icon="ep:delete" />删除
              </el-button>
            </Auth>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination v-if="total > 0" v-model:total="total" v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize" @pagination="handleQuery" />
    </el-card>

    <el-dialog :title="dialog.title" v-model="dialog.visible" width="900px" @close="closeDialog">
      <el-form ref="dataFormRef" :model="formData" :rules="rules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input v-model="formData.title" placeholder="请输入文章标题" />
        </el-form-item>
        <!-- 后端接收文件的URL -->
        <el-form-item label="封面" prop="poster">
          <el-upload class="avatar-uploader" action="#" :show-file-list="false" accept="image/jpg,image/jpeg,image/png"
            :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload" :on-change="changeImage"
            :http-request="UploadImage" :file-list="fileList">
            <img v-if="imageUrl" :src="imageUrl" class="avatar" title="点击重新上传" />
            <!-- <i class="el-icon-plus avatar-uploader-icon"></i> -->
            <el-icon v-else class="avatar-uploader-icon">
              <plus />
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <WangEditor v-if="hackReset" v-model="formData.content" />
          <!-- <el-input type="textarea" :rows="5" v-model="formData.content" placeholder="请输入文章内容" /> -->
        </el-form-item>
        <el-form-item label="发布时间" prop="publishedAt">
          <el-date-picker v-model="formData.publishedAt" type="date" placeholder="请选择发布时间" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <Auth value="post:edit">
            <el-button type="primary" @click="handleSubmit">确 定</el-button>
          </Auth>
          <el-button @click="closeDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

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