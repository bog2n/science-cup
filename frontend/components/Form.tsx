import React from "react";
import RightArrow from "./Icons/RightArrow";
import UploadIcon from "./Icons/UploadIcon";

export default function Form({dataHandler}:any) {
  // Handlers
  function submitHandler(e:any) {
	e.preventDefault();
	dataHandler(e);
  }
  return (
    <form
      onSubmit={submitHandler}
      className="flex flex-col items-center gap-7"
    >
      {/* Text input. */}
      <div className="relative w-[30rem]">
        <input
          type="text"
          placeholder="Wprowadź sekwencję DNA..."
          className="block w-full pl-8 pr-24 py-5 rounded-full  border-gray-300 border-2 "
        />
        {/* Arrow button. */}
        <button className="flex justify-center items-center h-20 aspect-square  absolute -right-2 top-1/2 rounded-full text-white bg-[#A2D2FF] transform -translate-y-1/2 hover:scale-110 transition-transform shadow-lg shadow-[#A2D2FF]">
          {/* Arrow icon */}
          <RightArrow classes="w-8 h-8" />
        </button>
      </div>
      <span className="text-xl text-gray-500">lub</span>
      {/* File input. */}
      <div className="flex items-center justify-center">
        {/* Load file custom input. */}
        <input
          id="dropzone-file"
          type="file"
          className="absolute -left-[100rem] opacity-0 peer"
        />
        <label
          htmlFor="dropzone-file"
          className="flex flex-col items-center justify-center w-full border-dashed border-2 border-gray-300 rounded-lg cursor-pointer  peer-focus-within:bg-gray-50  hover:bg-gray-50 px-10 py-5"
        >
          <div className="flex flex-col items-center justify-center text-gray-300">
            <UploadIcon />
            <p className=" text-sm text-gray-500">
              <span className="font-semibold">Kliknij lub upuść</span>, żeby
              załadować plik.
            </p>
          </div>
        </label>
      </div>
    </form>
  );
}
