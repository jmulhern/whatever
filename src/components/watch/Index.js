import Group from "./Group";
import {useEffect, useState} from "react";
import {useQuery} from "react-query";
import axios from "axios";

export default function Index() {
  const { isLoading, error, data } = useQuery({
    queryKey: ['movies'],
    queryFn: () => axios.get("/x/movies")
    .then((res) => res.data)
  })
console.log(data)
  return (
    <div className="my-4">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-3xl">
          <div className="overflow-hidden bg-white shadow rounded-md">
            <ul role="list" className="divide-y divide-gray-200">
              {error &&  <span>Shit - {error.message}</span>}
              {isLoading && <span>Loading...</span>}
              {data && data.map((group) => <Group key={group.unique} cdn={cdn} unique={group.unique} name={group.name} movies={group.movies} />)}
            </ul>
          </div>
        </div>
      </div>
    </div>

  )
}
