import { defineStore } from 'pinia'

export const userStore = defineStore('userInfo', {
  state: () => ({
    uid: 0,
    role: 0,
    token: '',
    username: '张三'
  }),
  getters: {
    GET_UID: state => (state.uid ? state.uid : localStorage.getItem('uid')),
    GET_USERNAME: state => (state.username ? state.username : localStorage.getItem('username')),
    GET_TOKEN: state => (state.token ? state.token : localStorage.getItem('token')),
    GET_ROLE: state => (state.role ? state.role : localStorage.getItem('role'))
  },
  actions: {
    SET_TOKEN(token: string) {
      localStorage.setItem('token', token)
      this.token = token
    },
    SET_UID(uid: number) {
      localStorage.setItem('uid', uid.toString())
      this.uid = uid
    },
    SET_ROLE(role: number) {
      localStorage.setItem('role', role.toString())
      this.role = role
    }
  }
})
