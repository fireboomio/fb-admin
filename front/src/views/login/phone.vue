<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { ref, reactive, computed } from "vue";
import { useRouter } from "vue-router";
import Motion from "./utils/motion";
import { message } from "@/utils/message";
import { phoneRules } from "./utils/rule";
import { ElMessage, type FormInstance } from "element-plus";
import { $t, transformI18n } from "@/plugins/i18n";
import { useVerifyCode } from "./utils/verifyCode";
import { useUserStoreHook } from "@/store/modules/user";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import Iphone from "@iconify-icons/ep/iphone";
import { sendVerifyCode } from "@/api/user";
import { usePermissionStoreHook } from "@/store/modules/permission";
import { addPathMatch } from "@/router/utils";

const { t } = useI18n();
const loading = ref(false);
const router = useRouter();
const ruleForm = reactive({
  phone: "18856264667",
  verifyCode: ""
});
const loginType = computed(() => {
  return useUserStoreHook().loginType;
});
const ruleFormRef = ref<FormInstance>();
const { isDisabled, text } = useVerifyCode();

const onLogin = async (formEl: FormInstance | undefined) => {
  loading.value = true;
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      // 模拟登录请求，需根据实际开发进行修改
      setTimeout(() => {
        useUserStoreHook()
          .loginByUsername({
            phone: ruleForm.phone,
            code: ruleForm.verifyCode,
            loginType: loginType.value
          })
          .then(res => {
            usePermissionStoreHook().handleWholeMenus([]);
            addPathMatch();
            if (res.success) {
              router.push("/");
              message("登录成功", { type: "success" });
            }
          })
          .catch(err => {
            loading.value = false;
            ElMessage.error(err);
          });
        message(transformI18n($t("login.loginSuccess")), { type: "success" });
        loading.value = false;
      }, 2000);
    } else {
      loading.value = false;
      return fields;
    }
  });
};

const sendCode = async () => {
  await sendVerifyCode({ dest: ruleForm.phone }).then(res => {
    console.log(res);
  });
};


function onBack() {
  useVerifyCode().end();
  useUserStoreHook().SET_LOGINTYPE("password");
}
</script>

<template>
  <el-form ref="ruleFormRef" :model="ruleForm" :rules="phoneRules" size="large">
    <Motion>
      <el-form-item prop="phone">
        <el-input clearable v-model="ruleForm.phone" :placeholder="t('login.phone')"
          :prefix-icon="useRenderIcon(Iphone)" />
      </el-form-item>
    </Motion>

    <Motion :delay="100">
      <el-form-item prop="verifyCode">
        <div class="w-full flex justify-between">
          <el-input clearable v-model="ruleForm.verifyCode" :placeholder="t('login.smsVerifyCode')"
            :prefix-icon="useRenderIcon('ri:shield-keyhole-line')" />
          <el-button :disabled="isDisabled" class="ml-2" @click="sendCode">
            {{
              text.length > 0
              ? text + t("login.info")
              : t("login.getVerifyCode")
            }}
          </el-button>
        </div>
      </el-form-item>
    </Motion>

    <Motion :delay="150">
      <el-form-item>
        <el-button class="w-full" size="default" type="primary" :loading="loading" @click="onLogin(ruleFormRef)">
          {{ t("login.login") }}
        </el-button>
      </el-form-item>
    </Motion>

    <Motion :delay="200">
      <el-form-item>
        <el-button class="w-full" size="default" @click="onBack">
          {{ t("login.back") }}
        </el-button>
      </el-form-item>
    </Motion>
  </el-form>
</template>
