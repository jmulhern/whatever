export default function Hero() {
  return (
    <div className="relative isolate px-6 pt-6 sm:pt-12 lg:px-8">
      <div className="mx-auto max-w-xl content-center">
        <picture>
          <source type="image/webp" srcSet="/public/desertcatcookies-logo.webp" />
          <source type="image/png" srcSet="/public/desertcatcookies-logo.png" />
          <img width="1024" height="1024" src="/public/desertcatcookies-logo.png" alt="Desert Cat Cookies"/>
        </picture>
      </div>
    </div>
  )
}