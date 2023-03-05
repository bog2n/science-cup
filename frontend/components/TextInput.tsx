import React from "react";
// Icons
import RightArrow from "./Icons/RightArrow";

interface IProps {
  value: string;
  handleChange: any;
}

export default function TextInput({ value, handleChange }: IProps) {
  return (
    <div className="relative w-[30rem]">
      <input
        value={value}
        onChange={(e: any) => handleChange(e.target.value)}
        type="text"
        placeholder="Wprowadź sekwencję DNA/RNA..."
        className="block w-full pl-8 pr-24 py-5 rounded-full bg-gray-50"
      />
      {/* Arrow button. */}
      <button className="flex justify-center items-center h-20 aspect-square  absolute -right-2 top-1/2 rounded-full text-white bg-[#A2D2FF] transform -translate-y-1/2 hover:scale-110 transition-transform">
        {/* Arrow icon */}
        <RightArrow classes="w-8 h-8" />
      </button>
    </div>
  );
}
