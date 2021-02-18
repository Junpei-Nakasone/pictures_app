<template>
  <section>
    <b-image :src="pictureData.image_url" ratio="16by9" />
    <b-button
      tag="router-link"
      to="/"
      type="is-info"
    >戻る</b-button>
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
