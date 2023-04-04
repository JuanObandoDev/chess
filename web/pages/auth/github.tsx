import { APIError } from "@/types/Error";
import { parseAPIError } from "@/utils/error";
import { auth } from "@/utils/auth";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";

type Props = {
  error?: APIError;
};

export const getServerSideProps: GetServerSideProps<Props> = async (ctx) => {
  try {
    const cookie = await auth("github", ctx);
    ctx.res.setHeader("set-cookie", cookie);
    ctx.res.writeHead(302, { Location: "/" });
    ctx.res.end();
    return { props: {} };
  } catch (e) {
    return { props: { error: parseAPIError(e) } };
  }
};

export default function Github({
  error,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  return (
    <>
      <div>{error?.error}</div>
      <div>{error?.message}</div>
      <div>{error?.suggested}</div>
    </>
  );
}
