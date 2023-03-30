<template>
  
  <el-table :data="domains" style="width: 100%">
    <el-table-column prop="id" label="ID" width="180" />
    <el-table-column prop="name" label="Name" width="180" />
    <el-table-column prop="created_on" label="Created On" width="180" />
    <el-table-column prop="modified_on" label="Modified On" width="180" />
    <el-table-column prop="status" label="Status" width="180" />
    <el-table-column prop="name_server" label="Name Server" width="180" />
  </el-table>

  <button @click="update">update</button>
</template>

<script  lang="ts">
import { defineComponent,ref } from 'vue'
import axios  from 'axios'
interface Domain{
      id: string
      name: string
      created_on:string
      modified_on:string
      name_servers: string[]
      status:string
  }
export default defineComponent({

  setup () {
    const domains = ref<Domain[]>([]);

    function update(){
      axios.get("http://192.168.96.128:8080/cloudflare")
        .then( (response) =>{
            console.log(response.data)
            domains.value=response.data.data as Domain[]
        })
        .catch( (err) => {
            console.log(err)
        })
    }
    return {update,
      domains}
  }
})
</script>

<style scoped>

</style>