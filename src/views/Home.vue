<template>
  <div class="home">
    <Form v-on:submit="sendWorkerDetails"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Form from '@/components/Form.vue';
import axios from "axios";
import {cors_url,db_url} from "../config";
axios.defaults.headers.post['Content-Type']='application/json'
export default {
  name: 'home',
  components: {
    Form
  },
  methods:{
    sendWorkerDetails: function(worker){
      axios({
        method:"post",
        url:cors_url+db_url+'registerworker/',
        data:{
        name:worker.name,
        age:+(worker.age),
        email:worker.email,
        salary:+(worker.salary)
      }},
      )
      .then((response)=>{
        console.log(response)
      })
      .catch((err)=>{
        console.log(err.message)
      })
    }
  }
}
</script>
