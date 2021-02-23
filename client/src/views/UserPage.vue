<template>
  <div>
    userpage
    {{ userData.user_name }}
  </div>
</template>

<script>
import api from '@/services/api'

export default {
  data() {
    return {
      userData: {}
    }
  },
  async mounted() {
    api.post('/fetchUserData', {
      user_id: this.$route.params.user_id
    })
    .then((res) => {
      this.userData = res.data
    })
    .catch((err) => {
      console.log(err)
    })
  }
}
</script>

<style>

</style>
