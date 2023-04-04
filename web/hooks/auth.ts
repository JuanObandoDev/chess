import useSWR from "swr";

import { APIError } from "@/types/Error";
import { Link, LinkSchema } from "@/types/GitHub";
import { apiFetcherSWR } from "@/utils/fetcher";

export function useLink(provider: "discord" | "github") {
  const { data, isLoading, error } = useSWR<Link, APIError>(
    `/auth/${provider}/link`,
    apiFetcherSWR({ schema: LinkSchema })
  );

  return {
    link: data,
    isLoading: isLoading,
    error: error,
  };
}
