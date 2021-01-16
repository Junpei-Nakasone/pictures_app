<template>
  <v-container>
    <div v-if="isLoggedIn">
    <v-toolbar flat class="mainHeader">
      <router-link to="/" id="title">
        <v-toolbar-title>Pictures</v-toolbar-title>
      </router-link>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-xs-only">
        <v-btn
          large
            class="ma-1"
            elevation="2"
            to="newPost"
        >投稿する</v-btn>
        <v-btn
        @click="logout"
        large
        class="ma-1"
        elevation="2"
      >ログアウト
      </v-btn>
      </v-toolbar-items>
      <div class="hidden-sm-and-up">
        <v-menu offset-y>
          <template v-slot:activator="{ on }">
            <v-app-bar-nav-icon v-on="on"></v-app-bar-nav-icon>
          </template>
          <v-list class="responsiveMenu">
            <v-list-item>
              <v-list-item-title>
                <v-btn
                color="white"
                  depressed
                  class="ma-1"
                  to="newPost"
                >
                  投稿する
                </v-btn>
                </v-list-item-title>
            </v-list-item>
            <v-list-item>
              <v-list-item-title>
                <v-btn
                color="white"
                depressed
                @click="logout"
                  class="ma-1"
                  to="newPost"
                >
                  ログアウト
                </v-btn>
                </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
    </v-toolbar>
    </div>

    <div v-if="!isLoggedIn">
    <v-toolbar flat class="mainHeader">
      <router-link to="/" id="title">
        <v-toolbar-title>Pictures</v-toolbar-title>
      </router-link>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-xs-only">
        <v-btn
            to="/signup"
            large
            class="ma-1"
            elevation="2"
          >新規登録
        </v-btn>

        <v-btn
            to="/login"
            large
            class="ma-1"
            elevation="2"
        >ログイン
        </v-btn>

      </v-toolbar-items>
      <div class="hidden-sm-and-up">
        <v-menu offset-y>
          <template v-slot:activator="{ on }">
            <v-app-bar-nav-icon v-on="on"></v-app-bar-nav-icon>
          </template>
          <v-list class="responsiveMenu">
            <v-list-item>
              <v-list-item-title><router-link to="/signup">新規登録</router-link></v-list-item-title>
            </v-list-item>
            <v-list-item>
              <v-list-item-title><router-link to="/login">ログイン</router-link></v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
    </v-toolbar>
    </div>
  </v-container>
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

<style scoped>
#title {
  color: black;
  text-decoration: none;
}
.no-decoration {
  text-decoration: none;
}

a {
  text-decoration: none;
}
</style>
