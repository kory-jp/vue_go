export class User {
  signIn = false;
  id = "";
  name = "hello";
  email = "";
  password = "";
}

export interface State {
  user: User;
}

const state: State = {
  user: new User()
}

export default state;