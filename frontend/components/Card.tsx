import React from "react";

interface IProps {
  children: React.ReactNode;
}

export default function Card({ children }: IProps) {
  return (
    <div className="flex flex-col justify-start gap-1 p-7 pt-5 rounded-lg shadow-md shadow-gray-100 bg-gray-50">
      {children}
    </div>
  );
}
