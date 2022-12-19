<template>
  <div class="bg-gray-200 box-border h-32 w-32">
    <h2 class="font-bold">ログイン!</h2>
    <div>
      <input type="text">
      <input type="text">
    </div>
    <router-link to="/">トップページへ</router-link>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
  @Component
  export default class Login extends Vue{}
</script>