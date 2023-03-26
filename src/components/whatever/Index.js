import {useQuery} from 'react-query'
import axios from "axios";

export default function Index() {
  const { isLoading, error, data } = useQuery({
    queryKey: ['things'],
    queryFn: () => axios.get("/x/things")
    .then((res) => res.data)
  })
  if (error) return "An error has occurred: " + error.message;
  if (isLoading) return "Loading..."
  console.log(data)
  return (
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-3xl">
          <div className="flex flex-col justify-center items-center">
            <div className="p-12">
              <div><pre>nothing</pre></div>
            </div>
          </div>
        </div>
      </div>
    )
}