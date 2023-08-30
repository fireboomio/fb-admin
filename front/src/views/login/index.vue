<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { message } from "@/utils/message";
import { loginRules } from "./utils/rule";
import { useNav } from "@/layout/hooks/useNav";
import { ElMessage, type FormInstance } from "element-plus";
import { $t, transformI18n } from "@/plugins/i18n";
import { useLayout } from "@/layout/hooks/useLayout";
import { useUserStoreHook } from "@/store/modules/user";
import { addPathMatch } from "@/router/utils";
import { bg, avatar, illustration } from "./utils/static";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { ref, reactive, toRaw, onMounted, onBeforeUnmount, computed } from "vue";
import { useTranslationLang } from "@/layout/hooks/useTranslationLang";
import { useDataThemeChange } from "@/layout/hooks/useDataThemeChange";
import Motion from "./utils/motion";
import phone from "./phone.vue";
import dayIcon from "@/assets/svg/day.svg?component";
import darkIcon from "@/assets/svg/dark.svg?component";
import globalization from "@/assets/svg/globalization.svg?component";
import Lock from "@iconify-icons/ri/lock-fill";
import Check from "@iconify-icons/ep/check";
import User from "@iconify-icons/ri/user-3-fill";
import { usePermissionStoreHook } from "@/store/modules/permission";
import { setInterval } from "timers";
import { removeToken } from "@/utils/auth";
import { Md5 } from "ts-md5"
const salt = "qwertyuiop";
defineOptions({
  name: "Login"
});
const router = useRouter();
const loading = ref(false);
const ruleFormRef = ref<FormInstance>();
const { initStorage } = useLayout();
initStorage();
const { t } = useI18n();
const { dataTheme, dataThemeChange } = useDataThemeChange();
dataThemeChange();
const { title, getDropdownItemStyle, getDropdownItemClass } = useNav();
const { locale, translationCh, translationEn } = useTranslationLang();

const loginType = computed(() => {
  return useUserStoreHook().loginType;
});

const ruleForm = reactive({
  username: "zql",
  password: "abc1234"
});


const onLogin = async (formEl: FormInstance | undefined) => {
  loading.value = true;
  // 对用户登录的密码进行MD5加密
  const md5: any = new Md5();
  md5.appendAsciiStr(ruleForm.password[0] + salt[5] + ruleForm.password.slice(1, ruleForm.password.length));
  const password = md5.end();
  // 清除token
  removeToken();
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      useUserStoreHook()
        .loginByUsername({
          username: ruleForm.username,
          password: password,
          loginType: loginType.value
        })
        .then(res => {
          // 全部采取静态路由模式
          usePermissionStoreHook().handleWholeMenus([]);
          addPathMatch();
          if (res.success) {
            message("登录成功", { type: "success" });
            router.push("/");
          }
        })
        .catch(err => {
          loading.value = false;
          ElMessage.error(err);
        });
    } else {
      loading.value = false;
      return fields;
    }
  });
};

/** 使用公共函数，避免`removeEventListener`失效 */
function onkeypress({ code }: KeyboardEvent) {
  if (code === "Enter") {
    onLogin(ruleFormRef.value);
  }
}

onMounted(() => {
  window.document.addEventListener("keypress", onkeypress);
});

onBeforeUnmount(() => {
  window.document.removeEventListener("keypress", onkeypress);
});
</script>

