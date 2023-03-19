import { default as DesertCatCookies } from "./desertcatcookies/Index";
import { default as GreasyShadows } from "./greasyshadows/Index";
import { default as TheBachelorette } from "./thebachelorette/Index";
import { default as Whatever } from "./whatever/Index";
export default function App() {
  const urlSearchParams = new URLSearchParams(window.location.search);
  const params = Object.fromEntries(urlSearchParams.entries());
  const parts = window.location.pathname.split("/")

  if (window.thing === 'whatever') {
    return (
      <Whatever params={params} parts={parts} />
    )
  } else if (window.thing === 'greasy-shadows') {
    return (
      <GreasyShadows params={params} parts={parts} />
    )
  } else if (window.thing  === 'the-bachelorette') {
    return (
      <TheBachelorette params={params} parts={parts} />
    )
  } else if (window.thing === 'desert-cat-cookies') {
    return (
      <DesertCatCookies params={params} parts={parts} />
    )
  } else {
    return (
      <h1>???</h1>
    )
  }
}