<template>
  <el-button type="primary" plain @click="showDialog">添加用户</el-button>
  <el-table :data="Users" style="width: 100%">

    <el-table-column prop="UserName" label="用户名" width="120" />
    <el-table-column prop="PassWord" label="PassWord" width="120" />
    <el-table-column prop="CreatedAt" label="创建日期" width="300" />
    <el-table-column prop="UpdatedAt" label="更新日期" width="300" />

    <el-table-column fixed="right" label="Operations" width="120">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="showdialogedit(scope.row)">修改密码</el-button>
      </template>
    </el-table-column>
  </el-table>

  <!-- 添加用户 dialog -->
  <el-dialog v-model="dialogFormVisible" title="添加用户">
    <el-form :model="adduser">
      <el-form-item label="用户名" :label-width="120">
        <el-input v-model="adduser.UserName" autocomplete="off" />
      </el-form-item>
      <el-form-item label="密码" :label-width="120">
        <el-input v-model="adduser.PassWord" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">关闭</el-button>
        <el-button type="primary" @click="addUser">添加</el-button>
      </span>
    </template>
  </el-dialog>

  <!-- 修改用户 dialog -->
  <el-dialog v-model="dialogeditFormVisible" title="修改密码">
    <el-form :model="adduser">
      <el-form-item label="修改密码" :label-width="120">
        <el-input v-model="adduser.PassWord" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer="data">
      <span class="dialog-footer">
        <el-button @click="dialogeditFormVisible = false">关闭</el-button>
        <el-button type="primary" @click="edit()">修改</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { User } from "../interface/index"
import { GetUserList, AddUser, ChangePassword } from "../api/userapi"
import { ref } from "vue"
import { onMounted } from 'vue'
import { ElMessage } from 'element-plus'


// 生命周期钩子
onMounted(() => {
  console.log('mounted in the composition api!')
  getuserList()
})
const adduser = ref<User>({
  UserName: "",
  PassWord: ""
})
const Users = ref<User[]>()
const dialogFormVisible = ref(false)
const dialogeditFormVisible = ref(false)

const addUser = async () => {
  console.log(adduser.value)
  const data = await AddUser(adduser.value)
  if (data.code = 200) {
    ElMessage({
      message: data.msg,
      type: 'success',
    })
  }
}

const getuserList = async () => {
  const data = await GetUserList()
  Users.value = data.data
}


const showDialog = () => {
  dialogFormVisible.value = true
}

const showdialogedit = (row: User) => {
  dialogeditFormVisible.value = true
  adduser.value = row
}


const edit = async () => {
  console.log(adduser.value)
  const data = await ChangePassword(adduser.value)
  if (data.code = 200) {
    ElMessage({
      message: data.msg,
      type: 'success',
    })
  }


}
const handleDelete = (row: any) => { }
</script>

<style scoped></style>