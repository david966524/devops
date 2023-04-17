<template>
    <div class="domain-checker">
      <el-input v-model="input" placeholder="输入你要检测的域名" class="input-with-select" style="width: 500px">
        <template #prepend>
          <el-select v-model="select" placeholder="Select" style="width: 100px">
            <el-option label="http://" value="http://" />
            <el-option label="https://" value="https://" />
          </el-select>
        </template>
        <template #append>
          <el-button :icon="Search" @click="checkdomain()" />
        </template>
      </el-input>
      <el-divider>
      <el-icon><star-filled /></el-icon>
    </el-divider>
      <div class="result-container">
        <el-form label-position="left" label-width="100px" :model="message" style="max-width: 500px; border-top: 2px dashed var(--el-border-color);">
          <el-form-item label="URL" class="result-item">
            {{ message.url }}
          </el-form-item>
          <el-form-item label="QQ检测" class="result-item">
            {{ message.qq_msg }}
          </el-form-item>
          <el-form-item label="VX检测" class="result-item">
            {{ message.vx_msg }}
          </el-form-item>
          <el-form-item label="原因" class="result-item">
            {{ message.cause }}
          </el-form-item>
          <el-form-item label="ICP名字" class="result-item">
            {{ message.icp_name }}
          </el-form-item>
          <el-form-item label="ICP" class="result-item">
            {{ message.icp }}
          </el-form-item>
        </el-form>
      </div>
    </div>
  </template>

<script setup lang="ts">
import { ref } from 'vue'
import { Delete, Edit, Search, Share, Upload } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { Ckeckdomain } from '../api/checkdomain'
import { checkdomainResult } from '../interface/index'
const input = ref<string>('')
const select = ref<string>('https://')
const message = ref<checkdomainResult>({})
const checkdomain = async () => {
    console.log(select, input)
    if (input.value == "") {
        ElMessage.error('域名不能为空')
    }
    const data = await Ckeckdomain(input.value)
    message.value = <checkdomainResult>data.data
}
</script>

<style scoped>
.domain-checker {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 20px;
}



.result-container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

.result-item {
  margin-bottom: 16px;
}
</style>