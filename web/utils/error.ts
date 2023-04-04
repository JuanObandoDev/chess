import { APIError, APIErrorSchema } from "@/types/Error";

export function parseAPIError(e: unknown): APIError {
  const parsed = APIErrorSchema.safeParse(e);
  if (parsed.success) {
    return parsed.data;
  } else if (e instanceof Error) {
    return { error: e.toString() };
  } else if (e === undefined) {
    return {
      error: "Unknown error",
      message: "deriveError was passed a value of `undefined`.",
    };
  } else {
    return {
      error: String(e),
      message:
        "deriveError was passed a value that was neither Error or APIError.",
    };
  }
}

export function handleError(raw: unknown, response: Response): never {
  throw parseAPIError(
    raw ?? { error: `${response.status}: ${response.statusText}` }
  );
}
