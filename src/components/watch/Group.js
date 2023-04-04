import Movie from "./Movie";

export default function Group({name, unique, movies}) {
  let classNames = "text-gray-600 focus:ring-gray-600"
  if(unique === "mcu") {
    classNames = "text-red-600 focus:ring-red-600"
  } else if (unique ===  "star-wars") {
    classNames = "text-yellow-400 focus:ring-yellow-400"
  }

  return (
    <li className="px-4 py-4 sm:px-6">
      <div className="relative mb-2">
        <div className="relative flex justify-start">
          <span className="bg-white pr-3 text-base font-bold leading-6 text-3xl text-gray-900">{name}</span>
        </div>
      </div>
      {movies.map((movie) => (
        <Movie key={movie.unique} movie={movie} classNames={classNames} />
      ))}
    </li>
  )
}
