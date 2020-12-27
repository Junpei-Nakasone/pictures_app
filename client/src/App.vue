<template>
  <v-app>
    <MyHeader />
    <v-main>
    <GlobalMessage />
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>

  </v-app>
</template>

<script>
import MyHeader from '@/components/MyHeader.vue'
import GlobalMessage from '@/components/GlobalMessage.vue'

export default {
  name: 'App',

  components: {
    MyHeader,
    GlobalMessage,
},

  data: () => ({
    //
  }),
};
</script>
