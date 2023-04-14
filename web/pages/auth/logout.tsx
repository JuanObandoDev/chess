import { withAuthSSR } from "@/utils/auth";

export const getServerSideProps = withAuthSSR((ctx) => {
  ctx.req.session.destroy();
  ctx.res.setHeader("clear-site-data", '"cookies"');
  return {
    redirect: {
      destination: "/",
      permanent: false,
    },
    props: {},
  };
});

export default function Logout() {
  return null;
}
