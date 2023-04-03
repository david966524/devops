<template>
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane v-for="(tab, index) in tabs" :key="index" :label="tab.name" :name="tab.name">{{ tab.name
        }}</el-tab-pane>
        <el-table :data="asiacloudtable" style="width: 100%">
            <el-table-column prop="domain" label="domain"  />
            <el-table-column prop="host" label="host"  />
            <el-table-column prop="vhost" label="vhost"  />
            <el-table-column prop="status" label="status"  />
            <el-table-column prop="cname" label="cname"  />
            <el-table-column fixed="right" label="Operations" width="120">
                <template #default>
                    <el-button link type="primary" size="small" @click="">Detail</el-button>
                    <el-button link type="primary" size="small">Edit</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-tabs>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { TabsPaneContext } from 'element-plus'
import { Asiacloudvhost, AsiacloudDomain } from '../interface/index'
import { VhostList, GetDomain } from '../api/asiacloudapi'

// const tabs = [
//   { label: 'User', name: 'first', content: 'User' },
//   { label: 'Config', name: 'second', content: 'Config' },
//   { label: 'Role', name: 'third', content: 'Role' },
//   { label: 'Task', name: 'fourth', content: 'Task' }
// ]

const asiacloudtable = ref<AsiacloudDomain[]>()

onMounted(() => {
    getVhosts()
})


const tabs = ref<Asiacloudvhost[]>()

const activeName = ref('first')

const handleClick = async (tab: TabsPaneContext, event: Event) => {
    console.log(tab.paneName, event)
    const data = await GetDomain(<string>tab.paneName)
    console.log(data.data)
    asiacloudtable.value = data.data

}

const getVhosts = async () => {
    const data = await VhostList()
    tabs.value = data.data
}
</script>

<style>
.demo-tabs>.el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
}
</style>