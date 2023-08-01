<script setup lang="ts">
const url = "https://cos.ap-nanjing.myqcloud.com/test-1314985928/admin/WechatIMG25.png";
const srcList = [url];
defineOptions({
  name: "Welcome"
});
</script>

<template>
  <div class="app-container home">
    <el-row :gutter="20">
      <el-col :sm="24" :lg="12" style="padding-left: 20px">
        <h1>Fireboom Admin管理框架</h1>
        <div>
          <p style="text-indent:2em;">市面上有很多管理后台，为什么还要做 Fireboom Admin 呢？若不是其独一无二的特点，我们断然不会班门弄斧。</p>
          <p style="text-indent:2em;">Fireboom
            Admin最大的特点是可视化的后台接口开发，注意不是可视化拖拽界面。TA适用于网站管理后台、CMS、CRM、OA等中后台系统开发。当然，您也可以对TA深度定制。</p>
          <p style="text-indent:2em;">我将从两个方面说明，TA的与众不同:</p>
          <p style="text-indent:2em;">相对于传统后台管理框架：Fireboom Admin
            能够可视化构建后台接口，但前端部分仍采用传统的方式进行开发。这既保证了接口开发的快速，又保证了前端实现复杂交互的可行性。而且，Fireboom还内置了类似navicate的数据库管理界面，很多数据管理功能无需开发都本自具足。
          </p>
          <p style="text-indent:2em;">相对于拖拽构建界面的低代码：Fireboom Admin 具备可视化API开发的能力，且具有PRO
            CODE能力，能够实现任意复杂的业务逻辑，无论是后端逻辑还是前端逻辑，但代价是部分牺牲了构建页面的速度。Fireboom
            Admin还支持OIDC身份验证和RBAC角色鉴权，能够实现复杂的权限控制，包括数据权限和接口权限。
          </p>
          <p style="text-indent:2em;">未来， Fireboom Admin将考虑与appsmith或lowcode-engine结合，进一步提升后台开发效率！</p>
        </div>
        <p>
          <b>当前版本:</b> <span>v0.3</span>
        </p>
        <p>
          <el-tag type="danger">&yen;免费开源</el-tag>
        </p>
        <p>
          <el-button type="primary" href="https://github.com/Echoidf/fireboom-admin-web.git" target="_blank"
            rel="noopener noreferrer" tag="a">访问github</el-button>
          <el-button plain href="https://www.fireboom.io" tag="a" rel="noopener noreferrer"
            target="_blank">访问主页</el-button>
        </p>
      </el-col>

      <el-col :sm="24" :lg="12" style="padding-left: 50px">
        <el-row>
          <el-col :span="12">
            <h2>技术选型</h2>
          </el-col>
        </el-row>
        <el-row><br></el-row>
        <el-row>
          <el-col :span="6">
            <h4 style="font-size: medium;">后端技术</h4>
            <ul>
              <li>飞布可视化开发框架</li>
              <li>Golang</li>
              <li>JWT</li>
              <li>OIDC</li>
              <li>MySQL</li>
              <li>...</li>
            </ul>
          </el-col>
          <el-col :span="2"></el-col>
          <el-col :span="6">
            <h4 style="font-size: medium;">前端技术</h4>
            <ul>
              <li>Vue3</li>
              <li>TypeScript</li>
              <li>Pinia</li>
              <li>Element-Plus</li>
              <li>Axios</li>
              <li>Sass</li>
              <li>...</li>
            </ul>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-divider />
    <el-row :gutter="20">
      <el-col :xs="24" :sm="24" :md="12" :lg="8">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>联系信息</span>
            </div>
          </template>
          <div class="body">
            <p>
              <el-icon>
                <Promotion />
              </el-icon> &nbsp;官网:<el-link href="https://www.fireboom.io"
                target="_blank">https://www.fireboom.io</el-link>
            </p>
            <p>
            </p>
            <p>
              <el-icon>
                <User />
              </el-icon> &nbsp;微信: luxuncode
              <br>
            </p>
            <p>
              <el-image :src="url" :preview-src-list="srcList" :hide-on-click-modal="true" style="width:50%" />
            </p>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="24" :md="12" :lg="8">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>致谢</span>
            </div>
          </template>
          <div class="body">
            <p>
              <el-link href="https://yiming_chang.gitee.io/pure-admin-doc/" target="_blank">Pure-admin</el-link>
            </p>
            <p>
              <el-link href="http://doc.ruoyi.vip/" target="_blank">Ruoyi-admin</el-link>
            </p>
            <p>
              <el-link href="https://casdoor.org/zh/" target="_blank">Casdoor</el-link>
            </p>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.home {
  blockquote {
    padding: 10px 20px;
    margin: 0 0 20px;
    font-size: 17.5px;
    border-left: 5px solid #eee;
  }

  hr {
    margin-top: 20px;
    margin-bottom: 20px;
    border: 0;
    border-top: 1px solid #eee;
  }

  .col-item {
    margin-bottom: 20px;
  }

  ul {
    padding: 0;
    margin: 0;
  }

  font-family: "open sans",
  "Helvetica Neue",
  Helvetica,
  Arial,
  sans-serif;
  font-size: 13px;
  color: #676a6c;
  overflow-x: hidden;

  ul {
    list-style-type: none;
  }

  h4 {
    margin-top: 0px;
  }

  h2 {
    margin-top: 10px;
    font-size: 26px;
    font-weight: 100;
  }

  p {
    margin-top: 10px;

    b {
      font-weight: 700;
    }
  }

  .update-log {
    ol {
      display: block;
      list-style-type: decimal;
      margin-block-start: 1em;
      margin-block-end: 1em;
      margin-inline-start: 0;
      margin-inline-end: 0;
      padding-inline-start: 40px;
    }
  }
}

.card-header {
  font-size: large;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>