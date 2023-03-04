import Head from "next/head";
import Image from "next/image";
import { Inter } from "@next/font/google";
import { useState } from "react";

// Components
import Form from "../components/Form";
import DataCard from "@/components/DataCard";
import SchemaCard from "@/components/SchemaCard";
import PHChart from "@/components/PHChart";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  const [fill, setFill] = useState(3);

  // TODO remove in near future
  function changeFillHandler() {
    setFill(Math.round(Math.random() * 14));
  }

  function processData(data:any) {
    let genome = data.target[0].value;
    let file = data.target[2].files[0];
    let requestParameters;

    if (genome == "") {
      let data = new FormData();
      data.append("file", file);
      requestParameters = {
        method: "POST",
        body: data
      };
    } else {
      requestParameters = {
        method: "POST",
        headers: {'Content-Type': 'application/x-www-form-urlencoded'},
        body: 'genome=' + genome
      };
    }

    fetch('/api/data', requestParameters)
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      // TODO handle this
    })
    .catch((error) => {
      // TODO handle this
    });
  }

  return (
    <main className="max-w-7xl mx-auto fir">
      {/* Main header. */}
      <header className="py-32">
        <h1 className="text-center max-w-3xl mx-auto">
          Potężna aplikacja webowa.
        </h1>
      </header>
      {/* Insert data to calculate. */}
      <section className="flex justify-center items-center mb-36">
        <Form dataHandler={processData} />
      </section>
      {/* Printed data. */}
      <section className="mb-10">
        {/* Data */}
        <div className="grid grid-cols-3 gap-10">
          <DataCard title="Masa" data="10" />
          <DataCard title="Index Hydrofobowy" data="10" />
          <DataCard title="Polarność" data="10" />
        </div>
        {/* PH chart */}
        {/* TODO dwie krechy... */}
        <div className="my-10">
          <button onClick={changeFillHandler}>Change fill</button>
          <PHChart fill={fill} />
        </div>
        {/* Schema */}
        <div>
          <SchemaCard />
        </div>
      </section>
    </main>
  );
}
