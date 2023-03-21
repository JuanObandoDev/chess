import { API_ADDRESS } from "@/config";
import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest) {
  const { searchParams, origin } = new URL(request.url);
  const body = {
    code: searchParams.get("code"),
    state: searchParams.get("state"),
  };

  const response = await fetch(`${API_ADDRESS}/auth/discord/callback`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
    mode: "cors",
    credentials: "include",
  });

  const redirect = NextResponse.redirect(origin);
  const cookie = response.headers.get("Set-Cookie");
  if (cookie) {
    redirect.headers.set("Set-Cookie", cookie);
  }
  return redirect;
}
