<template>
    <el-button type="primary" @click="dialogFormVisible = true">添加新包网</el-button>

    <el-table :data="ims" style="width: 100%">
        <el-table-column fixed prop="ID" label="Id" width="80" />
        <el-table-column prop="projectName" label="平台名称" width="200" />
        <el-table-column prop="serverip" label="服务器IP" width="200" />
        <el-table-column prop="jsonconfig" label="JSON配置地址" width="400" />
        <el-table-column prop="remark" label="备注" width="400" />
        <el-table-column fixed="right" label="Operations" width="300">
            <template #default="scope">
                <el-button link type="primary" size="small" @click="getLines(scope.row)">查看线路</el-button>
                <el-button link type="primary" size="small" @click="showdialog(scope.row)">修改</el-button>
                <el-button link type="primary" size="small" @click="deleteIm(scope.row)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <!--add dialog -->
    <el-dialog v-model="dialogFormVisible" title="添加新包网">
        <el-form :model="im" label-width="120px">
            <el-form-item label="包网名">
                <el-input v-model="im.projectName" />
            </el-form-item>
            <el-form-item label="服务器ip">
                <el-input v-model="im.serverip" />
            </el-form-item>
            <el-form-item label="配置文件">
                <el-input v-model="im.jsonconfig" />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="im.remark" type="textarea" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogFormVisible = false">Cancel</el-button>
                <el-button type="primary" @click="addIm">
                    添加
                </el-button>
            </span>
        </template>
    </el-dialog>

    <!-- show line table  -->
    <el-dialog v-model="dialogTableVisible" title="线路表" width="70%">
        <el-table :data="lines">
            <el-table-column property="base_url" label="线路"  width="220"/>
            <el-table-column property="res_url" label="资源线路" width="220" />
            <el-table-column property="socket_ip" label="socket ip" />
            <el-table-column property="socket_port" label="socket port" />
            <el-table-column property="timeout" label="timeout" />
            <el-table-column property="ssl" label="ssl" />
            <el-table-column property="remark" label="备注" />
            <el-table-column property="type" label="type" />
        </el-table>
    </el-dialog>

          <!-- editdialog -->
          <el-dialog v-model="dialogEditFormVisible" title="编辑新包网">
        <el-form :model="im" label-width="120px">
            <el-form-item label="包网名">
                <el-input v-model="im.projectName" />
            </el-form-item>
            <el-form-item label="服务器ip">
                <el-input v-model="im.serverip" />
            </el-form-item>
            <el-form-item label="json配置地址">
                <el-input v-model="im.jsonconfig" />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="im.remark" type="textarea" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogEditFormVisible = false">Cancel</el-button>
                <el-button type="primary" @click="updataIm">
                    修改
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { AddIm, GetImList, DeleteIm ,UpdateIm} from '../api/imapi';
import { GetImLine } from '../api/imLinesapi'
import { Im, Line } from '../interface';
import { ElMessage } from 'element-plus'

onMounted(() => {
    getIms()
})

const dialogFormVisible = ref(false)
const dialogTableVisible = ref(false)
const dialogEditFormVisible =ref(false)
const im = ref<Im>({
    projectName: "",
    serverip: "",
})

const ims = ref<Im[]>()
const lines = ref<Line[]>()

const deleteIm = async (row: Im) => {
    const data = await DeleteIm(row)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogFormVisible.value = false
        getIms()
    }
}

const addIm = async () => {
    const data = await AddIm(im.value)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogFormVisible.value = false
        getIms()
    }
}

const getIms = async () => {
    const data = await GetImList()
    ims.value = data.data
}

const showdialog = (row:Im)=>{
    dialogEditFormVisible.value=true
    im.value = row
    console.log(im.value)
}

const updataIm = async ()=>{
    const data = await UpdateIm(im.value)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogEditFormVisible.value=false
        getIms()
    }
}



const getLines = async (row: Im) => {
    dialogTableVisible.value = true
    const data = await GetImLine(row)
    lines.value = data.data
    console.log(lines.value)
}


</script>

<style scoped></style>