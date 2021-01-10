<template>
  <div>
    <v-row justify="center">
      <v-col justify="center">
        <v-card
          elevation="5"
          shaped
          color="blue-grey lighten-5"
          class="mx-auto"
          max-width="1200px"
        >
        <div class="pa-6">
          <h3 class="text-center">新規投稿</h3>
          <validationObserver>
            <form @submit.prevent="submitPost()">
              <v-row>
                <v-col cols="12" md="6">
                  <div uk-form-custom id="form_custom">
                    <div class="uk-placeholder uk-text-center">
                      <input type="file" @change="onUpload()" />
                      <div v-if="!post_id" id="preview">
                        <img
                          id="preview_image"
                          v-show="previewImage"
                          :src="previewImage"
                        />
                      </div>
                      <div v-else id="preview">
                        <img id="preview_image" :src="beforeImage" />
                      </div>
                      <div class="camera-choice">
                        <v-icon size="100">mdi-camera</v-icon>
                        <p>画像を選択してください</p>
                      </div>
                    </div>
                  </div>
                  <p id="error_message">{{ message }}</p>
                </v-col>
                <v-col cols="12" md="6">
                  <ValidationProvider
                    mode="lazy"
                    name="カテゴリ"
                    rules="required"
                    v-slot="{ errors }"
                  >
                  <v-select
                    :items="categories"
                    item-text="name"
                    item-value="id"
                    :error-message="errors"
                    v-model="category"
                    label="カテゴリ"
                    placeholder="選択してください"
                  ></v-select>
                  </ValidationProvider>
                </v-col>
              </v-row>
              <v-btn
                block
                large
                elevation="2"
                class="mr-4 mt-4"
                type="submit"
                :disabled="invalid"
                color="blue-grey lighten-2"
              >
              <span>
                <v-icon>mdi-send-outline</v-icon>
                投稿
              </span>
              </v-btn>
            </form>
          </validationObserver>
        </div>

        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import axios from 'axios'
import api from '@/services/api'

import {
  ValidationProvider,
  ValidationObserver,
  extend,
  localize,
} from "vee-validate";

export default {
  data() {
    return {
      images: null,

    }
  },
  methods: {
    onUpload: function() {
      this.images = event.target.files;
      console.log(this.images)
      this.selectedFile()
    },
    selectedFile(event) {
      this.image = event.target.files[0];
      this.createImage(event.target.files[0]);
    },
    createImage(file) {
      console.log("create image run")
      const reader = new FileReader();
      reader.onload = (event) => {
        this.previewImage = event.target.result;
      };
      reader.readAsDataURL(file);
    },
    submitPost() {
      const formData = new FormData();
      console.log(formData)
      formData.append("image", this.images[0]);
      formData.append("user_id", 1)
      console.log(formData)

      api.post('/addImage', formData,{
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
      .then((res) => {
        console.log(res)
        this.$router.replace("/")
      })
      .catch((err) => {
        console.log(err)
      })



    }
  }

}
</script>

<style>

</style>
