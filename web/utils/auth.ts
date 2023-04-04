import { GetServerSidePropsContext } from "next";

import { CallbackSchema } from "@/types/GitHub";
import { UserSchema } from "@/types/User";
import { apiFetcher, isSuccessStatus } from "@/utils/fetcher";
import { handleError, parseAPIError } from "@/utils/error";

export async function auth(
  provider: "discord" | "github",
  ctx: GetServerSidePropsContext
): Promise<string> {
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

  return setCookie;
}
