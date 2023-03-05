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
      <div className="h-96 overflow-x-scroll overflow-y-hidden text-center justify-center">
        {protein && <img className="max-w-max" alt="Schemat białka" src={protein} />}
        {!protein && <span className="relative top-40 text-hg">Wprowadź sekwencje genomu...</span>}
      </div>
    </Card>
  );
}
