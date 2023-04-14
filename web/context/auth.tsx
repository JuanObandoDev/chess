import { ReactNode, createContext, useContext, useEffect } from "react";
import useSWR from "swr";

import { APIError } from "@/types/Error";
import { User, UserSchema } from "@/types/User";
import { apiFetcherSWR } from "@/utils/fetcher";

type AuthenticationContext = {
  isLoading: boolean;
  user?: User;
};

const AuthContext = createContext<AuthenticationContext>({
  isLoading: true,
  user: undefined,
});

type AuthProviderProps = {
  children?: ReactNode;
};

export function AuthProvider({ children }: AuthProviderProps) {
  const { data: user, isLoading } = useSWR<User, APIError>(
    "/users/self",
    apiFetcherSWR({ schema: UserSchema })
  );

  return (
    <AuthContext.Provider
      value={{
        isLoading,
        user,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuthContext() {
  return useContext(AuthContext);
}
