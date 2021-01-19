<template>
  <v-card class="mx-auto">
    <v-list>
      <v-list-item-group>
        <!-- ログイン時のメニュー -->
        <router-link class="router-link d-sm-none" to="/newPost" v-if="isLoggedIn">
        <v-list-item>
          <v-list-item-icon>
            <v-icon>mdi-pencil-outline</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>
              投稿する
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        </router-link>
        <!-- ログイン前のメニュー -->
        <router-link class="router-link d-sm-none" to="/newPost" v-if="!isLoggedIn">
        <v-list-item>
          <v-list-item-icon>
            <v-icon>mdi-pencil-outline</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>
              投稿する
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        </router-link>
      </v-list-item-group>
    </v-list>
  </v-card>
</template>

<script>
export default {

}
</script>

<style>

</style>
