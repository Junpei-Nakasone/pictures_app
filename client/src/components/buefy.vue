<template>
    <div class="posts columns is-multiline is-4">
          <div class="column is-4"
              v-for="post in latestPosts"
              :key="post.picture_id">
              <div class="card">
                  <div class="card-image">
                      <figure class="image">
                        <router-link :to="{name: 'detail', params: {id: post.picture_id }}">
                          <img :src="post.image_url"
                              alt="Placeholder image">
                        </router-link>
                      </figure>
                  </div>
                  <div class="card-content">
                      <div class="media">
                          <div class="media-left">
                              <figure class="image is-48x48">
                                  <img :src="post.image_url">
                              </figure>
                          </div>
                          <div class="media-content">

                              <p class="subtitle is-6"></p>
                          </div>
                      </div>
                  </div>
              </div>
          </div>
      </div>
</template>

<script>
import api from '@/services/api'

export default {
    data() {
        return {
          latestPosts: []
        }
    },
    async mounted() {
      api.get('/fetchLatestImages')
      .then((res) => {
        this.latestPosts = res.data
      })
      .catch((err) => {
        console.log(err)
      })
    }
}
</script>
