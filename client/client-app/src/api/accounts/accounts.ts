import axios from "axios";
import { Account } from "@/store/types";

const  headers = {
  'Content-type': 'application/json'
}

const config = {
  method: '',
  url: '',
  headers,
  withCredentials: true,
  data: {},
}

export default {
  register: (accountInfo: Account) => {
    config.method = 'POST'
    config.url = process.env.VUE_APP_CLIENT_URL + "register"
    config.data = accountInfo
    return axios.request(config)
      .then(res => res)
      .catch(err => err)
  }
}