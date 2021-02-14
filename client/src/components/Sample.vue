<template>
  <div>
    <div>this is sample20200209HTTPS</div>
    <div>
      {{msg}}
    </div>
    <div>
      {{VUE_APP_API_ENDPOINT}}
    </div>
    <div>
      {{VUE_APP_SERVICE_URL}}
    </div>
  </div>
</template>

<script>
import api from '@/services/api'
import axios from 'axios'

export default {
  // components: {
  //   PostList,
  // },
  data() {
    return {
      VUE_APP_API_ENDPOINT: process.env.VUE_APP_API_ENDPOINT,
      msg: "",
      VUE_APP_SERVICE_URL: process.env.VUE_APP_SERVICE_URL
    }
  },
  async mounted() {
    // axios.get('http://localhost:9999/sample')
    // axios.get('http://18.180.158.120:9999/sample')
    api.get('/sample')
    .then((res) => {
      console.log(res)
      this.msg = res.data
    })
    .catch((err) => {
      this.msg = "server is not working.."
    })
  }
}

</script>

<style>

</style>
