<template>
  <div>
    <v-container>
      <h3 class="text-center">新規投稿</h3>
      <form @submit.prevent="submitPost()">
      <v-row>
        <v-col sm="5" md="6">
          <div class="FieldField__wrapper">
    <img v-bind:src="imagePreview" class="preview-image" v-on:click="openUpload">
    <div class="form-group">
      <input
        name="image"
        type="file"
        id="file-field"
        v-on:change="updatePreview"
        style="display:none;"
        >
    </div>
  </div>
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
      images: null,
      imagePreview: 'https://cdn.icon-icons.com/icons2/930/PNG/512/camera_icon-icons.com_72364.png',
      prefectureCategories: [],
      viewCategories: [],
      postData: {}
    }
  },
  computed: {
    username : function () {
      this.postData.id = this.$store.getters["auth/id"]
    }
  },
  async mounted() {
    api.get('/fetchPrefectureCategories')
      .then((res) => {
        this.prefectureCategories = res.data
        console.log(this.prefectureCategories)
      });
      api.get('/fetchViewCategories')
      .then((res) => {
        this.viewCategories = res.data
        console.log(this.viewCategories)
      });
  },
  methods: {
    openUpload () {
        document.getElementById('file-field').click()
      },
    updatePreview (e) {
        this.images = e.target.files
        console.log('e', e)
        var reader, files = e.target.files
        if (files.length === 0) {
          console.log('Empty')
        }
        reader = new FileReader();
        reader.onload = (e) => {
          this.imagePreview = e.target.result
        }
        reader.readAsDataURL(files[0])
      },
    submitPost() {
      const formData = new FormData();
      console.log(formData)
      formData.append("image", this.images[0]);
      formData.append("view_category_cd", this.postData.viewCategoryCd)
      formData.append("prefecture_category_cd", this.postData.prefectureCategoryCd)
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
    },
    updateViewData(value) {
      this.postData.viewCategoryCd = value
    },
    updatePrefectureData(value) {
      this.postData.prefectureCategoryCd = value
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
