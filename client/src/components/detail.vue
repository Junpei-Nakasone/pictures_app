<template>
  <section>
    <div class="columns is-mobile">
      <div class="column">
        <figure class="image is-48x48">
          <router-link :to="{name: 'userpage', params: {user_id: pictureData.user_id}}">
            <b-tooltip :label="pictureData.user_name">
              <img class="is-rounded" :src="pictureData.icon_image">
            </b-tooltip>
          </router-link>
        </figure>
        <p>{{ pictureData.image_note }}</p>
      </div>
    </div>
    <div class="columns">
      <div class="column is-1">
        <b-button
          tag="router-link"
          to="/"
          type="is-info"
        >戻る</b-button>
      </div>
      <div v-if="pictureData.user_id === id" class="column">
        <b-button
          @click="deletePicture"
          type="is-danger"
        >削除</b-button>
      </div>
    </div>
    <b-image :src="pictureData.image_url" />
  </section>
</template>

<script>
import api from '@/services/api'
import mixin from '@/mixin/mixin.js'

export default {
  data() {
      return {
        pictureData: {},
        userData: {}
      }
  },
  mixins: [
    mixin
  ],
  computed: {
    id: function() {
      return this.$store.getters["auth/id"]
    }
  },
  async mounted() {
      api.post('/fetchImageByPictureID', {
        picture_id: this.$route.params.id
      })
      .then((res) => {
        this.pictureData = res.data
      })
      .catch((err) => {
        console.log(err)
      })
    },
  methods: {
    deletePicture() {
      api.delete('/deletePicture', {
        data: {
          picture_id: this.pictureData.picture_id
        }
      })
      .then((res) => {
        this.showDeleteMessage()
        this.$router.push('/')
      })
      .catch((err) => {
        console.log(err)
      })
    }
  }
}
</script>

<style>

</style>
