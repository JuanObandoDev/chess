import { API_ADDRESS } from "@/config";

async function getUser(payload: any): Promise<string> {
  const req = new Request(`${API_ADDRESS}/auth/discord/callback`, {
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  const res = await fetch(req, {
    method: "POST",
    body: JSON.stringify(payload),
    cache: "no-store",
  });

  const data = await res.json();
  return data;
}

export default async function Discord({
  params,
  searchParams,
}: {
  params: { slug: string };
  searchParams?: { [key: string]: string | string[] | undefined };
}) {
  const data = await getUser(searchParams);

  return (
    <>
      <main>{JSON.stringify(data)}</main>
    </>
  );
}
