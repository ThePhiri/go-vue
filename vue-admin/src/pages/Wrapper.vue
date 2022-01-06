<template>
      <Nav />
  <div class="container-fluid">
    <div class="row">
      <Menu />

      <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script>
import Nav from "@/components/Nav";
import Menu from "@/components/Menu";
import { onMounted } from '@vue/runtime-core';
import axios from 'axios';
import {useRouter} from 'vue-router';


export default {
  components: {
    Nav,
    Menu,
  },
  setup(){
      const router = useRouter();
      onMounted(async ()  => {
          try {
               const {data} = await axios.get('user');
          } catch (error) {
              await router.push('/login');
          }
      
      });
  }
};
</script>

<style scoped>

</style>