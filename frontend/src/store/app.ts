// Utilities
import {defineStore} from 'pinia'
import axios from "axios";
import {computed, ref} from "vue";
import {Action} from "@/plugins/data";

// const API_URL = 'http://pub-3sigmas:3001' // Production
const API_URL = 'http://localhost:3001' // Development
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
                actionList.value = res.data as Action[]
            })
    }

    function uploadAction(action: any){
        axios.post(`${API_BASE}/action`, action)
            .then(res => {
                actionList.value = res.data as Action[]
            })
    }

    function deleteAction(action: any){
        console.log(action)
        axios.delete(`${API_BASE}/action/${action.idUuid}`)
            .then(res => {
                console.log(res.data)
                actionList.value = res.data as Action[]
                console.log(actionList.value)
            })
            .catch(err => {
                console.log(err)
            })
    }

    function preAddAction() {
        const a = new Action()
        a.active = false
        a.bucketName = ""
        a.equationName = ""
        a.listVariables = ""
        actionList.value.push(a)
    }

    function triggerAction(action: any){
        axios.post(`${API_BASE}/trigger/action/${action.idUuid}`)
            .then(res => {
                console.log(res.data)
            })
            .catch(err => {
                console.log(err)
            })
    }

    return {actionList, getActionList, fetchActions, uploadAction, deleteAction, preAddAction, triggerAction}

})

export const useBucketStore = defineStore('bucket', () => {
    const bucketList = ref([] as string[])
    const getBucketList = computed(() => {
        return bucketList.value
    })

    function fetchBuckets() {
        axios.get(`${API_BASE}/bucket`)
            .then(res => {
                bucketList.value = res.data as string[]
            })
    }

    return {bucketList, getBucketList, fetchBuckets}
})

export const useEquationStore = defineStore('equation', () => {
    const equationList = ref([] as string[])
    const getEquationList = computed(() => {
        return equationList.value
    })

    function fetchEquations() {
        axios.get(`${API_BASE}/function`)
            .then(res => {
                equationList.value = res.data as string[]
            })
    }

    return {equationList, getEquationList, fetchEquations}
})
