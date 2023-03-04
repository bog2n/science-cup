import { spawn } from "child_process";
import React from "react";

// Components
import Card from "./Card";

interface IProps {
  title: string;
  data?: null | number;
  unit?: string;
}

export default function DataCard({ title, data, unit }: IProps) {
  let roundedData = null;
  if (data) {
    roundedData = Math.round(data * 100) / 100;
  }
  return (
    <Card>
      <h2 className="text-gray-500 text-lg">{title}</h2>

      {!data && <span>-</span>}

      {data && (
        <span className="text-3xl text-black font-bold">{`${roundedData}${
          unit ? ` ${unit}` : ""
        }`}</span>
      )}
    </Card>
  );
}
