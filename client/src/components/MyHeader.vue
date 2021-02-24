<template>
<b-navbar>
  <template #brand>
    <b-navbar-item  tag="router-link" :to="{ path: '/' }">
      View Pictures
    </b-navbar-item>
  </template>

  <template #start>
    <b-navbar-item>
      <router-link to="/view_category">
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
          <router-link :to="{name: 'userpage', params: {user_id: id}}">
            <b-navbar-item>
            {{username}}のユーザーページ
          </b-navbar-item>
          </router-link>
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
</template>

<script>
import mixin from '@/mixin/mixin.js';

export default {
    mixins: [
      mixin
    ],
    computed: {
    isLoggedIn : function () {
      return this.$store.getters["auth/isLoggedIn"]
    },
    username : function () {
      return this.$store.getters["auth/username"]
    },
    id : function () {
      return this.$store.getters["auth/id"]
    }
  },
  methods: {
    logout: function() {
      sessionStorage.clear();
      this.$store.dispatch("auth/logout")
      this.showLogoutMessage()
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
