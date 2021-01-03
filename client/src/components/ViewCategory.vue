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
          :items="viewCategories"
          item-text="view_name"
          item-value="view_category_cd"
          label="景色"
          v-model="items"
          @change="fetchImageByPrefectureCategoryCd(items)"
        ></v-select>
      </v-col>
    </v-row>
  </v-container>
  <PostList
    :postType="viewPosts"
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
        viewCategories: [],
        viewPosts: []
      }
    },
    async mounted() {
      api.get('/fetchViewCategories')
      .then((res) => {
        this.viewCategories = res.data
        console.log(this.viewCategories)
      })
    },
    methods: {
      fetchImageByPrefectureCategoryCd(value) {
        alert(value)
        api.post('/fetchImageByViewCategoryCd',{
          'view_category_cd': value
        })
        .then((res) => {
          console.log(res)
          this.viewPosts = res.data
        })
      }
    }
  }
</script>

<style>

</style>
