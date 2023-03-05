import Link from "next/link";
import "@/styles/globals.css";
import Image from "next/image";
import Footer from "@/app/components/Footer";

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
                    <Image src="/logo.png" alt="chess Logo" width={30} height={30} />
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
