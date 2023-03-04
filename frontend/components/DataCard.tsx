import { spawn } from "child_process";
import React from "react";

// Components
import Card from "./Card";

interface IProps {
  title: string;
  data?: undefined | number;
  unit?: string;
}

export default function DataCard({ title, data, unit }: IProps) {
  // When protein is not selected.
  if (data === undefined) {
    return (
      <Card>
        <h2 className="text-gray-500 text-lg">{title}</h2>
        <span>-</span>
      </Card>
    );
  }

  // When protein is selected.
  const roundedData = Math.round(data * 100) / 100;

  // Add + or - prefix for polarization.
  let prefix = "";
  title === "Polaryzacja" && roundedData > 0 ? (prefix = "+") : (prefix = "");

  return (
    <Card>
      <h2 className="text-gray-500 text-lg">{title}</h2>

      <span className="text-3xl text-black font-bold">{`${prefix}${roundedData}
          ${unit ? ` ${unit}` : ""}`}</span>
    </Card>
  );
}
