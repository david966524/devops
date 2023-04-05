<template>
  <el-container class="layout-container-demo" style="height: 100vh">
    <el-aside width="200px">
      <el-scrollbar>
        <el-menu :default-openeds="['1', '3']" router>
          <el-sub-menu index="1">
            <template #title>
              <el-icon>
                <icon-menu />
              </el-icon>域名管理
            </template>
            <el-menu-item index="/cloudflare">Couldflare</el-menu-item>
            <el-menu-item index="">方能</el-menu-item>
            <el-menu-item index="/asiacloud">亚洲云海</el-menu-item>
            <el-sub-menu index="1-1">
              <template #title>包网项目</template>
              <el-menu-item index="/airborne">空降平台信息</el-menu-item>
              <el-menu-item index="/im">IM</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
           <el-sub-menu index="4">
            <template #title>
              <el-icon>
                <icon-menu />
              </el-icon>服务器
            </template>
            <el-sub-menu index="4-1">
              <template #title>aliyun</template>
              <el-menu-item index="">服务器列表</el-menu-item>
              <el-menu-item index="">list</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
          <el-sub-menu index="2">
            <template #title>
              <el-icon><ChromeFilled /></el-icon>工具
            </template>
            <el-menu-item-group>
              <template #title>DNS Tools</template>
              <el-menu-item index="/boce">拨测</el-menu-item>
              <el-menu-item index="/itdog">ITDog</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group 2">
              <el-menu-item index="2-3">Option 3</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="2-4">
              <template #title>Option 4</template>
              <el-menu-item index="2-4-1">Option 4-1</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title>
              <el-icon>
                <setting />
              </el-icon>System
            </template>
            <el-menu-item-group>
              <template #title>用户管理</template>
              <el-menu-item index="/user">用户列表</el-menu-item>
              <el-menu-item index="3-2">Option 2</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
        </el-menu>
      </el-scrollbar>
    </el-aside>

    <el-container>
      <el-header style="text-align: right; font-size: 12px">
        <div class="toolbar">
          <el-dropdown>
            <el-icon style="margin-right: 8px; margin-top: 1px">
              <setting />
            </el-icon>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>View</el-dropdown-item>
                <el-dropdown-item>Add</el-dropdown-item>
                <el-dropdown-item>Delete</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <span>{{ username }}</span>
        </div>
      </el-header>
      <!-- main  -->
      <el-main>
        <router-view></router-view>
      </el-main>

    </el-container>
  </el-container>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'





const username = ref(localStorage.getItem('username'))

</script>
  
<style scoped>
.layout-container-demo .el-header {
  position: relative;
  background-color: var(--el-color-primary-light-7);
  color: var(--el-text-color-primary);
}

.layout-container-demo .el-aside {
  color: var(--el-text-color-primary);
  background: var(--el-color-primary-light-8);
}

.layout-container-demo .el-menu {
  border-right: none;
}

.layout-container-demo .el-main {
  padding: 0;
}

.layout-container-demo .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}
</style>
  