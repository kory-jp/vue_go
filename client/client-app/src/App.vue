<template>
  <div id="app">
    <!-- <p>{{ message }}</p>
    <hello-world :msg="msg"></hello-world> -->
    <router-view />
  </div>
</template>

<script lang="ts">
import HelloWorld from './components/HelloWorld.vue';
import { Component, Vue, Prop } from 'vue-property-decorator';
import Top from '@/components/pages/Top.vue'
@Component({
  components: {
    HelloWorld,
    Top
  }
})

export default class App extends Vue {
  // message = 'Hello My World';
  // msg = "Goodbye World";
}
</script>
