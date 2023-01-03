import axios from "axios";
import { Account } from "@/store/types";

const  headers = {
  'Content-type': 'application/json'
}

const config = {
  method: '',
  url: process.env.VUE_APP_CLIENT_URL,
  headers,
  data: {},
}

export default {
  register: (accountInfo: Account) => {
    config.method = 'POST'
    config.data = accountInfo
    return axios.request(config)
      .then(res => res)
      .catch(err => err)
  }
}