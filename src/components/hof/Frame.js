
export default function Frame({cdn, name, unique}) {
  return (
    <li className="px-4 py-4 sm:px-6">
      <img width="800" height="600" src={cdn+"/public/hof/"+unique+".webp"} alt={name}></img>
      <p className="pt-2 text-center text-3xl font-serif font-bold">{name}</p>
    </li>
  )
}
