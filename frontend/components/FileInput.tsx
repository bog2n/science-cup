import React from "react";
import UploadIcon from "./Icons/UploadIcon";

interface IProps {
  uploadFileHandler: any;
}

export default function FileInput({ uploadFileHandler }: IProps) {
  return (
    <div className="relative w-full hover:bg-gray-50 hover:border-gray-300 focus-within:bg-gray-50 focus-within:border-gray-300 flex flex-col items-center  text-gray-400 py-6 px-12 rounded-lg border-2 border-dashed">
      <UploadIcon />
      <span>Załaduj sekwencję z pliku.</span>

      <input
        className="absolute top-0 left-0 rounded-lg block w-full h-full opacity-0 cursor-pointer"
        onChange={(e: any) => uploadFileHandler(e.target.files[0])}
        type="file"
        value=""
      />
    </div>
  );
}
