export interface RootState {
  varsion: string;
}

export interface Account {
  id: number;
  name: string;
  email: string;
  password: string;
  signIn: boolean;
}