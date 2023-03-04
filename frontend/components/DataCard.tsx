import React from "react";

// Components
import Card from "./Card";

interface IProps {
  title?: string;
  data?: number;
  unit?: string;
}

export default function DataCard({
  title = "Heading",
  data = 10,
  unit,
}: IProps) {
  const roundedData = Math.round(data * 100) / 100;
  let prefix = "";
  if (title == "Polaryzacja" && roundedData > 0) {
    prefix = "+";
  }
  return (
    <Card>
      <h2 className="text-gray-500 text-lg">{title}</h2>

      {unit ? (
        <span className="text-3xl text-black font-bold">{`${prefix}${roundedData} ${unit}`}</span>
      ) : (
        <span className="text-3xl text-black font-bold">{`${prefix}${roundedData}`}</span>
      )}
    </Card>
  );
}
