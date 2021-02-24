<template>
  <div>
    <div class="columns is-vcentered">
      <div class="column is-5">
        <figure class="image is-4by3">
          <img :src="userData.icon_image">
        </figure>
      </div>
      <div class="column is-6 is-offset-1">
        <h1 class="title is-2">{{ userData.user_name }}</h1>
        <h3 class="subtitle">{{ userData.note }}</h3>
      </div>
    </div>

    <div class="container has-text-centered">
      <h1>{{ userData.user_name }}さんがこれまでに投稿した写真</h1>
    </div>
    <div>
      <PostList
        :postType="userData.PostedPictures"
      />
    </div>
  </div>
</template>

<script>
import PostList from '@/components/PostList'
import api from '@/services/api'

export default {
  components: {
    PostList,
  },
  data() {
    return {
      userData: {}
    }
  },
  async mounted() {
    api.post('/fetchUserData', {
      user_id: Number(this.$route.params.user_id)
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
