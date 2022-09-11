import request from '.'

export const userLogin = v => request.post('/login', v)
export const userRegister = v => request.post('/register', v)
export const delUser = id => request.delete(`/user/${id}`)
export const getUserRole = () => request.get('/user/role')
export const getUserInfo = () => request.get('/user/info')
