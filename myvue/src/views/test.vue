<template>
    <div>
        <p>{{ username }}</p>
        <p>{{ age }}</p>
        <p></p>
    </div>
    <button @click="change">click</button>
</template>

<script  lang="ts">

import{ ref ,reactive ,toRef ,toRefs,defineComponent} from 'vue'
export default defineComponent({
    setup(){

        const userinfo = reactive({
            username: "david1",
            age: "23",
            sex: "man"
        })

        const userinforefs = toRefs(userinfo)
        const change= ()=>{
            userinforefs.username.value="1222"
        }

        interface User {
            name: string
            age:string
            email: string
        }
        const myuser = reactive<User>({
            name: "haha",
            age: "23",
            email: "djjjs@mail.com"
        })
        console.log(myuser)
        const username = toRef(myuser,"name")
        username.value="david"
        console.log(myuser)
        

        const myusertorefs = toRefs(myuser)
        const {name,age,email}=myusertorefs

        console.log(myusertorefs.name)
        console.log(age)
        console.log(email)
        email.value="sssss@gamil.com"
        console.log(myuser)

        return{
            ...userinforefs,
            change
        }
    }
})


</script>

<style scoped>

</style>