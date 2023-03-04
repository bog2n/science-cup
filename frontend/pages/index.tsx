import Head from "next/head";
import Image from "next/image";
import { Inter } from "@next/font/google";
import { useState } from "react";

// Components
import Form from "../components/Form";
import DataCard from "@/components/DataCard";
import SchemaCard from "@/components/SchemaCard";
import PHChart from "@/components/PHChart";
import RightArrow from "@/components/Icons/RightArrow";

const inter = Inter({ subsets: ["latin"] });

interface IProtein {
  hindex: number;
  isopoint: number;
  mass: number;
  ph: number;
  polarity: number;
  protein: string;
}
interface IProteinsData {
  ok: boolean;
  proteins: IProtein[];
}

export default function Home() {
  //* States
  // Each protein data.
  const [proteinsData, setProteinsData] = useState<null | IProteinsData>(null);
  const [currentProtein, setCurrentProtein] = useState<null | IProtein>(null);

  const [listIsOpen, setListIsOpen] = useState(false);

  function processData(genome: string, file: File) {
    console.log("from index,", genome, file);
    let requestParameters;

    if (genome == "") {
      let data = new FormData();
      data.append("file", file);
      requestParameters = {
        method: "POST",
        body: data,
      };
    } else {
      requestParameters = {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: "genome=" + genome,
      };
    }

    fetch("/api/data", requestParameters)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        // TODO handle this
        setProteinsData(data);
        setCurrentProtein(data.proteins[0]);
      })
      .catch((error) => {
        // TODO handle this
      });
  }

  //* Handlers
  function listOpenHandler() {
    setListIsOpen((prev) => !prev);
  }
  function changeProteinHandler(newProtein: IProtein) {
    setListIsOpen(false);
    // Update the protein which is going to be displayed.
    setCurrentProtein(newProtein);
  }

  return (
    <main className="max-w-7xl mx-auto px-16">
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
        {/* Whole screen overlay */}
        {listIsOpen && (
          <div
            onClick={() => setListIsOpen(false)}
            className="block absolute inset-0 z-40"
          ></div>
        )}
        {/* List with all proteins. */}
        <div className="relative z-50 mb-4">
          <button
            onClick={listOpenHandler}
            disabled={!proteinsData}
            className="text-left max-w-sm text-lg font-bold flex justify-center items-center gap-2 disabled:pointer-events-none disabled:text-gray-400"
          >
            <span className="truncate">
              {!proteinsData
                ? "Wprowadź sekwencję DNA"
                : currentProtein?.protein}
            </span>{" "}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="2"
              stroke="currentColor"
              className={`block w-12 h-5 transform transition-transform ${
                !listIsOpen ? "rotate-0" : "rotate-90"
              }`}
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M8.25 4.5l7.5 7.5-7.5 7.5"
              />
            </svg>
          </button>
          <ul
            className={`absolute top-full mt-3 flex flex-col justify-start gap-2 bg-gray-50 shadow-lg max-w-2xl max-h-52 overflow-y-scroll rounded-lg transform transition-all duration-300 ${
              !listIsOpen
                ? "-translate-y-10 opacity-0 pointer-events-none"
                : "translate-y-0 opacity-100 pointer-events-auto"
            }`}
          >
            {/* List of the proteins in each protein. */}
            {proteinsData &&
              proteinsData.proteins.map((protein: IProtein, index) => {
                return (
                  <li key={index}>
                    <button
                      className={`block px-5 py-2 w-full break-all text-left hover:bg-gray-300 ${
                        protein.protein === currentProtein?.protein
                          ? "bg-gray-200"
                          : ""
                      }`}
                      onClick={() => changeProteinHandler(protein)}
                    >
                      {protein.protein}
                    </button>
                  </li>
                );
              })}
          </ul>
        </div>
        <div className="flex ">
          {/* Data */}
          <div className="grid grid-cols-2 grid-rows-2 gap-4 flex-grow">
            <DataCard title="Mass" data={currentProtein?.mass} unit="U" />
            <DataCard title="Isopoint" data={currentProtein?.isopoint} />
            <DataCard title="Polarity" data={currentProtein?.polarity} />
            {/* PH chart legend */}
            <div className="flex flex-col gap-3 justify-end pl-7 pt-7 pr-7">
              <div className="flex justify-start items-center gap-3">
                <span className="block h-1 w-12 bg-green-400"></span>{" "}
                <span>PH białka.</span>
              </div>
              <div className="flex justify-start items-center gap-3">
                <span className="block h-1 w-12 bg-purple-400"></span>{" "}
                <span>Punkt izo coś tam.</span>
              </div>
            </div>
          </div>
          {/* PH chart */}
          <div className="">
            <PHChart ph={currentProtein?.ph} index={currentProtein?.isopoint} />
          </div>
        </div>
        {/* Schema */}
        <div>
          <SchemaCard />
        </div>
      </section>
    </main>
  );
}
