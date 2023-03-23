import {useState} from "react";
import Datepicker from "tailwind-datepicker-react"

export default function EstimateForm({submitted}) {
  const [neededBy, setNeededBy] = useState("");
  const [quantity, setQuantity] = useState(0);
  const [show, setShow] = useState(false);

  const handleQuantity = (e) => {
    setQuantity(e.target.value);
  }
  const handleSelectDate = (selectedDate) => {
    setNeededBy(selectedDate);
  }
  const handleClose = (state) => {
		setShow(state)
	}
  const defaultDate = new Date()
  //defaultDate.setDate(defaultDate.getDate() + 2 * 7);

  const datePickerOptions = {
    defaultDate: defaultDate,
    autoHide: true,
    todayBtn: false,
    clearBtn: false,
    theme: {
      input: "rounded-md border-0 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6",
      selected: "bg-pink-300",
    }
  }

  if(submitted){
    return (
      <div className="isolate bg-white py-24 px-6 sm:py-32 lg:px-8">
        <div className="mx-auto max-w-2xl text-center">
          <h2 className="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">Thank you!</h2>
        </div>
      </div>
      )
  }

  return (
    <div className="isolate bg-white py-24 px-6 sm:py-32 lg:px-8">
      <div className="mx-auto max-w-2xl text-center">
        <h2 className="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">Order Request</h2>
      </div>
      <form action="/x/estimates" method="POST" className="mx-auto mt-16 max-w-xl sm:mt-20">
        <div className="grid grid-cols-1 gap-y-6 gap-x-8 sm:grid-cols-2">
          <div>
            <label htmlFor="first-name" className="block text-sm font-semibold leading-6 text-gray-900">
              First Name
            </label>
            <div className="mt-2.5">
              <input
                type="text"
                name="first-name"
                id="first-name"
                autoComplete="given-name"
                className="block w-full rounded-md border-0 py-2 px-3.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
              />
            </div>
          </div>
          <div>
            <label htmlFor="last-name" className="block text-sm font-semibold leading-6 text-gray-900">
              Last Name
            </label>
            <div className="mt-2.5">
              <input
                type="text"
                name="last-name"
                id="last-name"
                autoComplete="family-name"
                className="block w-full rounded-md border-0 py-2 px-3.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
              />
            </div>
          </div>
          <div className="sm:col-span-2">
            <label htmlFor="email" className="block text-sm font-semibold leading-6 text-gray-900">
              Email
            </label>
            <div className="mt-2.5">
              <input
                type="email"
                name="email"
                id="email"
                autoComplete="email"
                className="block w-full rounded-md border-0 py-2 px-3.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
              />
            </div>
          </div>
          <div className="sm:col-span-2">
            <label htmlFor="phone-number" className="block text-sm font-semibold leading-6 text-gray-900">
              Phone Number
            </label>
            <div className="relative mt-2.5">
              <input
                type="tel"
                name="phone-number"
                id="phone-number"
                autoComplete="tel"
                className="block w-full rounded-md border-0 py-2 px-3.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
              />
            </div>
          </div>

          <div className="sm:col-span-2">
            <label htmlFor="theme" className="block text-sm font-semibold leading-6 text-gray-900">
              Cookie Theme
            </label>
            <div className="relative mt-2.5">
              <input
                type="text"
                name="theme"
                id="theme"
                autoComplete="off"
                className="block w-full rounded-md border-0 py-2 px-3.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
              />
            </div>
          </div>

          <div className="sm:col-span-1">
            <label htmlFor="quantity" className="block text-sm font-semibold leading-6 text-gray-900">
              Cookie Quantity (by the dozen)
            </label>
            <div className="relative mt-2.5">
              <input
                type="number"
                name="quantity"
                id="quantity"
                autoComplete="off"
                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
                aria-describedby="quantity-description"
                value={quantity}
                onChange={handleQuantity}
              />
            </div>
            <p className="mt-2 text-sm text-gray-500" id="quantity-description">
              {quantity} x 12 = {quantity * 12} cookies
            </p>
          </div>

          <div className="sm:col-span-1">
            <label htmlFor="needed-by" className="block text-sm font-semibold leading-6 text-gray-900">
              Requested Pick-Up Date
            </label>
            <div className="relative mt-2.5">
              <Datepicker options={datePickerOptions} onChange={handleSelectDate} show={show} setShow={handleClose} />
              <input
                type="hidden"
                name="needed-by"
                id="needed-by"
                autoComplete="off"
                className="block w-full rounded-md border-0 py-2 px-3.5 pl-20 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
                aria-describedby="needed-by-description"
                value={neededBy}
              />
            </div>
            <p className="mt-2 text-sm text-gray-500" id="needed-by-description">
              Minimum of two weeks notice
            </p>
          </div>

          <div className="sm:col-span-2">
            <label htmlFor="message" className="block text-sm font-semibold leading-6 text-gray-900">
              Anything else?
            </label>
            <div className="mt-2.5">
              <textarea
                name="message"
                id="message"
                rows={2}
                className="block w-full rounded-md border-0 py-2 px-3.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-200 sm:text-sm sm:leading-6"
                defaultValue={''}
              />
            </div>
          </div>

        </div>
        <div className="mt-10">
          <button
            type="submit"
            className="block w-full rounded-md bg-pink-400 px-3.5 py-2.5 text-center text-sm font-semibold text-white shadow-sm hover:bg-pink-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-pink-200"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  )
}