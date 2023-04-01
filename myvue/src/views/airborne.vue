<template>
    <el-button type="primary" @click="dialogFormVisible = true">添加新包网</el-button>

    <el-table :data="airbornelist" style="width: 100%">
        <el-table-column fixed prop="ID" label="Id" width="150" />
        <el-table-column prop="projectName" label="平台名称" width="200" />
        <el-table-column prop="serverip" label="服务器IP" width="200" />
        <el-table-column prop="remark" label="备注" width="400" />
        <el-table-column fixed="right" label="Operations" width="300">
            <template #default="scope">
                <el-button link type="primary" size="small" @click="showeditProject(scope.row)">修改</el-button>
                <el-button link type="primary" size="small" @click="deleteProject(scope.row)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <!--add dialog -->
    <el-dialog v-model="dialogFormVisible" title="添加新包网">
        <el-form :model="airborne" label-width="120px">
            <el-form-item label="包网名">
                <el-input v-model="airborne.projectName" />
            </el-form-item>
            <el-form-item label="服务器ip">
                <el-input v-model="airborne.serverip" />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="airborne.remark" type="textarea" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogFormVisible = false">Cancel</el-button>
                <el-button type="primary" @click="addProject">
                    添加
                </el-button>
            </span>
        </template>
    </el-dialog>

        <!-- editdialog -->
        <el-dialog v-model="dialogEditFormVisible" title="编辑新包网">
        <el-form :model="airborne" label-width="120px">
            <el-form-item label="包网名">
                <el-input v-model="airborne.projectName" />
            </el-form-item>
            <el-form-item label="服务器ip">
                <el-input v-model="airborne.serverip" />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="airborne.remark" type="textarea" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogEditFormVisible = false">Cancel</el-button>
                <el-button type="primary" @click="editProject">
                    修改
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>




<script setup lang="ts">
import { ref, onMounted } from "vue"
import { Airborne } from "../interface/index"
import { GetAirborneList, AddAirborne ,EditAirborne, DeleteAirborne } from "../api/airborneapi"
import { ElMessage } from 'element-plus'

onMounted(() => {
    getall()
})

const dialogFormVisible = ref(false)
const dialogEditFormVisible = ref(false)

const airbornelist = ref<Airborne[]>()

const airborne = ref<Airborne>({
    projectName: "",
    serverip: "",
    remark: "",
})

const getall = async () => {
    const data = await GetAirborneList()
    airbornelist.value = data.data
}

const addProject = async () => {
    const data = await AddAirborne(airborne.value)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogFormVisible.value=false
        getall()
    }
}

const showeditProject = async  (row:Airborne)=>{
    dialogEditFormVisible.value=true
    airborne.value = row
    console.log(airborne.value.ID)
}
const editProject =  async ()=>{
    const data =  await EditAirborne(airborne.value)
    if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        dialogEditFormVisible.value=false
    }
}

const deleteProject = async (row:Airborne) =>{
    
    const data = await DeleteAirborne(row)
        if (data.code == 200) {
        ElMessage({
            message: data.msg,
            type: 'success',
        })
        getall()
    }
}



</script>

<style scoped></style>