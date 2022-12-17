<template>
  <div>
    <p>TOP</p>
    <router-link to="/login">ログインページへ</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
@Component
export default class Top extends Vue {
  
}
</script>