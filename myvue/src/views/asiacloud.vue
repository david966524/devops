<template>
    <el-button type="primary" @click="dialogFormVisible = true">添加域名</el-button>
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane v-for="(tab, index) in tabs" :key="index" :label="tab.name" :name="tab.name">{{ tab.name
        }}</el-tab-pane>
        <el-table :data="asiacloudtable" style="width: 100%">
            <el-table-column prop="domain" label="domain" />
            <el-table-column prop="host" label="host" />
            <el-table-column prop="vhost" label="vhost" />
            <el-table-column prop="status" label="status" />
            <el-table-column prop="cname" label="cname" />
            <el-table-column fixed="right" label="Operations" width="120">
                <template #default="scope">
                    <el-button link type="primary" size="small" @click="">Edit</el-button>
                    <el-button link type="primary" size="small"  @click="deleteDomain(scope.row)">delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-tabs>

    <!-- add form -->
    <el-dialog v-model="dialogFormVisible" title="添加">
        <el-form :model="formdomain">
            <el-form-item label="域  名" label-width="100">
                <el-input v-model="formdomain.domain" autocomplete="off" />
            </el-form-item>
            <el-form-item label="回源地址" label-width="100">
                <el-input v-model="formdomain.host" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Zones" label-width="100">
                <el-select v-model="formdomain.vhost" placeholder="Please select a zone">
                    <el-option label="my101-site-03" value="my101-site-03" />
                    <el-option label="my101-site-04" value="my101-site-04" />
                    <el-option label="my101-site-05" value="my101-site-05" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogFormVisible = false">Cancel</el-button>
                <el-button type="primary" @click="addDomain">
                    Confirm
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted ,onBeforeMount} from 'vue'
import type { TabsPaneContext } from 'element-plus'
import { Asiacloudvhost, AsiacloudDomain } from '../interface/index'
import { VhostList, GetDomain, AddDomain,DeleteDomain } from '../api/asiacloudapi'
import { ElMessage } from 'element-plus'


onBeforeMount(() => {
    getVhosts()
})
const formdomain = ref<AsiacloudDomain>({
    domain: "",
    host: ""
})
const dialogFormVisible = ref(false)
const asiacloudtable = ref<AsiacloudDomain[]>()
const tabs = ref<Asiacloudvhost[]>()

const activeName = ref('first')

const handleClick = async (tab: TabsPaneContext, event: Event) => {
    console.log(tab.paneName, event)
    const data = await GetDomain(<string>tab.paneName)
    console.log(data.data)
    asiacloudtable.value = data.data

}

const addDomain = async () => {
    console.log(formdomain.value)
    const data = await AddDomain(formdomain.value)
    console.log(data)
    if (data.code !== 200) {
        ElMessage.error(data.msg)
        return
    }
    ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogFormVisible.value=true
}

const deleteDomain = async (row:AsiacloudDomain)=>{
    console.log("id = "+row.id)
    console.log("vhost = "+row.vhost)
    const data = await DeleteDomain(row)
    if (data.code !=200 ){
        ElMessage.error(data.msg)
        return
    }
    ElMessage({
            message: data.msg,
            type: 'success',
        })
    
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