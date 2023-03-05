import { API_ADDRESS } from "@/config";

async function getGithubLink(): Promise<string> {
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

async function getDiscordLink(): Promise<string> {
  const req = new Request(`${API_ADDRESS}/auth/discord/link`, {
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
  const githubLink = await getGithubLink();
  const discordLink = await getDiscordLink();

  return (
    <>
      <main>
        <a href={githubLink}>Github</a>
        <a href={discordLink}>Discord</a>
      </main>
    </>
  );
}
