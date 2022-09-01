import request from '.'

export const userLogin = v => request.post('/login', v)
export const userRegister = v => request.post('/register', v)
export const delUser = id => request.delete(`/user/${id}`)
