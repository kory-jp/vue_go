<template>
  <div class="h-screen w-screen flex justify-center items-center">
    <div class="bg-gray-200 box-border h-96 w-1/5 rounded-lg drop-shadow-lg">
      <div class="p-12">
        <h2 class="font-bold mb-8 text-xl text-center">ログイン</h2>
        <div class="mb-8 space-y-4">
          <input v-model="email" type="text" placeholder="mail address" class="shadow border rounded w-full  py-2 px-3 focus:outline-none forcus:shadow-outline">
          <input v-model="password" type="text" placeholder="password" class="shadow border rounded w-full  py-2 px-3 focus:outline-none forcus:shadow-outline">
        </div>
        <div class="mb-8">
          <button @click="login" class="p-3 bg-sky-500 hover:bg-sky-500/75 rounded-xl w-full shadow-lg">Login</button>
        </div>
        <router-link to="/">トップページへ</router-link>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Account } from "@/store/types";
import { Component, Prop, Vue } from "vue-property-decorator";
  @Component
  export default class Login extends Vue{
    email = "";
    password = "";
    login() {
      const account: Account = {
        id: 0,
        name: "",
        email: this.email,
        password: this.password,
        signIn: false,
      }
      this.$store.dispatch('account/login', account)
    }
  }
</script>