<template>
  <div class="select-none">
    <!-- <img :src="bg" class="wave" /> -->
    <div class="flex-c absolute right-5 top-3">
      <!-- 主题 -->
      <!-- <el-switch v-model="dataTheme" inline-prompt :active-icon="dayIcon" :inactive-icon="darkIcon"
        @change="dataThemeChange" /> -->
      <!-- 国际化 -->
      <el-dropdown trigger="click">
        <globalization
          class="hover:text-primary hover:!bg-[transparent] w-[20px] h-[20px] ml-1.5 cursor-pointer outline-none duration-300" />
        <template #dropdown>
          <el-dropdown-menu class="translation">
            <el-dropdown-item :style="getDropdownItemStyle(locale, 'zh')"
              :class="['dark:!tsext-white', getDropdownItemClass(locale, 'zh')]" @click="translationCh">
              <IconifyIconOffline class="check-zh" v-show="locale === 'zh'" :icon="Check" />
              简体中文
            </el-dropdown-item>
            <el-dropdown-item :style="getDropdownItemStyle(locale, 'en')"
              :class="['dark:!text-white', getDropdownItemClass(locale, 'en')]" @click="translationEn">
              <span class="check-en" v-show="locale === 'en'">
                <IconifyIconOffline :icon="Check" />
              </span>
              English
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <div class="lefttop">
      <a href="https://www.fireboom.io/" target=“_blank”>
        <img src="/src/assets/login/Logo.png" style="width:50%;" />
      </a>
    </div>
    <div class="login-container">
      <div class="img">
        <div class="title">
          <span class="highLight">前端</span>变全栈, <span class="highLight">后端</span>不搬砖
          <hr class="style-one">
          <ul class="table">
            <li><el-icon class="checkIcon"><Select /></el-icon>&nbsp;支持连接自定义数据库和三方数据源</li>
            <li><el-icon class="checkIcon"><Select /></el-icon>&nbsp;多语言兼容, 支持主流编程语言</li>
            <li><el-icon class="checkIcon"><Select /></el-icon>&nbsp;可视化操作面板, 简单易学</li>
          </ul>
          <div class="cards">
            <div class="card">
              <img src="/src/assets/login/database.png" class="cardsImg">
              <div class="cardTitle">
                数据库建模
              </div>
              <div class="cardContent">
                内置数据库建模功能，技术小白也能驾驭。
              </div>
            </div>
            <div class="card">
              <img src="/src/assets/login/Github.png" class="cardsImg">
              <div class="cardTitle">
                GitHub集成
              </div>
              <div class="cardContent">
                不断地将生成的应用程序推送到你的GitHub存储库，实现版本管理。
              </div>
            </div>
            <div class="card">
              <img src="/src/assets/login/SDK.png" class="cardsImg">
              <div class="cardTitle">
                SDK生成
              </div>
              <div class="cardContent">
                生成带有自动登录功能的客户端SDK方便前后端对接接口。
              </div>
            </div>
            <div class="card">
              <img src="/src/assets/login/VSCODE.png" class="cardsImg">
              <div class="cardTitle">
                VSCODE
              </div>
              <div class="cardContent">
                提供vscode插件，方便高端玩家使用更深功能。
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="login-box">

        <div class="login-form">
          <!-- <avatar class="avatar" /> -->
          <div class="loginTitle">登录你的飞布账号</div>
          <div class="signUp">还没有账号? <el-button type="danger" link><span class="signUpNow">立即注册</span> <el-icon style="border-radius: 50%;  background-color: #FE4D60; font-size:10px;margin-top: -3px;margin-left: 2px;
  color: white;">
                <ArrowRight />
              </el-icon>
            </el-button>
          </div>
          <!-- <Motion>
            <h2 class="outline-none">{{ title }}</h2>
          </Motion> -->
          <el-form ref="ruleFormRef" :model="ruleForm" :rules="loginRules" size="large" v-if="loginType === 'password'">
            <Motion :delay="100">
              <el-form-item :rules="[
                {
                  required: true,
                  message: transformI18n($t('login.usernameReg')),
                  trigger: 'blur'
                }
              ]" prop="username">
                <el-input clearable v-model="ruleForm.username" :placeholder="t('login.username')"
                  :prefix-icon="useRenderIcon(User)" />
              </el-form-item>
            </Motion>

            <Motion :delay="150">
              <el-form-item prop="password">
                <el-input clearable v-model="ruleForm.password" :placeholder="t('login.password')" show-password
                  :prefix-icon="useRenderIcon(Lock)" />
              </el-form-item>
            </Motion>
            <Motion :delay="150">
              <el-checkbox size="large"></el-checkbox> <span style="
                position: absolute;
                top: 11px;
                left: 19px;
                font-family: PingFangSC-Regular;
                font-size: 12px;
                color: rgba(95,98,105,0.60);
                margin-bottom: 10px;
                ">记住账号
              </span>
            </Motion>
            <Motion :delay="250">
              <el-button class="w-full mt-4" size="default" type="primary" :loading="loading"
                @click="onLogin(ruleFormRef)">
                {{ t("login.login") }}
              </el-button>
            </Motion>
            <Motion :delay="300">
              <div class="w-full h-[20px] flex justify-between items-center mt-4">
                <el-button class="w-full mt-4" size="default" @click="useUserStoreHook().SET_LOGINTYPE('sms')">
                  {{ t("login.loginBySms") }}
                </el-button>
              </div>
            </Motion>
          </el-form>


          <!-- 手机号登录 -->
          <phone v-if="loginType === 'sms'" />
          <div class="xieyi">点击登录表示你同意 <span class="fy">服务协议</span> 和 <span class="fy">隐私协议</span></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url("@/style/login.css");
