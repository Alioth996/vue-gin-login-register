import { defineStore } from 'pinia'

export const userStore = defineStore('userInfo', {
  state: () => ({
    uid: 0,
    role: 0,
    token: ''
  }),
  getters: {
    GET_UID: state => state.uid,
    GET_TOKEN: state => state.token,
    GET_ROLE: state => state.role
  },
  actions: {
    SET_TOKEN(token: string) {
      this.token = token
    },
    SET_UID(uid: number) {
      this.uid = uid
    },
    SET_ROLE(role: number) {
      this.role = role
    }
  }
})
