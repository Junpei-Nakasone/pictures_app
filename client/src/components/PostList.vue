<template>
  <div>
    <div class="columns is-gapless is-multiline">
      <div class="column is-4" v-for="post in postType" :key="post.picture_id">
        <router-link :to="{name: 'detail', params: {id: post.picture_id}}">
          <img :src="post.image_url">
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: [
    "postType"
  ],

}
</script>

<style>

</style>
