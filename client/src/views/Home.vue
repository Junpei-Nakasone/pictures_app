<template>
  <div>
    <v-tabs color="blue-grey lighten-2" centered show-arrows>
      <v-tab to="/">
        <v-icon left>mdi-heart</v-icon>
          <span>新着投稿</span>
      </v-tab>
      <v-tab to="/view_category">
        <v-icon left>mdi-shape</v-icon>
          <span>景色カテゴリー別</span>
      </v-tab>
      <v-tab to="/prefecture_category">
        <v-icon left>mdi-shape</v-icon>
          <span>都道府県別</span>
      </v-tab>
      <v-tab to="/newPost">
        <v-icon left>mdi-plus</v-icon>
        <span>新規投稿</span>
      </v-tab>
      <v-tab to="/sample">
        <v-icon left>mdi-football</v-icon>
        <span>Sample</span>
      </v-tab>
      <!-- <v-tab to="/map">
        <v-icon left>mdi-map-marker</v-icon>
          <span>MAP</span>
      </v-tab> -->
    </v-tabs>
    <div class="content">
      <transition appear>
        <router-view />
      </transition>
    </div>
  </div>

</template>

<script>
// @ is an alias to /src

export default {
  computed: {
  },
  methods: {
  },
  name: 'Home',
  components: {
  }
}
</script>
