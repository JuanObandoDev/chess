import Link from "next/link";

import { useAuthContext } from "@/context/auth";

export default function Home() {
  const { user, isLoading } = useAuthContext();

  return (
    <>
      {user ? <div> {user.name}</div> : <Link href={"/auth/login"}>login</Link>}
    </>
  );
}
