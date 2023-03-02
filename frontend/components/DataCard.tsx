import React from "react";

// Components
import Card from "./Card";

interface IProps {
  title?: string;
  data?: string | number;
}

export default function DataCard({ title = "Heading", data = 10 }: IProps) {
  return (
    <Card>
      <h3 className="text-gray-500 text-lg">{title}</h3>
      <span className="text-3xl text-black font-bold">{data}</span>
    </Card>
  );
}
