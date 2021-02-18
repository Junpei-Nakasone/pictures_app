<template>
<b-navbar>
  <template #brand>
    <b-navbar-item  tag="router-link" :to="{ path: '/' }">
      View Pictures
    </b-navbar-item>
  </template>

  <template #start>
    <b-navbar-item>
      <router-link to="/view-category">
        景色別
      </router-link>
    </b-navbar-item>
    <b-navbar-item>
      <router-link to="/prefecture_category">
        都道府県別
      </router-link>
    </b-navbar-item>
  </template>

  <template #end>
    <b-navbar-item>
      <div v-if="!isLoggedIn" class="buttons">
        <b-button
          tag="router-link"
          to="/signup"
          type="is-info is-light">
          新規登録
        </b-button>
        <b-button
          tag="router-link"
          to="/login"
          type="is-info is-light">
          ログイン
        </b-button>
      </div>
      <div v-if="isLoggedIn" class="buttons">
        <b-navbar-dropdown label="info">
          <b-navbar-item href="#">
            マイメニュー
          </b-navbar-item>
          <b-navbar-item @click="logout">
            ログアウト
          </b-navbar-item>
        </b-navbar-dropdown>
        <b-button
          tag="router-link"
          to="/newPost"
          type="is-info is-light">
          写真投稿
        </b-button>
      </div>
    </b-navbar-item>
  </template>

</b-navbar>
  <!-- <v-container>
    <div v-if="isLoggedIn">
    <v-toolbar flat class="mainHeader">
      <router-link to="/" id="title">
        <v-toolbar-title>View Pictures</v-toolbar-title>
      </router-link>
      <v-spacer></v-spacer>
      {{username}}&nbsp;様
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
        <v-toolbar-title>View Pictures</v-toolbar-title>
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
              <v-list-item-title>
                <v-btn
                  color="white"
                depressed
                  class="ma-1"
                  to="signup"
                >新規登録
                </v-btn>
                </v-list-item-title>
            </v-list-item>
            <v-list-item>
              <v-list-item-title>
                <v-btn
                  color="white"
                  depressed
                  class="ma-1"
                  to="login"
                >
                  ログイン
                </v-btn>
                </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
    </v-toolbar>
    </div>
  </v-container> -->
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
