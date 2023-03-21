import { default as DesertCatCookies } from "./desertcatcookies/Index";
import { default as GreasyShadows } from "./greasyshadows/Index";
import { default as TheBachelorette } from "./thebachelorette/Index";
import { default as Whatever } from "./whatever/Index";
export default function App() {
  const urlSearchParams = new URLSearchParams(window.location.search);
  const params = Object.fromEntries(urlSearchParams.entries());
  const parts = window.location.pathname.split("/")

  if (window.seed === 'whatever') {
    return (
      <Whatever params={params} parts={parts} />
    )
  } else if (window.seed === 'greasy-shadows') {
    return (
      <GreasyShadows params={params} parts={parts} />
    )
  } else if (window.seed  === 'the-bachelorette') {
    return (
      <TheBachelorette params={params} parts={parts} />
    )
  } else if (window.seed === 'desert-cat-cookies') {
    return (
      <DesertCatCookies params={params} parts={parts} />
    )
  } else {
    return (
      <h1>???</h1>
    )
  }
}