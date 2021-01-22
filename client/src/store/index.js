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
      console.log("in sert",payload)
      state.username = payload.user_name
      state.isLoggedIn = true
      state.id = payload.user_id
      console.log("ログインデータ", state)
    },
    clear(state) {
      state.username = ''
      state.isLoggedIn = false
      state.id = false
    }
  },
  actions: {
    login(context, payload) {

      console.log("logindata ",payload)
      // TODO: send tdo login API

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
    info: state => state.info,
    warnings: state => state.warnings,
    error: state => state.error,
  },
  mutations: {
    set(state, payload) {
      if (payload.success) {
        state.success = payload.success
      }
      if (payload.info) {
        state.info = payload.info
      }
      if (payload.warnings) {
        state.warnings = payload.warnings
      }
      if (payload.error) {
        state.error = payload.error
      }
    },
    clear(state) {
      state.success = ''
      state.info = ''
      state.warnings = []
      state.error = ''
    }
  },
  actions: {
    setSuccessMessage(context, payload) {
      context.commit('set', {
        'success': payload.message
      })
      setTimeout(() => {
        context.dispatch('clearMessages')
      }, 1500)
    },
    setInfoMessage(context, payload) {
      context.commit('set', {
        'info': payload.message
      })
      setTimeout(() => {
        context.dispath('clearMessages')
      }, 1500)
    },
    setWarningMessages(context, payload) {
      context.commit('set', {
        'warnings': payload.messages
      })
      setTimeout(() => {
        context.dispath('clearMessages')
      }, 1500)
    },
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
