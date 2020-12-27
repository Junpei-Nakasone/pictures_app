import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import axios from 'axios'
import api from '@/services/api'

Vue.use(Vuex)

// 認証情報
const authModule = {
  // strict: true,
  namespaced: true,
  state: {
    username: '',
    isLoggedIn: false,
    id: '',
  },
  getters: {
    username: state => state.username,
    isLoggedIn: state => state.isLoggedIn,
    id: state => state.id,
  },
  mutations: {
    set(state, payload) {
      alert('set run')
      alert('set check')
      // console.log(payload)
      console.log("in sert",payload)
      state.username = payload.user_name
      state.isLoggedIn = true
      state.id = payload.user_id
      console.log("ログインデータ", state)
      alert('set終了')
    },
    clear(state) {
      state.username = ''
      state.isLoggedIn = false
      state.id = false
    }
  },
  actions: {
    login(context, payload) {
      alert('login on store running')
      console.log("logindata ",payload)
      // TODO: send to login API

      return api.post('/login', {
        'user_name': payload.username,
        'password': payload.password,
      })
      .then((res) => {
        console.log("then in loginmethod",res.data)
        return context.commit('set',res.data)
        // context.dispatch('reload')
        // .then(user => user)
      })
      .catch(error => error.response)
    },
    reload(context) {
      alert('reload run')
      console.log('reloadデータ', context.user_name)
      const userdata = {
        username: context.username,
        password: context.password,
      }
      return context.commit('set', {
        user: userdata
      })
    },
    logout(context) {
      alert('logout vuex')
      return context.commit('clear')
    }
  }
}

const messageModule = {
  namespaced: true,
  state: {
    success: '',
    info: '',
    warnings: '',
    error: '',
  },
  getters: {
    success: state => state.success,
    info: state => state.warnings,
    error: state => state.error,
  },
  mutations: {
    set(state, payload) {
      if (payload.error) {
        alert("setError", payload)
        console.log("setError", payload)
        state.error = payload.error
      }
    },
    clear(state) {
      state.error = ''
    }
  },
  actions: {
    setErrorMessage(context, payload) {
      console.log("setErrorM",payload)
      context.commit('set', {
        'error': payload.message
      })
      setTimeout(() => {
        context.dispatch('clearMessage')
      }, 1500)
    },
    // メッセージ削除
    clearMessage(context) {
      context.commit('clear')
    }
  }
}


const store = new Vuex.Store({
  modules: {
    auth: authModule,
    message: messageModule,
  },
  plugins: [createPersistedState({
    key: 'vuexdata',
    storage: window.sessionStorage
  })]
})

export default store
