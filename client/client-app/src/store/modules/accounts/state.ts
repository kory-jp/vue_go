export class Account {
  signIn = false;
  id = "";
  name = "hello";
  email = "";
  password = "";
}

export interface State {
  account: Account;
}

const state: State = {
  account: new Account()
}

export default state;