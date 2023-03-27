export default function Hero({cdn}) {
  return (
    <div className="relative isolate px-6 pt-6 sm:pt-12 lg:px-8">
      <div className="mx-auto max-w-xl content-center">
        <picture>
          <img width="1024" height="1024" src={cdn+"/public/desertcatcookies-logo.webp"} alt="Desert Cat Cookies"/>
        </picture>
      </div>
    </div>
  )
}