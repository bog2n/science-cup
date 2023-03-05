import React from "react";
// Types
import { ReactNode } from "react";

interface IProps {
  children: ReactNode;
}

export default function Card({ children }: IProps) {
  return (
    <div className="flex flex-col justify-start gap-2 px-7 pb-9 pt-7 rounded-lg bg-gray-50">
      {children}
    </div>
  );
}
