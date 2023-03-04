import React from "react";
import handler from "@/pages/api/data";

// Components
import Card from "./Card";
import Image from "next/image";

export default function SchemaCard() {
  return (
    <Card>
      {/* Heading */}
      <h2 className="text-gray-500 text-lg text-center">Schemat</h2>
      <div className="relative h-24 overflow-x-scroll">
        zdjęcie białka czy tam cczegoś 
      </div>
    </Card>
  );
}
