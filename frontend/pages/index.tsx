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
import ErrorCard from "@/components/ErrorCard";

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
  const [schema, setSchema] = useState<null | string>(null);
  const [errorValue, setErrorValue] = useState<null | string>(null);

  const [listIsOpen, setListIsOpen] = useState(false);

  function processData(genome: string, file: File) {
    let requestParameters;

    if (file == null) {
      requestParameters = {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: "genome=" + genome,
      };
    } else {
      let data = new FormData();
      data.append("file", file);
      requestParameters = {
        method: "POST",
        body: data,
      };
    }

    fetch("/api/data", requestParameters)
      .then((response) => response.json())
      .then((data) => {
        if (data.ok) {
          setProteinsData(data);
          setCurrentProtein(data.proteins[0]);
          setSchema("/api/image/" + data.proteins[0].protein);
        } else {
          setProteinsData(null);
          setCurrentProtein(null);
          setSchema(null);
          setErrorValue("Nie znaleziono białek");
          setTimeout(() => {
            setErrorValue(null);
          }, 3000);
        }
      })
      .catch((error) => {
        setErrorValue("Nie udało nawiązać połączenia z serwerem");
        setTimeout(() => {
          setErrorValue(null);
        }, 3000);
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
    setSchema("/api/image/" + newProtein.protein);
  }

  return (
    <main className="max-w-6xl mx-auto px-16">
      {/* Main header. */}
      <header className="py-24">
        <h1 className="text-center max-w-3xl mx-auto">DNAnalyzer</h1>
      </header>
      <ErrorCard error={errorValue} />
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
            className="block fixed inset-0 z-40"
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
                : proteinsData.proteins.map((protein: IProtein, index) => {
                    if (protein === currentProtein) {
                      return index + 1 + ". " + currentProtein.protein;
                    }
                  })}
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
            className={`absolute top-full mt-3 flex flex-col justify-start bg-gray-50 max-w-2xl max-h-96 overflow-y-scroll rounded-lg transform transition-all duration-300 ${
              !listIsOpen
                ? "-translate-y-10 opacity-0 pointer-events-none"
                : "translate-y-0 opacity-100 pointer-events-auto shadow-md"
            }`}
          >
            {/* List of the proteins in each protein. */}
            {proteinsData &&
              proteinsData.proteins.map((protein: IProtein, index) => {
                return (
                  <li key={index}>
                    <button
                      className={`grid grid-cols-[auto_1fr] gap-2 px-5 py-3 w-full break-all text-left hover:bg-gray-300 ${
                        protein.protein === currentProtein?.protein
                          ? "bg-gray-200"
                          : ""
                      }`}
                      onClick={() => changeProteinHandler(protein)}
                    >
                      <span>{`${index + 1}.`}</span>
                      <span>{protein.protein}</span>
                    </button>
                  </li>
                );
              })}
          </ul>
        </div>
        <div className="flex mb-8">
          {/* Data */}
          <div className="grid grid-cols-2 grid-rows-2 gap-3 flex-grow">
            <DataCard
              title="Masa"
              data={currentProtein?.mass}
              unit="u"
            />
            <DataCard
              title="Indeks Hydrofobowy"
              data={currentProtein?.hindex}
              unit="kcal/mol"
            />
            <DataCard title="Polaryzacja" data={currentProtein?.polarity} />
            {/* PH chart legend */}
            <div className="flex flex-col gap-3 p-7 justify-end">
              <div className="flex justify-start items-center gap-3">
                <span className="block h-1 w-12 bg-green-400"></span>{" "}
                <span>PH białka</span>
              </div>
              <div className="flex justify-start items-center gap-3">
                <span className="block h-1 w-12 bg-purple-400"></span>{" "}
                <span>Punkt izoelektryczny</span>
              </div>
            </div>
          </div>
          {/* PH chart */}
          <div className="pl-16 pr-0">
            <PHChart ph={currentProtein?.ph} index={currentProtein?.isopoint} />
          </div>
        </div>
        {/* Schema */}
        <div>
          <SchemaCard protein={schema} />
        </div>
      </section>
    </main>
  );
}
