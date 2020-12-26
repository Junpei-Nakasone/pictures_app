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
                Guest Login
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
  methods: {
    login(username="GuestUser", password="test") {
      alert('login run')
      this.form.username = username;
      this.form.password = password;
      console.log(this.form.username)

      this.$store.dispatch("auth/login", {
        username: username,
        password: password,
      })
      this.$router.replace("/")
      // .then(() => {
      //   alert('login success')
      // })
      // .catch(() => {
      //   alert('err')
      // })
    }
  }
}
</script>

<style>

</style>
