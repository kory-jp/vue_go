<template>
  <div>
    <p class="text-sky-400 font-bold">ログイン</p>
    <router-link to="/">トップページへ</router-link>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
  @Component
  export default class Login extends Vue{}
</script>