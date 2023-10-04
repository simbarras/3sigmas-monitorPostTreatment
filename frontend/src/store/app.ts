// Utilities
import { defineStore } from 'pinia'

export const useActionStore = defineStore('action', {
  state: () => ({
    actions: [],
  }),
})
