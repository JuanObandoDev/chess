import { useSubscribe } from "@/hooks/subscribe";

export default function Games() {
  const { isLoading, ws, data, error } = useSubscribe<string>("/games/ws");

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {JSON.stringify(error)}</div>;

  return (
    <>
      <button
        onClick={() => {
          if (!isLoading) {
            if (ws?.readyState === WebSocket.OPEN) {
              ws.send(new Date().toISOString());
            }
          }
        }}
      >
        send
      </button>
      {data && <div>{data}</div>}
    </>
  );
}
