<template>
  <div>
    <v-container>
      <h3 class="text-center">新規投稿</h3>
      <form @submit.prevent="submitPost()">
      <v-row>
        <v-col sm="5" md="6">
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
        </v-col>
        <v-col sm="5" md="6" >
          <v-select
            :items="viewCategories"
            item-text="view_name"
            item-value="view_category_cd"
            label="景色を選択してください"
            v-model="items"
            @change="updateViewData(items)"
          ></v-select>
          <v-select
            :items="prefectureCategories"
            item-text="prefecture_name"
            item-value="prefecture_category_cd"
            label="都道府県を選択してください"
            v-model="items"
            @change="updatePrefectureData(items)"
          ></v-select>
          <v-textarea
              v-model="image_note"
              :error-messages="errors"
              required
              placeholder="写真の説明"
              prepend-inner-icon="mdi-person"
              @change="updateImageNote"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-btn
        block
        class="mr-4 mt-14"
        x-large
        elevation="3"
        type=submit
      >
        <span>
          <v-icon>mdi-send-outline</v-icon>
          投稿
          </span>
      </v-btn>
      <div v-if="showProgress">
        <b-progress></b-progress>
      </div>
      </form>
    </v-container>
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
      prefectureCategories: [],
      viewCategories: [],
      postData: {},
      file: null,
      showProgress: false
    }
  },
  computed: {
    userID() {
      this.postData.id = this.$store.getters["auth/id"]
    },
    image() {
      return this.file ? window.URL.createObjectURL(this.file): ""
    }
  },
  async mounted() {
    api.get('/fetchPrefectureCategories')
      .then((res) => {
        this.prefectureCategories = res.data
      });
      api.get('/fetchViewCategories')
      .then((res) => {
        this.viewCategories = res.data
      });
  },
  methods: {
    submitPost() {
      this.showProgress = true;
      const formData = new FormData();
      formData.append("image", this.file);
      formData.append("view_category_cd", this.postData.viewCategoryCd)
      formData.append("prefecture_category_cd", this.postData.prefectureCategoryCd)
      formData.append("image_note", this.postData.imageNote)
      formData.append("user_id", this.postData.id)

      api.post('/addImage', formData,{
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
      .then((res) => {
        this.$router.replace("/")
      })
      .catch((err) => {
        console.log(err)
      })
    },
    updateViewData(value) {
      this.postData.viewCategoryCd = value
    },
    updatePrefectureData(value) {
      this.postData.prefectureCategoryCd = value
    },
    updateImageNote(value) {
      this.postData.imageNote = value
    }
}
}
</script>

<style>
.preview-image {
    width: 327px;
    height: 184px;
  }
</style>
