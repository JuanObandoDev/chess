import { API_ADDRESS } from "@/config";
import styles from "@/styles/login.module.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGithub, faDiscord } from "@fortawesome/free-brands-svg-icons";

async function getGithubLink(): Promise<string> {
  const req = new Request(`${API_ADDRESS}/auth/github/link`, {
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  const res = await fetch(req, { cache: "no-store" });
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

  const res = await fetch(req, { cache: "no-store" });
  const data = await res.json();
  return data.url;
}

export default async function Login() {
  const githubLink = await getGithubLink();
  const discordLink = await getDiscordLink();

  return (
    <>
      <main className={styles.main}>
        <h1 className={styles.title}>Login</h1>
        <p className={styles.description}>
          Login with your Github or Discord account
        </p>
        <a href={githubLink} className={styles.github}>
          {" "}
          <FontAwesomeIcon icon={faGithub} className={styles.ghicon} /> Login
          with Github
        </a>
        <p className={styles.or}>or</p>
        <a href={discordLink} className={styles.discord}>
          {" "}
          <FontAwesomeIcon icon={faDiscord} className={styles.drdicon} /> Login
          with Discord
        </a>
      </main>
    </>
  );
}
