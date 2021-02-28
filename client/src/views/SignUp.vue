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
                    name="パスワード"
                    rules="required|min:8|password"

                  >
                    <v-text-field
                      v-model="password"
                      :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
                      :type="show1 ? 'text' : 'password'"
                      counter
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
                      required
                      placeholder="パスワード(確認)"
                      @click:append="show2 = !show2"
                      prepend-inner-icon="mdi-lock"
                    ></v-text-field>
                  </ValidationProvider>
                  <ValidationProvider
                    mode="lazy"
                    name="入力内容"
                    rules="required|email"
                  >
                    <v-text-field
                      v-model="emailAddress"
                      required
                      email
                      placeholder="メールアドレス"
                      prepend-inner-icon="mdi-email"
                    ></v-text-field>
                  </ValidationProvider>
                  <ValidationProvider
                    mode="lazy"
                    name="備考"
                    rules="required|note"
                  >
                    <v-textarea
                      v-model="note"
                      required
                      placeholder="自己紹介"
                      prepend-inner-icon="mdi-person"
                    ></v-textarea>
                  </ValidationProvider>

                  <b-field class="file">
                  <b-upload v-model="file" required>
                    <a class="button is-primary">
                      <b-icon icon="camera"></b-icon>
                      <span v-if="file == null">
                        写真を選択してください
                      </span>
                    </a>
                  </b-upload>
                  <span class="file-name" v-if="file">
                    {{ file.name }}
                  </span>
                </b-field>
                <figure>
                  <img :src="image" />
                </figure>
                  <v-btn
                    block
                    elevation="2"
                    class="mr-4 mt-4"
                    type="submit"
                  >
                    送信
                  </v-btn>
                  <div v-if="showProgress">
                  <b-progress></b-progress>
                </div>
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
      emailAddress: "",
      password: "",
      passwordConfirm: "",
      note: "",
      isLoading: false,
      show1: false,
      show2: false,
      file: null,
      showProgress: false
    };
  },
  computed: {
    ...mapGetters("auth", {
      isLoggedIn: "isLoggedIn",
      id: "id",
    }),
    image() {
      return this.file ? window.URL.createObjectURL(this.file): ""
    }
  },
  methods: {
    submitUser() {
      this.showProgress = true;
      const formData = new FormData();
      formData.append("user_name", this.username)
      formData.append("password", this.password)
      formData.append("note", this.note)
      formData.append("email_address", this.emailAddress)
      formData.append("icon_image", this.file)

      api.post('/addNewUser', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
      .then((response) => {
        this.loginAfterSignup();
      })
      .catch((error) => {
        this.showProgress = false;
        this.showSignupFailMessage()
        this.$router.push("/signup")
      })
    },
    loginAfterSignup() {
      this.$store.dispatch("auth/login", {
        emailAddress: this.emailAddress,
        password: this.password
      })
      .then(() => {
        if (this.isLoggedIn) {
          this.showProgress = false;
          this.showSignupSuccessMessage()
          this.$router.push("/")
        }
      })
      .catch((err) => {
        this.showProgress = false;
        this.showSignupFailMessage()
        this.$router.push("/signup")
      })
    }
  }
}
</script>

<style>

</style>
