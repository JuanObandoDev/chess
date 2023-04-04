import { GetServerSidePropsContext } from "next";
import { ZodSchema } from "zod";

import { API_ADDRESS } from "@/utils/config";
import { handleError, parseAPIError } from "@/utils/error";
import { SSP } from "@/types/SSP";

type Opts<T> = {
  schema?: ZodSchema<T>;
  query?: URLSearchParams;
  isSetCookie?: boolean;
};

export type APIOptions<T> = Opts<T> & RequestInit;

export function apiFetcherSWR<T>(opts?: APIOptions<T>) {
  return (path: string): Promise<T> =>
    apiFetcher(path, opts).then(({ data }) => data);
}

export function apiFetcherSSP<T>(
  path: string,
  ctx: GetServerSidePropsContext,
  opts?: APIOptions<T>
): Promise<SSP<T>> {
  const headers = new Headers(opts?.headers);

  if (ctx.req.headers.cookie) {
    headers.append("cookie", ctx.req.headers.cookie);
  }

  return apiFetcher(path, { ...opts, headers })
    .then(({ data }) => ({
      success: true as const,
      data,
    }))
    .catch((e) => ({
      success: false as const,
      error: parseAPIError(e),
    }));
}

export async function apiFetcher<T>(
  path: string,
  opts?: APIOptions<T>
): Promise<{ data: T; response: Response }> {
  const pathWithQuery = `${path}${
    opts?.query ? "?" + opts.query.toString() : ""
  }`;

  const request = buildRequest(pathWithQuery, opts);
  const response = await fetch(request);
  console.log("response", response);
  const data = await getData(response);

  if (!isSuccessStatus(response.status)) {
    handleError(data, response);
  }

  return { data: opts?.schema?.parse(data) ?? (data as T), response: response };
}

export const buildRequest = (path: string, opts?: RequestInit): Request => {
  const req = new Request(`${"http://localhost:8080"}${path}`, {
    headers: {
      "Content-Type": "application/json",
    },
    ...opts,
  });
  return req;
};

export function isSuccessStatus(code: number) {
  return code >= 200 && code <= 299;
}

export async function getData(response: Response): Promise<unknown> {
  if (isJSON(response)) return response.json();
  return undefined;
}

function isJSON(response: Response): boolean {
  const contentType = response.headers.get("Content-Type");
  return contentType === "application/json";
}
