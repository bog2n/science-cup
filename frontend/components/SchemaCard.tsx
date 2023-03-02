import React from "react";
import handler from "@/pages/api/data";
import Card from "./Card";
import DummySchema from "/public/dummyschema.png";
import Image from "next/image";

export default function SchemaCard() {
  return (
    <Card>
      <div className="h-24 overflow-x-scroll">
        <div className="h-full w-[100rem] bg-pink-300"></div>
      </div>
    </Card>
  );
}
