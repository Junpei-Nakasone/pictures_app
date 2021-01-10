<template>
  <div v-if="isLoggedIn">
    <div class="d-none d-sm-flex">
      <v-layout>
        <span style="font-size:20px">
          {{username}}
        </span>
      </v-layout>

      <router-link class="router-link" to="/newPost">
        <v-btn
          rounded
          depressed
          elevation="3"
          color="blue-grey lighten-4"
          large
          class="ma-1"
        ><v-icon>mdi-camera</v-icon>投稿する</v-btn>
      </router-link>
      <v-btn
      rounded
        @click="logout"
        depressed
        elevation="3"
        color="blue-grey lighten-4"
        large
        class="ma-1"
      ><v-icon>mdi-minus</v-icon>
        ログアウト
      </v-btn>
    </div>
  </div>
<!-- ログインしていない時 -->
  <div v-else>
    <div class="d-none d-sm-flex">
      <router-link class="router-link" id="post" to="/signup">
        <v-btn
        rounded
        depressed
        elevation="3"
        large
         color="blue-grey lighten-4" class="ma-1">
          <v-icon>mdi-account-plus</v-icon>
          新規登録
        </v-btn>
      </router-link>
      <router-link class="router-link" id="post" to="/login">
        <v-btn
        rounded
        depressed
        elevation="3"
        large
         color="blue-grey lighten-4" class="ma-1">
          <v-icon>mdi-login</v-icon>
          ログイン
        </v-btn>
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    isLoggedIn : function () {
      return this.$store.getters["auth/isLoggedIn"]
    },
    username : function () {
      return this.$store.getters["auth/username"]
    }
  },
  methods: {
    logout: function() {
      sessionStorage.clear();
      this.$store.dispatch("auth/logout")
      this.$store.dispatch('message/setSuccessMessage', {
        message: 'ログアウトしました',
      });
      this.$router.replace("/login")
    }
  }
}
</script>

<style>

</style>
