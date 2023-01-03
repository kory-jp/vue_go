<template>
  <div class="h-screen w-screen flex justify-center items-center">
    <div class="bg-gray-200 box-border h-auto w-1/5 rounded-lg drop-shadow-lg">
      <div class="p-12">
        <h2 class="font-bold mb-8 text-xl text-center">新規登録</h2>
        <div class="mb-8 space-y-4">
          <input v-model="name" type="text" placeholder="name" class="shadow border rounded w-full  py-2 px-3 focus:outline-none forcus:shadow-outline">
          <input v-model="email" type="text" placeholder="mail address" class="shadow border rounded w-full  py-2 px-3 focus:outline-none forcus:shadow-outline">
          <input v-model="password" type="password" placeholder="password" class="shadow border rounded w-full  py-2 px-3 focus:outline-none forcus:shadow-outline">
        </div>
        <div
          class="mb-8"
        >
          <button @click="registerAccount" class="p-3 bg-sky-500 hover:bg-sky-500/75 rounded-xl w-full shadow-lg">Singup</button>
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
    name = '';
    email = '';
    password = '';
    registerAccount() {
      const account: Account = {
        id: "",
        name: this.name,
        email: this.email,
        password: this.password,
        signIn: false,
      }
      console.log("hello", account);
      this.$store.dispatch('account/register', account)
    }
  }
</script>