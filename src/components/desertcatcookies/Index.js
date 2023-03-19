import Container from "./Container";
import EstimateForm from "./EstimateForm";
import Hero from "./Hero";

export default function Index({params}) {
  return (
    <Container>
      <Hero />
      <EstimateForm submitted={params.submitted === "true"}/>
    </Container>
  )
}