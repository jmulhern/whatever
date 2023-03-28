import Header from "./Header";
import Frame from "./Frame";

const all = [
  { unique: 'clarice', kind: 'pet', name: 'Clarice'  },
  { unique: 'gigi', kind: 'pet', name: 'Gigi'  },
  { unique: 'ginger', kind: 'pet', name: 'Ginger 🕊'  },
  { unique: 'chuck', kind: 'foster', name: 'Chuck' },
  { unique: 'meg', kind: 'foster', name: 'Meg' },
  { unique: 'toofer', kind: 'foster', name: 'James William Bottomtooth III' },
  { unique: 'missy', kind: 'foster', name: 'Missy' },
  { unique: 'ollie', kind: 'foster', name: 'Ollie' },
  { unique: 'captain-lee', kind: 'foster', name: 'Captain Lee' },
  { unique: 'ruby', kind: 'foster', name: 'Ruby' },
  { unique: 'luna-w', kind: 'foster', name: 'Luna W' },
  { unique: 'peaches', kind: 'foster', name: 'Peaches' },
  { unique: 'tessa', kind: 'foster', name: 'Tessa' },
  { unique: 'casper', kind: 'foster', name: 'Casper' },
  { unique: 'koi', kind: 'foster', name: 'Koi' },
]

export default function Index({cdn, params}) {
  let spotlight = "pet" // default
  if (params !== undefined && params.spotlight !== undefined && params.spotlight !== "") {
    spotlight = params.spotlight
  }
  let frames = []
  all.forEach((frame) => {
    if(spotlight === frame.kind) {
      frames.push(frame)
    }
  })
  return (
    <div className="mb-4">
      <div className="md:flex md:items-center md:justify-between">
        <div className="min-w-0 flex-1">
          <Header spotlight={spotlight} />
        </div>
      </div>
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-3xl">
          <div className="overflow-hidden bg-white shadow sm:rounded-md">
            <ul role="list" className="divide-y divide-gray-200">
              {frames.map((frame) => (
                <Frame key={frame.unique} cdn={cdn} name={frame.name} unique={frame.unique} />
              ))}
            </ul>
          </div>
        </div>
      </div>
    </div>

  )
}