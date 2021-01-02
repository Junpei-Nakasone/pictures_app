<template>
<div>
  <v-container fluid>
    <v-row align="center">
      <v-col
        class="d-flex"
        cols="12"
        sm="6"
      >
        <v-select
          :items="prefectureCategories"
          item-text="prefecture_name"
          item-value="prefecture_category_cd"
          label="都道府県"
          v-model="items"
          @change="fetchImageByPrefectureCd(items)"
        ></v-select>
      </v-col>
    </v-row>
  </v-container>
  <PostList
    :postType="prefecturePosts"
    />
  </div>
</template>
<script>
import PostList from '@/components/PostList'
import axios from 'axios'
import api from '@/services/api'


export default {
  components: {
    PostList,
  },
    data() {
      return {
        prefectureCategories: [],
        prefecturePosts: []
      }
    },
    async mounted() {
      api.get('/fetchPrefectureCategories')
      .then((res) => {
        this.prefectureCategories = res.data
        console.log(this.prefectureCategories)
      })
    },
    methods: {
      fetchImageByPrefectureCd(value) {
        api.post('/fetchImageByPrefectureCd',{
          'prefecture_category_cd': value
        })
        .then((res) => {
          console.log(res)
          this.prefecturePosts = res.data
        })
      }
    }
  }
</script>

<style>

</style>
