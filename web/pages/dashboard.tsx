import { InferGetServerSidePropsType } from "next";

import { User } from "@/types/User";
import { withAuthSSR } from "@/utils/auth";

type DashboardProps = {
  user: User;
};

export const getServerSideProps = withAuthSSR<DashboardProps>((ctx) => ({
  props: { user: ctx.req.session.user as User },
}));

export default function Dashboard({
  user,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  return (
    <main>
      <div>{user.name}</div>
      <div>{user.email}</div>
    </main>
  );
}
