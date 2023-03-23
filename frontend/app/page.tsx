import { API_ADDRESS } from "@/config";
import styles from "@/styles/page.module.css";
import { cookies } from "next/headers";

export default async function Home() {
  const cookieStore = cookies();

  const session = cookieStore.get("x-session");

  let user = null;

  if (session) {
    const response = await fetch(`${API_ADDRESS}/users/self`, {
      headers: {
        "Content-Type": "application/json",
        Cookie: session.name + "=" + session.value,
      },
    });

    if (response.status === 200) {
      const data = await response.json();
      user = data;
    }
  }

  return (
    <>
      {!user ? (
        <main className={styles.main}>
          <h2>
            Thank you for choosing our web site. Please login or signup to start
            to play.
          </h2>
        </main>
      ) : (
        <main className={styles.main}>
          <h2>
            Welcome {user.name}! You can start to play by clicking on the Play
            button.
          </h2>
        </main>
      )}
    </>
  );
}
