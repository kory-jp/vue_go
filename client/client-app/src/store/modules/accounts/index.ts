import { Account, RootState } from "@/store/types";
import { Module } from "vuex";
import state from "./state";
import actions from "./actions";
import mutations from "./mutations";

export const account: Module<Account, RootState> = {
  namespaced: true,
  state,
  actions,
  mutations,
}