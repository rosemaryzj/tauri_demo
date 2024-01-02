import {createPinia, defineStore} from "pinia";
import {ref} from "vue";

export const useSideBarStore = defineStore("sideBar",()=>{
    const collapsed = ref(true)
    function toggle(){
        collapsed.value = !collapsed.value
    }
    return {
        collapsed,
        toggle
    }
})

export const userStore = defineStore("user", ()=>{
    const username = ref("");
    function setUsername(payload){
        username.value = payload;
    }
    return {
        username,
        setUsername,
    }
})


const pinia = createPinia();
export default pinia;