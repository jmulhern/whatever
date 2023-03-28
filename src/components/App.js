import { default as Whatever } from "./whatever/Index";
import { default as DesertCatCookies } from "./desertcatcookies/Index";
import { default as GreasyShadows } from "./greasyshadows/Index";
import { default as TheBachelorette } from "./thebachelorette/Index";
import { default as HallOfFame } from "./hof/Index";


import { QueryClient, QueryClientProvider } from 'react-query';

const queryClient = new QueryClient()

export default function App() {
  const urlSearchParams = new URLSearchParams(window.location.search);
  const params = Object.fromEntries(urlSearchParams.entries());
  const parts = window.location.pathname.split("/");
  const cdn = window.cdn

  let site = <h1>???</h1>
  if (window.seed === 'whatever') {
    site = <Whatever cdn={cdn} params={params} parts={parts} />
  } else if (window.seed === 'desert-cat-cookies') {
    site = <DesertCatCookies cdn={cdn} params={params} parts={parts} />
  } else if (window.seed === 'greasy-shadows') {
    site = <GreasyShadows cdn={cdn} params={params} parts={parts} />
  } else if (window.seed  === 'the-bachelorette') {
    site = <TheBachelorette cdn={cdn} params={params} parts={parts} />
  } else if (window.seed === 'hall-of-fame') {
    site = <HallOfFame cdn={cdn} params={params} parts={parts} />
  }
  return (
    <QueryClientProvider client={queryClient}>
      { site }
    </QueryClientProvider>
  )
}