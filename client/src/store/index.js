import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

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
      console.log(payload)
      state.username = payload.user.username
      state.isLoggedIn = true
      state.id = payload.user.id
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
      console.log(payload)
      // TODO: send to login API

      return context.dispatch('reload')
    },
    reload(context) {
      alert('reload run')
      const userdata = {
        username: "testuser",
        password: "testpassword"
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



const store = new Vuex.Store({
  modules: {
    auth: authModule
  },
  plugins: [createPersistedState({
    key: 'vuexdata',
    storage: window.sessionStorage
  })]
})

export default store
