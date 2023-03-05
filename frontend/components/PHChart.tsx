import React from "react";

interface IProps {
  ph?: number | null; // 0-14
  index?: number | null; // 0-14
}

export default function PHChart({ ph, index }: IProps) {
  let roundedPh = 0;
  let roundedIndex = 0;
  let phPosition = 0;
  let indexPosition = 0;
  let offset = 0;

  // When ph and index are present.
  if (ph) {
    roundedPh = Math.round(ph * 100) / 100;
    phPosition = (ph / 14) * 100; // in percentage
  }
  if (index) {
    roundedIndex = Math.round(index * 100) / 100;
    indexPosition = (index / 14) * 100; // in percentage
  }

  // Add offset (position) to each value if values are too close.
  if (Math.abs(roundedPh - roundedIndex) < 1.5) {
    // Check which one is higher.
    if (roundedPh > roundedIndex) {
      offset = 4;
    } else {
      offset = -4;
    }
  } else {
    offset = 0;
  }

  return (
    // Chart
    <div className="relative flex flex-col h-full justify-center gap-1 mx-10">
      {/* Heading */}
      <h2 className="text-center absolute bottom-full left-0 right-0 mx-auto">
        Skala PH
      </h2>

      {/* Scale & numbers wrapper */}
      <div className="relative w-full h-full">
        {/* Fill */}
        <div
          style={{
            backgroundImage:
              "linear-gradient(0deg, rgba(247,4,4,1) 0%, rgba(255,248,0,1) 25%, rgba(7,195,0,1) 50%, rgba(22,73,250,1) 75%, rgba(45,4,144,1) 100%)",
          }}
          className="relative w-32 h-full rounded-lg"
        >
          {/* PH line */}
          <div
            style={{ bottom: `${phPosition}%` }}
            className={`w-full h-1 bg-green-500 left-0 absolute transform -translate-y-1/2 transition-all duration-700 ${
              ph ? "opacity-100" : "opacity-0"
            }`}
          ></div>
          {/* Index line */}
          <div
            style={{ bottom: `${indexPosition}%` }}
            className={`w-full h-1 bg-purple-500 left-0 absolute transform -translate-y-1/2 transition-all duration-1000 ${
              ph ? "opacity-100" : "opacity-0"
            }`}
          ></div>
        </div>
        {/* PH scale 'numbers' */}
        <span className="text-md absolute top-0 left-full ml-1">14</span>
        <span
          style={{ bottom: `${offset ? phPosition + offset : phPosition}%` }}
          className={`text-xl font-bold right-full mr-2 absolute transform translate-y-[35%] transition-all duration-700 ${
            ph ? "opacity-100" : "opacity-0"
          }`}
        >
          {roundedPh}
        </span>
        <span
          style={{
            bottom: `${offset ? indexPosition - offset : indexPosition}%`,
          }}
          className={`text-xl font-bold right-full mr-2 absolute transform translate-y-[35%] transition-all duration-1000 ${
            ph ? "opacity-100" : "opacity-0"
          }`}
        >
          {roundedIndex}
        </span>
        <span className="text-md absolute bottom-0 left-full ml-1">0</span>
      </div>
    </div>
  );
}
