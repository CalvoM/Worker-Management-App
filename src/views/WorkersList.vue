<template>
  <div class="Workers">
    <Search v-on:search_begin="filterWorkers"/>{{searched_value}}
      <Card v-for="(worker,index) in rendered_workers_list.reverse()" :key="index" :worker="worker" />
  </div>
</template>
<script>
import Card from "@/components/Card.vue";
import Search from "@/components/Search.vue";
import axios from "axios";
import {cors_url,db_url} from "../config";
export default {
  name:"WorkersList",
  components:{
    Card,
    Search
  },
  data(){
    return{
      workers_list:[],
      searched_value:"",
      search_filter:"name"
    }
  },
  methods:{
    filterWorkers:function(value) {
      //handle the search values from the search component upon change of value
      let [searchValue,searchFilter,filterId] = value;
      if(filterId==="0") {
        this.searched_value = searchValue;
        this.search_filter = "name"
        }
        else{
          this.searched_value = searchValue;
        this.search_filter = searchFilter
        }
    }
  },
  mounted:function() {
    axios({
      method:"get",
      url:cors_url+db_url+"/workers/"
    })
    .then((response)=>{
      let workers = response.data.Data
      workers.forEach(worker => {
        this.workers_list.push(worker)
      });

    })
    .catch((err)=>{
      console.log(err.message)
    })
  },
  computed:{
    rendered_workers_list(){
      //Returns filtered list of workers based on the value in the search-box
      return this.workers_list.filter((worker)=>{
       return worker[this.search_filter.toLowerCase()].toString().toLowerCase().includes(this.searched_value.toLowerCase())
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.Workers{
  font-family: 'Manjari';
}
</style>