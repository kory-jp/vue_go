<template>
  <div id="app">
    <p>{{ message }}</p>
    <hello-world :msg="msg"></hello-world>
  </div>
</template>

<script lang="ts">
import HelloWorld from './components/HelloWorld.vue';
import { Component, Vue, Prop } from 'vue-property-decorator';
console.log("OK?")
@Component({
  components: {
    HelloWorld
  }
})

export default class App extends Vue {
  message = 'Hello My World';
  msg = "Goodbye World";
}
</script>
