import axios from 'axios'
import { ElMessage } from 'element-plus'
const instance = axios.create({
  baseURL: 'http://127.0.0.1:5000/api/v1',
  timeout: 1000
})
// 添加请求拦截器
instance.interceptors.request.use(
  config => {
    // 在发送请求之前做些什么
    // 加token
    const token = `Bearer ${localStorage.getItem('token')}` ?? ''
    config.headers.Authorization = token
    return config
  },
  error => {
    // 对请求错误做些什么
    return Promise.reject(error)
  }
)

// 添加响应拦截器
instance.interceptors.response.use(
  res => {
    const { data } = res
    return data
  },
  error => {
    ElMessage.error(error.response.data.msg)
    return Promise.reject(error)
  }
)
export default instance
