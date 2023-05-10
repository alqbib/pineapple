import { createStore } from 'vuex'

const store = createStore({
  state:{
    token: ''
  },
  mutations: {
    updateToken(state, token){
      state.token = token
    }
  }
})

export default store