<template>
  <v-app>
    <MyHeader />
    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>

  </v-app>
</template>

<script>
import MyHeader from '@/components/MyHeader.vue'
export default {
  name: 'App',

  components: {
    MyHeader,
},

  data: () => ({
    //
  }),
};
</script>
