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
    api.get('/fetchLatestImages')
    .then((res) => {
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
