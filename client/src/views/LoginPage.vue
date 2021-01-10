<template>
  <v-container>
    <h3 class="h3 text-center pt-8">ログイン</h3>
    <v-row justify="center">
      <v-col justify="center">
        <v-card
          elevation="5"
          shaped
          color="blue-grey lighten-5"
          class="mx-auto"
          max-width=500px
        >
          <div class="pa-8">
            <form action="">
              <v-text-field
                v-model="form.username"
                required
                placeholder="username"
                prepend-inner-icon="mdi-account"
              ></v-text-field>
              <v-text-field
              v-model="form.password"
                :append-icon="show1 ? 'mdi-eye' : 'mdi-off'"
                :type="show1 ? 'text': 'password'"
                required
                placeholder="password"
                @click:append="show1 = !show1"
                prepend-inner-icon="mdi-lock"
              ></v-text-field>
              <v-btn block elevation="2" class="mr-4 mt-4" @click="login()">
                Login
              </v-btn>
              </form>
              <v-btn
                block
                color="blue-grey lighten-3"
                elevation="2"
                class="mr-4 mt-4"
                @click="login()"
              >
                デモユーザーでログイン(確認用)
              </v-btn>
              <div class="pa-4 text-center">
                登録していない方
                <router-link id="to_signup" class="router-link" to="/signup">アカウント作成</router-link>
              </div>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  data() {
    return {
      form: {
        username: '',
        password: ''
      },
      show1: false,
    };
  },
  computed: {
    ...mapGetters("auth", {
      username: "username",
      isLoggedIn: "isLoggedIn",
      id: "id",
    })
  },
  methods: {
    login() {
      this.$store.dispatch("auth/login", {
        username: this.form.username,
        password: this.form.password,
      })
      .then(() => {
        if (this.isLoggedIn) {
          this.$store.dispatch("message/setSuccessMessage", {
            message: "ログインしました"
          })
          this.$router.push("/")
        }
        if (!this.isLoggedIn) {
          this.$store.dispatch("message/setErrorMessage", {
          message: "ユーザー名かパスワード名が間違っています"
        })
        }
      })
      .catch((err) => {

      })
    }
  }
}
</script>

<style>

</style>
