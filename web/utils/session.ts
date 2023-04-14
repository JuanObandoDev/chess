import { GetServerSidePropsContext, GetServerSidePropsResult } from "next";
import { IronSessionOptions } from "iron-session";
import { withIronSessionSsr } from "iron-session/next";

import { User } from "@/types/User";

export const ironSessionOptions: IronSessionOptions = {
  cookieName: "x-chess",
  password: process.env.SECRET_COOKIE_PASSWORD as string,
  cookieOptions: {
    secure: process.env.NODE_ENV === "production",
  },
};

export function withSessionSsr<
  P extends {
    [key: string]: unknown;
  } = {
    [key: string]: unknown;
  }
>(
  handler: (
    context: GetServerSidePropsContext
  ) => GetServerSidePropsResult<P> | Promise<GetServerSidePropsResult<P>>
) {
  return withIronSessionSsr<P>(handler, ironSessionOptions);
}

declare module "iron-session" {
  interface IronSessionData {
    user?: User;
  }
}
