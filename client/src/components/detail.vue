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
      <div class="column">
        <b-button
          tag="router-link"
          to="/"
          type="is-info"
        >戻る</b-button>
      </div>
    </div>
    <b-image :src="pictureData.image_url" />
  </section>
</template>

<script>
import api from '@/services/api'

export default {
  data() {
        return {
          pictureData: {}
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
    }
}
</script>

<style>

</style>
