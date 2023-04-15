<template>
    <div>
<!-- table -->
<el-table  :data="ecsList" style="width: 100%">
            <!-- <el-table-column prop="id" label="Id" width="180" /> -->
            <el-table-column prop="IpAddress" label="弹性ip"  />
            <el-table-column prop="ImageId" label="镜像ID" />
            <el-table-column prop="InstanceId" label="ecsID"  />
            <el-table-column prop="InstanceName" label="实例名"  />
            <el-table-column prop="InstanceType" label="实例规格" />
            <el-table-column prop="KeyPairName" label="密钥名"  />
            <el-table-column prop="OSName" label="操作系统"  />
            <el-table-column fixed="right" label="Operations" width="200">
                <template #default="scope">
                    <el-button link type="primary" size="small" @click="">Show</el-button>
                    <el-button link type="primary" size="small" @click="">Edit</el-button>
                    <el-button link type="primary" size="small" @click="">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted ,onBeforeMount} from 'vue'
import {AliyunEcs} from '../interface/index'
import {GetEcsList} from '../api/aliyunecsapi'

onMounted(()=>{
    getecslist()
})

const ecsList  = ref<AliyunEcs[]>()
const getecslist = async ()=>{
    const data = await GetEcsList()
    ecsList.value = data.data
}

</script>

<style scoped>

</style>