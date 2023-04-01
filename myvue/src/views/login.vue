<template>
    <div class="login-container">
        <el-card class="login-box">
            <div class="title">用户登录</div>
            <el-form :model="Userform" ref="loginForm" :rules="rules" label-width="80px" class="login-form">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model.trim="Userform.username" placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input type="password" v-model.trim="Userform.password" placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="Login" class="login-btn">登录</el-button>
                    <el-button @click="handleReset" class="reset-btn">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script setup lang="ts" >
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import useUserStore from '../store/index'
import {UserData} from '../interface/index'
import { loginReq} from "../api/login"


const router = useRouter();
const Userform = reactive<UserData>({
    username: '',
    password: '',
});


const rules = {
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
};

// 实例化 store
const userStore = useUserStore()

const Login = async () => {

    await useUserStore().login(Userform)
    if (useUserStore().token !==""){
         router.push("home")
    }
    
    console.log(useUserStore().token)
    console.log(useUserStore().username)
}
const handleReset = () => {
    // 重置表单
    Userform.username = ''
    Userform.password = ''


};








</script>

<style scoped>
.login-container {
    height: 100vh;
    background-color: #f0f2f5;
    display: flex;
    justify-content: center;
    align-items: center;
}

.login-box {
    width: 420px;
    margin: 30px auto;
    padding: 30px;
    border-radius: 10px;
    box-shadow: 0 0 30px rgba(0, 0, 0, 0.1);
}

.title {
    font-size: 26px;
    font-weight: bold;
    text-align: center;
    margin-bottom: 20px;
}

.login-form .el-form-item__label {
    color: #333;
    font-weight: bold;
}

.login-btn {

    display: flex;
    justify-content: center;
}

.reset-btn {

    display: flex;
    justify-content: center;
}
</style>