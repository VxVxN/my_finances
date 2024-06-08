type AuthForm = {
  username: string;
  password: string;
};

export type SignUpData = {} & AuthForm;

export type SignInData = {} & AuthForm;