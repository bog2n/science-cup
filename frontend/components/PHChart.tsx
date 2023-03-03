import React from "react";

interface IProps {
  fill?: number; // 0-14
}

export default function PHChart({ fill = 3 }: IProps) {
  const fillPercentage = fill / 14;
  return (
    // Chart
    <div className="relative">
      {/* Heading */}
      <h2 className="text-gray-500 text-lg text-center">Skala PH</h2>

      {/* Fill container */}
      <div className="relative w-full h-36 bg-gradient-to-r from-red-600 to-blue-600 rounded-lg overflow-x-hidden">
        {/* Filled */}
        <div
          style={{ transform: `scaleX(${1 - fillPercentage})` }}
          className={`absolute h-full w-full bg-gray-300 transform origin-right  transition-all duration-[1500ms]`}
        ></div>
      </div>
      {/* PH scale 'numbers' */}
      <span className="text-xl absolute top-full left-0">1</span>
      <span
        style={{ left: `${fillPercentage * 100}%` }}
        className="text-2xl font-bold absolute top-full  transform -translate-x-1/2 transition-all duration-[1500ms]"
      >
        {fill}
      </span>
      <span className="text-xl absolute top-full right-0">14</span>
    </div>
  );
}
