import accounts from "@/api/accounts/accounts";
import { Account, RootState } from "@/store/types";
import { ActionTree } from "vuex";

const actions: ActionTree<Account, RootState> = {
  register: async ({commit}, account: Account) => {
    return accounts.register(account)
      .then((res) => {
        commit(res.data)
      })
      // .catch((err => console.log(err)))
      .catch((err) => {
        console.log(err)
      })
  }
}

export default actions