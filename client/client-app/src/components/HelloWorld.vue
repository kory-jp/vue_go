<template>
  <div>
    <h1>{{ msg }}</h1>
    <p>"All right?"</p>
  </div>
</template>


<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

console.log("child OK?")
@Component({
  name: 'HelloWorld'
})
export default class App extends Vue {
  @Prop({type: String, required: true}) msg!: string
}
</script>