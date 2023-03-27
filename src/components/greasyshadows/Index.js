import Track from "./Track";
import Album from "./Album";

export default function Index() {
  return (
      <div className="bg-black mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-3xl">
          <div className="flex flex-col justify-center items-center">

            <div className="p-8 text-white font-bold justify-center text-3xl">
              <h1>Spiritual Warfare and the Greasy Shadows.</h1>
            </div>

            <div className="py-4">
              <Album id={1483488955}
                     url="http://spiritualwarfare.bandcamp.com/album/ad-hoc"
                     title="'ad hoc' by Spiritual Warfare and the Greasy Shadows" />
            </div>

            <div className="py-4">
              <Track id={3979850911}
                     url="http://spiritualwarfare.bandcamp.com/track/take-it-over"
                     title="Take It Over by SPIRITUAL WARFARE AND THE GREASY SHADOWS!" />
            </div>

            <div className="py-4">
              <Track id={2901956873}
                     url="http://spiritualwarfare.bandcamp.com/track/death-was-an-olympic-speedskater"
                     title="Death Was An Olympic Speedskater by SPIRITUAL WARFARE AND THE GREASY SHADOWS" />
            </div>

            <div className="py-4">
              <Album id={2731713295}
                     url="http://spiritualwarfare.bandcamp.com/album/i-hope-my-grave-is-a-gutter"
                     title="I Hope My Grave Is A Gutter by SPIRITUAL WARFARE AND THE GREASY SHADOWS" />
            </div>

            <div className="py-4">
              <Album id={2018954812}
                     url="http://spiritualwarfare.bandcamp.com/album/double-voices"
                     title="Double Voices by SPIRITUAL WARFARE AND THE GREASY SHADOWS" />
            </div>

            <div className="py-4">
              <Album id={3889473414}
                     url="http://spiritualwarfare.bandcamp.com/album/vol-1"
                     title="Vol. 1 by SPIRITUAL WARFARE AND THE GREASY SHADOWS" />
            </div>

            <div className="py-4">
              <Album id={3882193566}
                     url="http://spiritualwarfare.bandcamp.com/album/suite-16"
                     title="Suite 16 by SPIRITUAL WARFARE AND THE GREASY SHADOWS" />
            </div>

            <div className="p-8 text-white font-bold justify-center text-3xl">
              <h1>...</h1>
            </div>

          </div>
        </div>
      </div>
    )
}