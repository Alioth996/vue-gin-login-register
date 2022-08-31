import request from '.'

export const userLogin = v => request.post('/login', v)
export const userRegister = v => request.post('/register', v)
