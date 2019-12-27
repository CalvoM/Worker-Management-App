<template>
    <div id="Search">
            <input type="text" name="search" id="search" placeholder="Search..." v-model="search_value" @input="searchBegin">
            Search by:
            <select name="search_tag" id="search_tag" v-model="search_filter_value">
                <option v-for="(option,id) in options" :value="id" :key="id">{{option}}</option>
            </select>
    </div>
</template>
<script>
export default {
   name:"Search",
   data(){
       return{
           search_value:"",
           search_filter_value:"0",
           options:["Choose search filter","Name","Age","Email","Salary"]
       }
   },
   methods:{
       searchBegin:function(){
           this.$emit('search_begin',[this.search_value,this.options[+this.search_filter_value],this.search_filter_value])
       }
   }
}
</script>
<style lang="scss" scoped>
#Search{
    width:80%;
    margin:auto;
}
#search{
    padding: 0.6em;
    width:30%;
    font-family: Manjari;
    font-size: 105%;
}
select{
    margin-left: 1em;
}
</style>