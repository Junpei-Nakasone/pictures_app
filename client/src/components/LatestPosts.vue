<template>
  <div>
    <PostList
    :postType="latestPosts"
    />
  </div>
</template>

<script>
import PostList from '@/components/PostList'
import axios from 'axios'
import api from '@/services/api'

export default {
  components: {
    PostList,
  },
  data() {
    return {
      latestPosts: []
    }
  },
  async mounted() {
    // alert('mounted')
    // axios.get('http://localhost:9999/fetchLatestImages')
    // axios.get('http://18.180.158.120:9999/fetchLatestImages')
    api.get('/fetchLatestImages')
    .then((res) => {
      // alert('fetchlatestImages')
      console.log(res)
      this.latestPosts = res.data
    })
    .catch((res) => {
      console.log(err)
    })
  }
}
</script>

<style>

</style>
