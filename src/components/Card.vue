<template>
    <div id="Card">
        <p v-for="(key,index) in Object.keys(worker)" :key="index" class="credentials"><span class="tag">{{key}}</span>{{worker[key]}}</p> 
    </div>
</template>
<script>
export default {
    name:"Card",
    props:{
        worker : Object
    }
}
</script>
<style lang="scss" scoped>
#Card{
    padding: 1em;
    border:1px solid #ddd;
    margin:1em auto;
    width:65%;
    display: flex;
}
p{
    display: inline;
    flex:1;
}
.tag{
    color:#aaa;
    margin: auto 0.8em;
}
</style>