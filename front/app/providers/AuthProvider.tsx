import { createContext, ReactNode, useContext } from "react";


type AuthContextType = {
    isAuth: boolean;
}

const initial: AuthContextType = {
    isAuth: true,
};

const AuthContext = createContext<AuthContextType>(initial);

type AuthProviderProps = {
    children: ReactNode;
}

export function AuthProvider({children}: AuthProviderProps) {

    const isAuth = true;
    const loading = true;


    return (
        <AuthContext.Provider value={{isAuth  }}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    return useContext(AuthContext);
}