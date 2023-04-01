<template>
    <div class="container">
       
        <!-- input -->
        <div class="search-bar">
            <el-input v-model="queryInput" placeholder="请输入要查询的域名" />
            <el-button type="primary" plain @click="queryDomain">查询域名</el-button>
            <el-button type="warning" plain @click="clearQuery">清空输入框</el-button>
        </div>

        <!-- button -->
        <div class="button-group">
            <el-button type="primary" plain @click="getAllDomain">获取所有域名</el-button>
            <el-button type="success" plain @click="showFrom">新增域名</el-button>
        </div>

        <!-- table -->
        <el-table v-loading="loading" :data="domains" style="width: 100%">
            <!-- <el-table-column prop="id" label="Id" width="180" /> -->
            <el-table-column prop="name" label="Name" width="180" />
            <el-table-column prop="created_on" label="Created" width="200" />
            <el-table-column prop="modified_on" label="Modified" width="200" />
            <el-table-column prop="name_servers" label="Name Servers" width="350" />
            <el-table-column prop="status" label="Status" width="120" />
            <el-table-column fixed="right" label="Operations" width="200">
                <template #default="scope">
                    <el-button link type="primary" size="small" @click="show(scope.row)">Show</el-button>
                    <el-button link type="primary" size="small" @click="showAddDnsDialog(scope.row)">Edit</el-button>
                    <el-button link type="primary" size="small" @click="deleteDomain(scope.row)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>

        <!-- add dialog -->
        <el-dialog v-model="dialogFormVisible" title="添加域名">
            <el-form :model="tableform">
                <el-form-item label="Domain name" :label-width="120">
                    <el-input v-model="tableform.name" autocomplete="off" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">关闭</el-button>
                    <el-button type="primary" @click="addDomainName">添加</el-button>
                </span>
            </template>
        </el-dialog>

        <!-- show dialog -->
        <el-dialog v-model="showDialogTableVisible" title="DNS 解析">
            <el-table :data="dnsList">
                <el-table-column property="type" label="Type" width="80" />
                <el-table-column property="name" label="Name" width="250" />
                <el-table-column property="content" label="Content" width="350" />
                <el-table-column property="proxied" label="Proxy" width="100" />
                <el-table-column property="ttl" label="TTL(1=AUTO)" width="150" />
                <el-table-column fixed="right" label="" width="50">
                    <template #default="scope">
                        <el-button link type="primary" size="small" @click="deleteDNSRecord(scope.row)">Delete</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>

        <!-- show edit form -->
        <el-dialog v-model="showDialogEditVisible" title="DNS 解析">
            <el-form :model="dnsForm" label-width="120px">
                <el-form-item label="Type">
                    <el-select v-model="dnsForm.type" placeholder="请选择类型">
                        <el-option label="A" value="A" />
                        <el-option label="CNAME" value="CNAME" />
                        <el-option label="TXT" value="TXT" />
                    </el-select>
                </el-form-item>
                <el-row>
                    <el-form-item label="Domain name">
                        <el-input v-model="dnsForm.name" />
                    </el-form-item>
                </el-row>
                <el-form-item label="Proxy">
                    <el-switch v-model="dnsForm.proxied" />
                </el-form-item>
                <el-row>
                    <el-form-item label="Content">
                        <el-input v-model="dnsForm.content" />
                    </el-form-item>
                </el-row>
            </el-form>
            <el-button type="primary" @click="addDNSRecord">添加</el-button>
        </el-dialog>




        <!-- back to top button -->
        <el-backtop :right="100" :bottom="100" />
    </div>
</template>
  
  
<script setup lang="ts">

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { defineComponent } from 'vue';
import { onMounted } from 'vue'
import useUserStore from '../store/index'
import { Domain, DNS } from '../interface/index'
import { GetAllDomain, QueryDoaminReq, AddDomainNameReq, ShowDnsReq, DeleteDomainReq, AddDNSRecordReq, DeleteDNSRecordReq } from '../api/cloudflare'



// 生命周期钩子
onMounted(() => {
    console.log('mounted in the composition api!')
    getAllDomain()
})


// 域名对象列表
const domains = ref<Domain[]>([]);

// DNS对象列表
const dnsList = ref<DNS[]>([]);

// 查询框
let queryInput = ref<string>()

//初始化 新增域名 表单
const tableform = ref({
    name: "",
})

// 加载画面
const loading = ref(false)





// 显示 添加域名 add dialog 
let dialogFormVisible = ref(false)
// 显示 修改域名 edit dialog
let showDialogTableVisible = ref(false)




//function

// 清空 查询框方法
const clearQuery = () => {
    queryInput.value = ""
}

//显示 新增域名 表单  并初始化 name
const showFrom = () => {
    dialogFormVisible.value = true
    tableform.value.name = ""
}


// 获取所有域名
const getAllDomain = async () => {
    console.log("getAllDomain")
    loading.value = true
    const data = await GetAllDomain()
    domains.value = <Domain[]>data.data
    loading.value = false
}

// 添加域名方法
const addDomainName = async () => {
    console.log(tableform.value.name)
    const data = await AddDomainNameReq(tableform.value.name)
    if (data.code = 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogFormVisible.value = false
        domains.value = <Domain[]>data.data
    }

}

//点击 show 方法
const show = async (row: DNS) => {
    showDialogTableVisible.value = true
    console.log(row.name)
    const data = await ShowDnsReq(row.name)
    if (data.code !== 200) {

    }
    dnsList.value = <DNS[]>data.data
}

// 查询单个域名
const queryDomain = async () => {
    console.log(queryInput.value)
    if (queryInput.value === undefined) {
        ElMessage({
            showClose: true,
            message: '请输入要查询的域名',
            type: 'warning',
        })
        return
    }
    const data = await QueryDoaminReq(<string>queryInput.value)
    domains.value = <Domain[]>data.data
}

//删除 域名
const deleteDomain = async (row: Domain) => {
    console.log(row.id)
    const data = await DeleteDomainReq(row.id)
    if (data.code = 200) {
        ElMessage({
            message: row.name + ' 删除成功！ 请刷新页面~',
            type: 'success',
        })
    }
}
//添加 dns解析

// 初始化 编辑dns dialog 框为 flase
const showDialogEditVisible = ref(false)
const addDnsId = ref()
//初始化 编辑dns form 数据
const dnsForm = ref<DNS>({
    id: "",
    type: "",
    name: "",
    content: "",
    proxied: false,
    ttl: 600,
})

const showAddDnsDialog = (row: any) => {
    console.log(row.id)
    showDialogEditVisible.value = true
    addDnsId.value = row.id
}

const addDNSRecord = async () => {
    console.log(dnsForm.value)
    const data = await AddDNSRecordReq(addDnsId.value, dnsForm.value)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        showDialogEditVisible.value = false
        dnsForm.value.name = ""
        return
    }
    ElMessage({
        showClose: true,
        message: data.msg,
        type: 'error',
    })
}


//删除 dns解析
const deleteDNSRecord = async (row: any) => {
    console.log(row)
    console.log(row.id)
    console.log(row.zone_id)
    const data = await DeleteDNSRecordReq(row.id, row.zone_id)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
    }
}




</script>
  
<style scoped>
.container {
    width: 80%;
    margin: 0 auto;
    padding-top: 20px;
}

.search-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.button-group {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.el-table-column--link {
    color: #409eff;
    cursor: pointer;
}

.dialog-footer {
    display: flex;
    justify-content: flex-end;
}

.el-switch {
    margin-left: 10px;
}
</style>