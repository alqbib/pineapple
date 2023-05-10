import axios from 'axios'
import { message as $message } from 'ant-design-vue';



axios.defaults.baseURL = import.meta.env.VITE_GLOB_API_URL 
axios.defaults.timeout = 6000

// 请求头，headers 信息
axios.defaults.headers['Accept'] = 'application/json'

// const token = sessionStorage.getItem("search-token")
// axios.defaults.headers['Authorization'] = (token??'') != ''?token:''
axios.defaults.headers.post['Content-Type'] = 'application/json;charset=utf-8'


axios.interceptors.response.use(res => {
  if (typeof res.data !== 'object') {

    $message.error('提示: 服务器异常');
    
    return Promise.reject(res)
  }
  if (res.data.code != 200) {

    $message.error('提示:' + res.data.msg)
    
    return Promise.reject(res.data)
  }

  return res.data
},
(error=>{
  $message.error('提示:' + error.message)
  return error
}))

export default axios