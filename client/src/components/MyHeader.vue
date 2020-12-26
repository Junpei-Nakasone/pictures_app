<template>
  <v-container>
  <v-app-bar app color="teal lighten-3">
    <v-app-bar-nav-icon class="d-none"></v-app-bar-nav-icon>
    <v-toolbar-title class="pl-0 pl-sm-5">
      <router-link to="/" id="title">
        <v-btn icon class="mb-2" depressed disabled>
          <!-- <v-img
            src="../assets/logo.png"
            max-width="35px"
            max-height="35px"
          /> -->
        </v-btn>
        Pictures
      </router-link>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <HeaderSide />
  </v-app-bar>
  </v-container>
</template>

<script>
import HeaderSide from '@/components/HeaderSide.vue'
export default {
  components: {
    HeaderSide,
  }
}
</script>

<style scoped>
@import "../assets/common.css";

#title {
  text-decoration: none;
  font-size: 40px;
  font-family: "Economica";
  color: black;
}

</style>
