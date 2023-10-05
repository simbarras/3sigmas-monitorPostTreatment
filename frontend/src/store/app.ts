// Utilities
import {defineStore} from 'pinia'
import axios from "axios";
import {computed, ref} from "vue";
import {Action} from "@/plugins/data";

const API_URL = 'http://localhost:3001'
const API_VERSION = 'v0'
const API_BASE = `${API_URL}/api/${API_VERSION}`

export const useActionStore = defineStore('action', () => {

    const actionList = ref([] as Action[])

    const getActionList = computed(() => {
        return actionList.value
    })

    function fetchActions() {
        axios.get(`${API_BASE}/action`)
            .then(res => {
                console.log(res.data)
                actionList.value = res.data as Action[]
                console.log(actionList.value)
            })
    }

    return {actionList, getActionList, fetchActions}

})


