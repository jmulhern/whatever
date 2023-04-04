import {useState} from "react";

export default function Movie({movie, classNames}) {
  if (localStorage.getItem(movie.unique) === null) {
    localStorage.setItem(movie.unique, false.toString())
  }
  const [watched, setWatched] = useState(/^true$/i.test(localStorage.getItem(movie.unique)))
  const handleChange = (e) => {
    localStorage.setItem(movie.unique, e.target.checked)
    setWatched(e.target.checked)
  };
  return (
    <fieldset key={movie.unique}>
      <legend className="sr-only">Movies</legend>
      <div className="space-y-5">
        <div className="relative flex items-start">
          <div className="flex h-6 items-center">
            <input
              id={movie.unique}
              aria-describedby={movie.unique+"-description"}
              name={movie.unique}
              type="checkbox"
              className={"h-4 w-4 rounded border-gray-300 "+classNames}
              onChange={handleChange}
              checked={watched}
            />
          </div>
          <div className="ml-3 text-sm leading-6">
            <label htmlFor={movie.unique} className="font-semibold text-gray-900">
              {movie.title}
            </label>{' '}
            <span id={movie.unique+"-description"} className="text-xs text-gray-500">
                  <span className="sr-only">{movie.title} </span>{movie.released}
                </span>
          </div>
        </div>
      </div>
    </fieldset>
  )
}
