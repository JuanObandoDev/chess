import { API_ADDRESS } from "@/config";

async function getLink(): Promise<string> {
  const req = new Request(`${API_ADDRESS}/auth/github/link`, {
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  const res = await fetch(req, { cache: "force-cache" });
  const data = await res.json();
  return data.url;
}

export default async function Login() {
  const link = await getLink();

  return (
    <>
      <main>
        <a href={link}>Github</a>
      </main>
    </>
  );
}