</style>

<style lang="scss" scoped>
:deep(.el-input-group__append, .el-input-group__prepend) {
  padding: 0;
}

.translation {
  ::v-deep(.el-dropdown-menu__item) {
    padding: 5px 40px;
  }

  .check-zh {
    position: absolute;
    left: 20px;
  }

  .check-en {
    position: absolute;
    left: 20px;
  }
}

.left {
  width: 100%;
  height: 900px;
  background-image: linear-gradient(133deg, #FFEEEE 0%, #FFB9B9 100%);
}

.lefttop {
  position: absolute;
  top: 20px;
  left: 30px;

}

.highLight {
  // width: 62px;
  // height: 44px;
  font-family: PingFangSC-Semibold;
  font-size: 28px;
  color: #FF4D60;
  letter-spacing: 2.59px;
  line-height: 44px;
  font-weight: 600;
}

.title {
  font-family: PingFangSC-Semibold;
  font-size: 28px;
  color: #1F2328;
  letter-spacing: 2.59px;
  line-height: 44px;
  font-weight: 600;
}

.checkIcon {
  position: relative;
  top: 4px;
  color: blue;
}

hr.style-one {
  margin-top: 25px;
  border: 0;
  height: 1px;
  width: 331px;
  background: #333;
  background-image: linear-gradient(to left, #EE5F6B 40%, #F18BC7 100%);
}

.table {
  margin-top: 20px;
  font-family: PingFangSC-Regular;
  font-size: 14px;
  color: rgba(40, 31, 31, 0.80);
  letter-spacing: 1.44px;
  line-height: 35px;
  font-weight: 400;
}


.cards {
  margin-top: 48px;
  display: grid;
  grid-template-columns: 158px 158px;
  grid-template-rows: 195px 195px;
  gap: 14px;
}

.card {
  font-size: 24px;
  background-image: linear-gradient(136deg, rgba(255, 242, 242, 0.63) 0%, #FECED3 100%);
  box-shadow: 0px 2px 7px 0px rgba(207, 169, 169, 0.48);
  border-radius: 4px;
}

.card:hover {
  transform: scale(1.1);
  /* 当鼠标悬停时，将卡片放大 10% */
  box-shadow: 0px 5px 10px 0px rgba(207, 169, 169, 0.8);
  /* 修改阴影效果 */
}

.cardsImg {
  height: 24px;
  width: 24px !important;
  position: inherit;
  margin-top: 18px;
  margin-left: 16px;

}

.cardTitle {
  width: 120px;
  font-family: PingFangSC-Semibold;
  font-size: 18px;
  color: #FF4D60;
  letter-spacing: 1.66px;
  position: relative;
  font-weight: 600;
  margin-top: 16x;
  margin-left: 16px;
}

.cardContent {
  display: inline-block;
  width: 126px;
  opacity: 0.8;
  font-family: PingFangSC-Regular;
  font-size: 14px;
  color: rgba(40, 31, 31, 0.80);
  letter-spacing: 1.44px;
  line-height: 21px;
  font-weight: 400;
  margin-left: 16px;
}

.loginTitle {
  // margin-top: 253px;
  font-family: PingFangSC-Semibold;
  font-size: 32px;
  color: #1F2328;
  letter-spacing: 3.3px;
  line-height: 47px;
  font-weight: 600;
  position: relative;
}

.signUp {
  height: 47px;
  font-family: PingFangSC-Regular;
  font-size: 12px;
  color: rgba(40, 31, 31, 0.80);
  letter-spacing: 1.24px;
  line-height: 47px;
  font-weight: 400;
}

.signUpNow {
  color: #FF4D60;
  height: 47px;
  font-family: PingFangSC-Regular;
  font-size: 12px;
  letter-spacing: 1.24px;
  line-height: 44px;
  font-weight: 400;
}

.signUpNow:hover {
  text-decoration: underline;
  color: red;
}

@media screen and (max-width: 1450px) {
  .lefttop {
    display: none;
  }
}

.xieyi {
  text-align: center;
  margin-top: 25px;
  font-family: PingFangSC-Regular;
  font-size: 14px;
  color: rgba(95, 98, 105, 0.60);
  letter-spacing: 0;
  line-height: 20px;
  font-weight: 400;
}

.fy {
  color: #E92E5E;
}

.fy:hover {
  color: red;
}
</style>
