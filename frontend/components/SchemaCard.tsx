import React from "react";
import handler from "@/pages/api/data";

// Components
import Card from "./Card";
import Image from "next/image";

export default function SchemaCard({protein}:any) {
  return (
    <Card>
      {/* Heading */}
      <h2 className="text-gray-500 text-lg text-center">Schemat</h2>
      <div className="relative overflow-x-scroll">
        {protein && <img alt="Protein schematic image" src={protein} />}
      </div>
    </Card>
  );
}
