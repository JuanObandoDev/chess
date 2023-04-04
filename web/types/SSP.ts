import { APIError } from "@/types/Error";

export type SSP<T> =
  | {
      success: true;
      data: T;
    }
  | {
      success: false;
      error: APIError;
    };
