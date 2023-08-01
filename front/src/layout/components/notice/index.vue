<script setup lang="ts">
import { onMounted, onUnmounted, ref } from "vue";
import { noticesData } from "./data";
import NoticeList from "./noticeList.vue";
import Bell from "@iconify-icons/ep/bell";
import { useUserStore } from "@/store/modules/user";

const noticesNum = ref(0);
const notices = ref(noticesData);
const activeKey = ref(noticesData[0].key);
const SSE = ref(true);
// 从store中获取用户的信息
const store = useUserStore();

function reverseTime(updatedAt) {
  const date = new Date(updatedAt);
  // 使用 Date 对象的方法获取各个时间部分
  const year = date.getFullYear(); // 年份
  const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份
  const day = String(date.getDate()).padStart(2, '0'); // 日
  const hours = String(date.getHours()).padStart(2, '0'); // 小时
  const minutes = String(date.getMinutes()).padStart(2, '0'); // 分钟
  const seconds = String(date.getSeconds()).padStart(2, '0'); // 秒
  const milliseconds = String(date.getMilliseconds()).padStart(3, '0');
  // 拼接成字符串格式的时间
  const formattedTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}.${milliseconds}`;
  return formattedTime;
}
async function createEventSource(role: string) {
  const resp = await fetch(`/operations/System/Sub/SSE?roles=${role}&wg_live=true`);
  if (resp.ok) {
    const reader = resp.body.getReader();
    let accumulatedData = ''; // 用于累积数据
    while (SSE.value) {
      const { value, done } = await reader.read();
      if (done) {
        break;
      }
      accumulatedData += new TextDecoder().decode(value);
      // 检查是否完整接收了一条 SSE 事件数据
      if (accumulatedData.includes('\n\n')) {
        const [eventData, remainingData] = accumulatedData.split('\n\n');
        try {
          const currentInfo = JSON.parse(eventData).data.data;
          notices.value.forEach(item => {
            item.list.length = 0;
          })
          console.log(currentInfo);
          // 将currentInfo放入到消息通知里
          currentInfo.forEach(item => {
            if (item.type === "notice") {
              notices.value[0].list.push({
                avatar: "",
                title: `来自${item.create_role}的通知`,
                datetime: reverseTime(item.updatedAt),
                description: item.message,
                type: "1",
              });
            } else if (item.type === "message") {
              notices.value[1].list.push({
                avatar: "",
                title: `来自${item.create_role}的消息`,
                datetime: reverseTime(item.updatedAt),
                description: item.message,
                type: "2",
              });
            } else if (item.type === "todo") {
              notices.value[2].list.push({
                avatar: "",
                title: `${item.create_role}提醒您的待办`,
                datetime: reverseTime(item.updatedAt),
                description: item.message,
                type: "3",
              });
            }
          })
          noticesNum.value = currentInfo.length;
        } catch (error) {
          console.log(error);
        }
        accumulatedData = remainingData; // 处理完一条数据后，将剩余数据留作下一次拼接
      }
    }
  }
}

onMounted(() => {
  notices.value.forEach(item => {
    item.list.length = 0;
  })
  console.log(notices.value);
  store.roles.forEach(item => {
    createEventSource(item);
  })
});

onUnmounted(() => {
  SSE.value = false;
})

</script>

<template>
  <el-dropdown trigger="click" placement="bottom-end">
    <span class="dropdown-badge navbar-bg-hover select-none">
      <el-badge :value="noticesNum" :max="99">
        <span class="header-notice-icon">
          <IconifyIconOffline :icon="Bell" />
        </span>
      </el-badge>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-tabs :stretch="true" v-model="activeKey" class="dropdown-tabs"
          :style="{ width: notices.length === 0 ? '200px' : '330px' }">
          <el-empty v-if="notices.length === 0" description="暂无消息" :image-size="60" />
          <span v-else>
            <template v-for="item in notices" :key="item.key">
              <el-tab-pane :label="`${item.name}(${item.list.length})`" :name="`${item.key}`">
                <el-scrollbar max-height="330px">
                  <div class="noticeList-container">
                    <NoticeList :list="item.list" />
                  </div>
                </el-scrollbar>
              </el-tab-pane>
            </template>
          </span>
        </el-tabs>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<style lang="scss" scoped>
.dropdown-badge {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 48px;
  margin-right: 10px;
  cursor: pointer;

  .header-notice-icon {
    font-size: 18px;
  }
}

.dropdown-tabs {
  .noticeList-container {
    padding: 15px 24px 0;
  }

  :deep(.el-tabs__header) {
    margin: 0;
  }

  :deep(.el-tabs__nav-wrap)::after {
    height: 1px;
  }

  :deep(.el-tabs__nav-wrap) {
    padding: 0 36px;
  }
}
</style>
