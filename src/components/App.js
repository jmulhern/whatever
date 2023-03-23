import { default as DesertCatCookies } from "./desertcatcookies/Index";
import { default as GreasyShadows } from "./greasyshadows/Index";
import { default as TheBachelorette } from "./thebachelorette/Index";
import { default as Whatever } from "./whatever/Index";
import { QueryClient, QueryClientProvider } from 'react-query';

const queryClient = new QueryClient()

export default function App() {
  const urlSearchParams = new URLSearchParams(window.location.search);
  const params = Object.fromEntries(urlSearchParams.entries());
  const parts = window.location.pathname.split("/");

  let site = <h1>???</h1>
  if (window.seed === 'whatever') {
    site = <Whatever params={params} parts={parts} />
  } else if (window.seed === 'greasy-shadows') {
    site = <GreasyShadows params={params} parts={parts} />
  } else if (window.seed  === 'the-bachelorette') {
    site = <TheBachelorette params={params} parts={parts} />
  } else if (window.seed === 'desert-cat-cookies') {
    site = <DesertCatCookies params={params} parts={parts} />
  }
  return (
    <QueryClientProvider client={queryClient}>
      { site }
    </QueryClientProvider>
  )
}