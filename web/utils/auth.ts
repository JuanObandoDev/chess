import { GetServerSidePropsContext, GetServerSidePropsResult } from "next";

import { CallbackSchema } from "@/types/GitHub";
import { User, UserSchema } from "@/types/User";
import { apiFetcher, isSuccessStatus } from "@/utils/fetcher";
import { handleError, parseAPIError } from "@/utils/error";
import { withSessionSsr } from "./session";

export function withAuthSSR<
  P extends {
    [key: string]: unknown;
  } = {
    [key: string]: unknown;
  }
>(
  handler: (
    context: GetServerSidePropsContext
  ) => GetServerSidePropsResult<P> | Promise<GetServerSidePropsResult<P>>,
  destination = "/auth/login"
) {
  return withSessionSsr<P>((ctx) => {
    if (ctx.req.session.user === undefined)
      return {
        redirect: {
          destination,
          permanent: false,
        },
        props: {},
      };
    return handler(ctx);
  });
}

export function withoutAuthSSR<
  P extends {
    [key: string]: unknown;
  } = {
    [key: string]: unknown;
  }
>(
  handler: (
    context: GetServerSidePropsContext
  ) => GetServerSidePropsResult<P> | Promise<GetServerSidePropsResult<P>>,
  destination = "/dashboard"
) {
  return withSessionSsr<P>((ctx) => {
    if (ctx.req.session.user !== undefined)
      return {
        redirect: {
          destination,
          permanent: false,
        },
        props: {},
      };
    return handler(ctx);
  });
}

export async function getAuth(
  provider: "discord" | "github",
  ctx: GetServerSidePropsContext
): Promise<{ user: User; setCookie: string }> {
  const payload = CallbackSchema.parse(ctx.query);

  const { response, data } = await apiFetcher(`/auth/${provider}/callback`, {
    method: "POST",
    body: JSON.stringify(payload),
    schema: UserSchema,
  });

  if (!isSuccessStatus(response.status)) {
    handleError(data, response);
  }

  const setCookie = response.headers.get("set-cookie");
  if (!setCookie) {
    throw parseAPIError({
      error: "No cookie set",
      message: `The server did not set a cookie after authenticating with ${provider}`,
    });
  }

  return { user: data, setCookie };
}
