import Link from "next/link";
import "@/styles/globals.css";
import { cookies } from "next/headers";
import { API_ADDRESS } from "@/config";
import Image from "next/image";
import Footer from "@/app/components/Footer";

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
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

  console.log(user);

  return (
    <html lang="en">
      <head />
      <body>
        <header>
          <nav>
            <ul>
              <div>
                <li>
                  <Link href="/" className="logo">
                    <Image
                      src="/logo.png"
                      alt="chess Logo"
                      width={30}
                      height={30}
                    />
                    Chess
                  </Link>
                </li>
              </div>
              <div className="links">
                <li>
                  <Link href="/login" className="login">
                    Login
                  </Link>
                </li>
              </div>
            </ul>
          </nav>
        </header>
        {children}
        <Footer />
      </body>
    </html>
  );
}
