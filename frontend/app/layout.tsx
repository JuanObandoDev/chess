import Link from "next/link";
import "@/styles/globals.css";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
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
        <footer>
          <p>Copyright © 2023 SIS</p>
        </footer>
      </body>
    </html>
  );
}
