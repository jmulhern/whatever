export default function Error({status, message}) {
  return (
    <>
      <main className="relative isolate min-h-full">
        <img src="/public/mad-cat.jpg" alt="" className="absolute inset-0 -z-10 h-full w-full object-cover object-top" />
        <div className="mx-auto max-w-7xl px-6 py-32 text-center sm:py-40 lg:px-8">
          <p className="text-base font-semibold leading-8 text-white">{status}</p>
          <h1 className="mt-4 text-3xl font-bold tracking-tight text-white sm:text-5xl">{message}</h1>
          <div className="mt-10 flex justify-center">
            <a href="/" className="text-sm font-semibold leading-7 text-white">
              <span aria-hidden="true">&larr;</span> Back to home
            </a>
          </div>
        </div>
      </main>
    </>
  )
}
