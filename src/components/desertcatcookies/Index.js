import Container from "./Container";
import EstimateForm from "./EstimateForm";
import Hero from "./Hero";

export default function Index({cdn, params}) {
  return (
    <Container>
      <Hero cdn={cdn} />
      <EstimateForm submitted={params.submitted === "true"}/>
    </Container>
  )
}