import React from "react";

interface IProps {
  ph?: number; // 0-14
  index?: number; // 0-14
}

export default function PHChart({ ph = 3, index = 3 }: IProps) {
  // TODO dont display lines when no protein is selected
  const roundedPh = Math.round(ph * 100) / 100;
  const roundedIndex = Math.round(index * 100) / 100;
  const phPosition = (ph / 14) * 100; // in percentage
  const indexPosition = (index / 14) * 100; // in percentage

  console.log(ph, roundedPh);
  console.log(index, roundedIndex);
  return (
    // Chart
    <div className="flex flex-col h-full justify-center gap-1 mx-10">
      {/* Heading */}
      <h2 className="text-gray-500 text-lg text-center">Skala PH</h2>

      {/* Scale & numbers wrapper */}
      <div className="relative w-full h-full">
        {/* Fill */}
        <div className="relative w-32 h-full bg-gradient-to-t from-red-600 via-yellow-500 to-blue-600 rounded-lg">
          {/* TODO dwie krechy */}
          {/* PH line */}
          <div
            style={{ top: `${phPosition}%` }}
            className="w-full h-1 bg-green-500 left-0 absolute transform -translate-y-1/2 transition-all duration-700"
          ></div>
          {/* Index line */}
          <div
            style={{ top: `${indexPosition}%` }}
            className="w-full h-1 bg-purple-500 left-0 absolute transform -translate-y-1/2 transition-all duration-1000"
          ></div>
        </div>
        {/* PH scale 'numbers' */}
        <span className="text-md absolute top-0  transform -translate-y-1/2 left-full ml-1">
          14
        </span>
        <span
          style={{ top: `${phPosition}%` }}
          className="text-xl font-bold right-full mr-2 absolute transform -translate-y-1/2 transition-all duration-700"
        >
          {roundedPh}
        </span>
        <span
          style={{ top: `${indexPosition}%` }}
          className="text-xl font-bold right-full mr-2 absolute transform -translate-y-1/2 transition-all duration-1000"
        >
          {roundedIndex}
        </span>
        <span className="text-md absolute bottom-0 transform translate-y-1/2 left-full ml-1">
          0
        </span>
      </div>
    </div>
  );
}
