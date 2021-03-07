<template>
  <div>
    <div class="columns is-vcentered">
      <div class="column is-3">
        <figure class="image">
          <img class="is-rounded" :src="userData.icon_image">
        </figure>
      </div>
      <div class="column is-6 is-offset-1">
        <h1 class="title is-2">{{ userData.user_name }}</h1>
        <h3 class="subtitle">{{ userData.note }}</h3>
      </div>
    </div>

    <div class="container has-text-centered">
      <div v-if="userData.PostedPictures.length > 0">
        <h1>{{ userData.user_name }}さんがこれまでに投稿した写真</h1>
      </div>
      <div v-else>
        <h1>{{ userData.user_name }}さんが投稿した写真はまだありません</h1>
      </div>
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
