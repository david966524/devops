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
        <dev>
            <el-form :inline="true" :model="formInline" class="demo-form-inline">
                <el-form-item label="线路域名">
                    <el-input v-model="formInline.base_url" placeholder="线路域名" />
                </el-form-item>
                <el-form-item label="资源域名">
                    <el-input v-model="formInline.res_url" placeholder="资源域名" />
                </el-form-item>
                <el-form-item label="socket ip">
                    <el-input v-model="formInline.socket_ip" placeholder="socket ip" />
                </el-form-item>
                <el-form-item label="port">
                    <el-select v-model="formInline.socket_port" placeholder="port">
                        <el-option label="9326" :value="parseInt('9326')" />
                        <el-option label="9325" :value="parseInt('9325')" />
                    </el-select>
                </el-form-item>
                <el-form-item label="ssl">
                    <el-select v-model="formInline.ssl" placeholder="ssl">
                        <el-option label="1" :value="parseInt('1')" />
                        <el-option label="2" :value="parseInt('2')" />
                    </el-select>
                </el-form-item>
                <el-form-item label="type">
                    <el-select v-model="formInline.type" placeholder="type">
                        <el-option label="0" :value="parseInt('0')" />
                        <el-option label="1" :value="parseInt('1')" />
                    </el-select>
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="formInline.remark" placeholder="备注" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="addline()">添加</el-button>
                </el-form-item>
            </el-form>
        </dev>
        <el-divider />
        <el-table :data="lines" :edit-config="{ editMethod: 'cell' }">
            <el-table-column property="base_url" label="线路" width="220">
            </el-table-column>
            <el-table-column property="res_url" label="资源线路" width="220" />
            <el-table-column property="socket_ip" label="socket ip" />
            <el-table-column property="socket_port" label="socket port" />
            <el-table-column property="timeout" label="timeout" />
            <el-table-column property="ssl" label="ssl" />
            <el-table-column property="remark" label="备注" />
            <el-table-column property="type" label="type" />
            <el-table-column fixed="right" label="Operations">
                <template #default="scope">
                    <el-button link type="primary" size="small" @click="deleteLine(scope.$index)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-divider content-position="left">修改完记得保存</el-divider>
        <div>
            <el-button type="primary" @click="saveLines()">save</el-button>
        </div>
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
import { AddIm, GetImList, DeleteIm, UpdateIm } from '../api/imapi';
import { GetImLine, ChangeLine } from '../api/imLinesapi'
import { Im, Line } from '../interface';
import { ElMessage } from 'element-plus'

onMounted(() => {
    getIms()
})

const dialogFormVisible = ref(false)
const dialogTableVisible = ref(false)
const dialogEditFormVisible = ref(false)
const im = ref<Im>({
    projectName: "",
    serverip: "",
})

const ims = ref<Im[]>()
const lines = ref<Line[]>([])
const formInline = ref<Line>({
    base_url: "",
    res_url: "",
    socket_ip: "",
    socket_port: 9326,
    timeout: 30000,
    ssl: 2,
    remark: "",
    type: 0
})

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

const showdialog = (row: Im) => {
    dialogEditFormVisible.value = true
    im.value = row
    console.log(im.value)
}

const updataIm = async () => {
    const data = await UpdateIm(im.value)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogEditFormVisible.value = false
        getIms()
    }
}



const getLines = async (row: Im) => {
    dialogTableVisible.value = true
    const data = await GetImLine(row)
    lines.value = <Line[]>data.data
    console.log(lines.value)
    imid.value = <string>row.ID
    formInline.value = {
        timeout: 30000,
        res_url: lines.value[0].res_url,
        socket_ip: lines.value[0].socket_ip
    }
}

const deleteLine = async (index: number) => {
    console.log(index)
    lines.value?.splice(index, 1)
}
const imid = ref<string>()
const saveLines = async () => {
    console.log(lines.value)
    console.log(imid.value)
    const data = await ChangeLine(lines.value, <string>imid.value)
    if (data.code !== 200) {
        ElMessage.error('操作失败')
    }
    ElMessage({
        message: '操作成功',
        type: 'success',
    })
}

const addline = async () => {
    lines.value?.push(formInline.value)
    console.log(lines.value)
    formInline.value = {
        timeout: 30000,
        ssl: 2,
        type: 0,
    }
}


</script>

<style scoped></style>