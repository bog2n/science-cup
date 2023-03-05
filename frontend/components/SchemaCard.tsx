import React from "react";

// Components
import Card from "./Card";

export default function SchemaCard({ protein }: any) {
  return (
    <Card>
      {/* Heading */}
      <h2 className="text-center">Schemat</h2>
      <div className="h-96 overflow-x-scroll overflow-y-hidden text-center justify-center">
        {protein && (
          <img className="max-w-max" alt="Schemat białka" src={protein} />
        )}
        {!protein && (
          <span className="relative top-40 text-lg">
            Wprowadź sekwencje genomu...
          </span>
        )}
      </div>
    </Card>
  );
}
