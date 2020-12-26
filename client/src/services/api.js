import axios from 'axios'
import store from '@/store'

const api = axios.create({
  baseURL: process.env.VUE_APP_API_ENDPOINT,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  }
})


// API共通処理
// トークン認証など入れたい


export default api
