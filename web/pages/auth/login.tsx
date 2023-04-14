import { useLink } from "@/hooks/auth";
import { withoutAuthSSR } from "@/utils/auth";

export const getServerSideProps = withoutAuthSSR(() => ({ props: {} }));

export default function Login() {
  const {
    link: linkDiscord,
    error: errorDiscord,
    isLoading: isLoadingDiscord,
  } = useLink("discord");

  const {
    link: linkGithub,
    error: errorGithub,
    isLoading: isLoadingGithub,
  } = useLink("github");

  return (
    <>
      <a href={`${linkGithub?.url}`}>Github</a>
      <a href={`${linkDiscord?.url}`}>Discord</a>
    </>
  );
}
