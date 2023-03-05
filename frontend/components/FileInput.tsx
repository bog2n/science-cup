import React from "react";
// Icons
import UploadIcon from "./Icons/UploadIcon";

interface IProps {
  uploadFileHandler: any;
}

export default function FileInput({ uploadFileHandler }: IProps) {
  return (
    <div className="bg-gray-50 relative w-full ease-linear duration-100 hover:bg-gray-200 hover:border-gray-300 focus-within:bg-gray-300 focus-within:border-gray-300 flex flex-col items-center  text-gray-400 py-6 px-12 rounded-lg border-2 border-dashed">
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
