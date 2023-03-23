import {useQuery} from 'react-query'
import axios from "axios";
import Cost from "./Cost";

export default function Index() {
  const { isLoading, error, data } = useQuery({
    queryKey: ['costs'],
    queryFn: () => axios.get("/x/costs")
    .then((res) => res.data)
  })
  if (error) return "An error has occurred: " + error.message;
  if (isLoading) return "Loading..."
  return (
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-3xl">
          <div className="flex flex-col justify-center items-center">
            <div className="p-12">
              <div><pre>{data.map((x)=>(
                <Cost key={x.name} name={x.name} amount={x.amount} />
              ))}</pre></div>
            </div>
          </div>
        </div>
      </div>
    )
}