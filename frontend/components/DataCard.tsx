import React from "react";

// Components
import Card from "./Card";

interface IProps {
  title: string;
  data?: undefined | number;
  unit?: string;
}

export default function DataCard({ title, data, unit }: IProps) {
  //* When protein is not selected.
  if (data === undefined) {
    return (
      <Card>
        <h2 className="">{title}</h2>
        <span>-</span>
      </Card>
    );
  }

  //* When protein is selected.
  // Round data to 2 decimal places.
  const roundedData = Math.round(data * 100) / 100;

  // Add + or - prefix for polarization.
  let prefix = "";
  title === "Polaryzacja" && roundedData > 0 ? (prefix = "+") : (prefix = "");

  return (
    <Card>
      <h2 className="">{title}</h2>

      <span className="data-text">{`${prefix}${roundedData}
          ${unit ? ` ${unit}` : ""}`}</span>
    </Card>
  );
}
