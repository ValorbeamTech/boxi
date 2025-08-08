// Local Imports
import { Hotel, HotelCard } from "./HotelCard";

// ----------------------------------------------------------------------

const hotels: Hotel[] = [
  {
    id: "1",
    cover: "/images/800x600.png",
    name: "Emerald Bay Inn.",
    tag: "plus",
    price: 100,
    rating: 4.9,
    bedsCount: 3,
    adultCount: 3,
  },
  {
    id: "2",
    cover: "/images/800x600.png",
    name: "Crowne Plaza.",
    tag: "lux",
    price: 80,
    rating: 4.8,
    bedsCount: 2,
    adultCount: 4,
  },
  {
    id: "3",
    cover: "/images/800x600.png",
    name: "Sunset Lodge.",
    tag: "penthouse",
    price: 100,
    rating: 4.9,
    bedsCount: 3,
    adultCount: 3,
  },
  {
    id: "4",
    cover: "/images/800x600.png",
    name: "Hotel Elite.",
    tag: "plus",
    price: 120,
    rating: 4.9,
    bedsCount: 1,
    adultCount: 2,
  },
  {
    id: "5",
    cover: "/images/800x600.png",
    name: "Hotel Bliss.",
    tag: "lux",
    price: 90,
    rating: 4.5,
    bedsCount: 2,
    adultCount: 4,
  },
  {
    id: "6",
    cover: "/images/800x600.png",
    name: "Emerald Bay Inn.",
    tag: "plus",
    price: 100,
    rating: 4.9,
    bedsCount: 3,
    adultCount: 3,
  },
  {
    id: "7",
    cover: "/images/800x600.png",
    name: "Crowne Plaza.",
    tag: "lux",
    price: 80,
    rating: 4.8,
    bedsCount: 2,
    adultCount: 4,
  },
  {
    id: "8",
    cover: "/images/800x600.png",
    name: "Sunset Lodge.",
    tag: "penthouse",
    price: 100,
    rating: 4.9,
    bedsCount: 3,
    adultCount: 3,
  },
  {
    id: "9",
    cover: "/images/800x600.png",
    name: "Hotel Elite.",
    tag: "plus",
    price: 120,
    rating: 4.9,
    bedsCount: 1,
    adultCount: 2,
  },
  {
    id: "10",
    cover: "/images/800x600.png",
    name: "Hotel Bliss.",
    tag: "lux",
    price: 90,
    rating: 4.5,
    bedsCount: 2,
    adultCount: 4,
  },
];

export function TopHotels() {
  return (
    <div className="transition-content mt-4 pl-(--margin-x) sm:mt-5 lg:mt-6">
      <div className="bg-gray-150 dark:bg-dark-800 rounded-l-lg pt-4 pb-1">
        <h2 className="dark:text-dark-100 truncate px-4 text-base font-medium tracking-wide text-gray-800 sm:px-5 lg:text-lg">
          Top Hotels
        </h2>
        <div
          className="custom-scrollbar mt-4 flex space-x-4 overflow-x-auto px-4 pb-4 sm:px-5"
          style={{ "--margin-scroll": "1.25rem" } as React.CSSProperties}
        >
          {hotels.map((hotel) => (
            <HotelCard key={hotel.id} {...hotel} />
          ))}
        </div>
      </div>
    </div>
  );
}
