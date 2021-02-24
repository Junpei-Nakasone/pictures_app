<template>
  <div>
    <v-container>
      <h3 class="h3 text-center">新規ユーザー登録</h3>
      <v-row justify="center">
        <v-col justify="center">
          <v-card
            elevation="5"
            shaped
            color="blue-grey lighten-5"
            class="mx-auto"
            max-width="500px"
          >
            <div class="pa-8">
              <ValidationObserver>
                <form @submit.prevent="submitUser()">
                  <ValidationProvider
                    mode="eager"
                    name="ユーザー名"
                    rules="required|max:10"

                  >
                  <v-text-field
                    v-model="username"
                    required
                    placeholder="ユーザー名"
                    prepend-inner-icon="mdi-account"
                  ></v-text-field>
                  </ValidationProvider>
                  <ValidationProvider
                    mode="lazy"
                    name="入力内容"
                    rules="required|email"
                  >
                    <v-text-field
                      v-model="email_address"
                      :error-messages="errors"
                      required
                      placeholder="メールアドレス"
                      prepend-inner-icon="mdi-email"
                    ></v-text-field>
                  </ValidationProvider>
                  <ValidationProvider
                    mode="lazy"
                    name="パスワード"
                    rules="required|min:8|password"

                  >
                    <v-text-field
                      v-model="password"
                      :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
                      :type="show1 ? 'text' : 'password'"
                      counter
                      :error-messages="errors"
                      required
                      placeholder="パスワード"
                      @click:append="show1 = !show1"
                      prepend-inner-icon="mdi-lock"
                    ></v-text-field>
                  </ValidationProvider>
                  <ValidationProvider
                    mode="aggressive"
                    name="パスワード"
                    rules="required|confirmed:password1"
                  >
                    <v-text-field
                      v-model="password_confirm"
                      :append-icon="show2 ? 'mdi-eye' : 'mdi-eye-off'"
                      :type="show2 ? 'text' : 'password'"
                      counter
                      :error-messages="errors"
                      required
                      placeholder="パスワード(確認)"
                      @click:append="show2 = !show2"
                      prepend-inner-icon="mdi-lock"
                    ></v-text-field>
                  </ValidationProvider>
                  <ValidationProvider
                    mode="lazy"
                    name="備考"
                    rules="required|note"
                  >
                    <v-textarea
                      v-model="note"
                      :error-messages="errors"
                      required
                      placeholder="自己紹介"
                      prepend-inner-icon="mdi-person"
                    ></v-textarea>
                  </ValidationProvider>
                  <v-btn
                    block
                    elevation="2"
                    class="mr-4 mt-4"
                    type="submit"
                    :disabled="invalid"
                  >
                    送信
                  </v-btn>
                </form>
              </ValidationObserver>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import axios from 'axios'
import api from '@/services/api'
import { mapGetters } from "vuex"
import mixin from '@/mixin/mixin.js';

import {
  ValidationProvider,
  ValidationObserver,
  // extend,
} from "vee-validate";
export default {
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  mixins: [
    mixin
  ],
  data() {
    return {
      message: "",
      username: "",
      email_address: "",
      password: "",
      password_confirm: "",
      note: "",
      isLoading: false,
      show1: false,
      show2: false,
    };
  },
  computed: {
    ...mapGetters("auth", {
      isLoggedIn: "isLoggedIn",
      id: "id",
    })
  },
  methods: {
    submitUser() {
      api.post('/addNewUser', {
        "user_name": this.username,
        "password": this.password,
	      "email_address": this.email_address,
        "note": this.note,
      })
      .then((response) => {
        this.loginAfterSignup();
      })
      .catch((error) => {
        console.log("エラー", error)
      })
    },
    loginAfterSignup() {
      this.$store.dispatch("auth/login", {
        username: this.username,
        password: this.password
      })
      .then(() => {
        if (this.isLoggedIn) {
          this.showSignupSuccessMessage()
          this.$router.push("/")
        }
      })
      .catch((err) => {
        this.showSignupFailMessage()
      })
    }
  }
}
</script>

<style>

</style>
