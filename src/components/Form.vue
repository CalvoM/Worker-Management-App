<template>
<div id="Form">
     <form id="form" v-on:submit.prevent="registerWorker" >
        <header>Please input the details of the new worker:</header>
        <input type="text" name="workerName" id="workerName" placeholder="Enter worker's name" v-model="Worker.name" required>
        <input type="email" name="workerEmail" id="workerEmail" placeholder="Enter worker's email" v-model="Worker.email" required>
        <input type="number" name="workerAge" id="workerAge" placeholder="Enter worker's age" v-model="Worker.age" required min="18">
        <input type="range" name="workerSalary" id="workerSalary" v-model="Worker.salary" min="1000" max="100000" step="500" required>
        <span v-show="Worker.salary">Worker Salary = {{Worker.salary}}</span>
        <input type="submit" value="Submit">
  </form>
  <router-link to="/worker_list/" id="workerLink">Check <span v-if="worker_list_updated" class="updated">updated</span> list of workers</router-link>
</div>
</template>
<script>
export default {
  name:"Form",
  data(){
    return{
        Worker:{
            name:"",
            email:"",
            age:"",
            salary:""
        },
        worker_list_updated:false
    }
  },
  methods:{
      registerWorker: function(){
          //Sends to server to be
          if(this.Worker.salary==="") alert("Please provide the salary amount")
          else {
                this.$emit("submit",this.Worker)
                this.Worker.name=""
                this.Worker.age=""
                this.Worker.email=""
                this.Worker.salary=""
                this.worker_list_updated = true;
              }
       

      }
  }
}
</script>
<style lang="scss" scoped>
div#Form{
    width:60%;
    margin:auto;
    padding:0.5em;
    font-family: 'Manjari'
}
form{
    margin-bottom: 2em;
}
input{
    padding:0.9em;
    display: block;
    margin:1em auto;
    font-size:100%;
    font-family: 'Manjari'
}
input[type='submit']{
    background-color: green;
    outline: none;
    border:none;
    color:white;
}
input:not(:last-child){
    width:50%;
}
input:required:focus:invalid{
    border:none;
    outline: none;
    box-shadow:1px 1px 2px 2px red;
}
#workerLink{
    text-decoration: none;
    font-size: 110%;
}
.updated{
    font-size:115%;
    color:red;
}
</style>