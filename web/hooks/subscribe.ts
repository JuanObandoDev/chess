import { APIError } from "@/types/Error";
import { WEBSOCKET_ADDRESS } from "@/utils/config";
import { parseAPIError } from "@/utils/error";
import { useEffect, useRef, useState } from "react";

export function useSubscribe<T>(key: string) {
  const ws = useRef<undefined | WebSocket>(undefined);
  const [isLoading, setIsLoading] = useState(true);
  const [data, setData] = useState<undefined | T>(undefined);
  const [error, setError] = useState<undefined | APIError>(undefined);

  useEffect(() => {
    ws.current = new WebSocket(`${WEBSOCKET_ADDRESS}${key}`);
    ws.current.addEventListener("open", () => {
      if (ws.current?.readyState === WebSocket.OPEN) setIsLoading(false);
    });
    ws.current.addEventListener("message", (event) => {
      setData(event.data);
    });

    ws.current.addEventListener("error", (event) => {
      setError(parseAPIError(new Error("WebSocket error")));
    });

    ws.current.addEventListener("close", () => {
      setIsLoading(true);
    });

    return () => {
      if (ws.current?.readyState === WebSocket.OPEN) ws.current.close();
    };
  }, [key]);

  return {
    isLoading,
    ws: ws.current,
    data,
    error,
  };
}
