<template>
  <div>
    <v-card :ripple="false" color="blue-grey darken-1" elevation="5" shaped tabindex="-1" uk-slider="finite: false">
      <div class="uk-position-relative" id="category?_card_content">
        <ul
          id="category_choice"
          class="uk-slider-items uk-grid uk-grid-column-small uk-child-width-1-2 uk-child-width-1-4-@s uk-child-width-1-6"
        ></ul>
      </div>
    </v-card>
  </div>
</template>

<script>
import PostList from '@/components/PostList';
import api from '@/services/api';

export default {
  components: {
    PostList,
  },

  data() {
    return {
      categories: [
        {
          viewCategoryCd: "1",
          viewCategoryName: "空"
        },
        {
          viewCategoryCd: "2",
          viewCategoryName: "海"
        },
      ]
    }
  }
}
</script>

<style>

</style>
