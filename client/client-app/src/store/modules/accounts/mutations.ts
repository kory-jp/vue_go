import { Account } from "@/store/types";
import { MutationTree } from "vuex";

const mutations: MutationTree<Account> = {
  add: (state, account: Account) => {
    state = account
  }
}

export default mutations