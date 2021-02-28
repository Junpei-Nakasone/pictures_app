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
                v-model="form.emailAddress"
                required
                placeholder="email_address"
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
                ログイン
              </v-btn>
              </form>
              <v-btn
                block
                color="blue-grey lighten-3"
                elevation="2"
                class="mr-4 mt-4"
                @click="demoLogin()"
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
import mixin from '@/mixin/mixin.js';

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
  mixins: [
    mixin
  ],
  computed: {
    ...mapGetters("auth", {
      username: "username",
      isLoggedIn: "isLoggedIn",
      id: "id",
    }),

  },
  methods: {
    login() {
      this.$store.dispatch("auth/login", {
        emailAddress: this.form.emailAddress,
        password: this.form.password,
      })
      .then(() => {
        if (this.isLoggedIn) {
          this.showLoginSeccessMessage()
          this.$router.push("/")
        }
        if (!this.isLoggedIn) {
          this.showLoginFailMessage()
        }
      })
      .catch((err) => {
      })
    },
    demoLogin()  {

      this.$store.dispatch("auth/login", {
        emailAddress: "user1@email.com",
        password: "demo_password",
      })
      .then(() => {
        if (this.isLoggedIn) {
          this.showLoginSeccessMessage()
          this.$router.push("/")
        }
        if (!this.isLoggedIn) {
          this.showLoginFailMessage()
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
