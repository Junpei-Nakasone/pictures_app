<template>
  <div>
    <v-container>
      <v-row>
        <v-col
          class="pa-1 pa-sm-2"
          v-for="(post,key) in postType"
          :key="key"
          cols="6"
          lg="4"
          md="4"
          sm="6"
          xs="3"
        >
        {{post.post_id}}
        <v-img :src="post.image_url" max-height="230"></v-img>

        </v-col>
      </v-row>
    </v-container>
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
