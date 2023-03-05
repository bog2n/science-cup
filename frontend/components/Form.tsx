import React, { ChangeEvent, FormEvent, useEffect } from "react";
import { useState } from "react";
// Icons
import RightArrow from "./Icons/RightArrow";
// Components
import FileInput from "./FileInput";
import TextInput from "./TextInput";

export default function Form({ dataHandler }: any) {
  // States
  const [genomeString, setGenomeString] = useState<string>("");
  const [genomeFile, setGenomeFile] = useState<any>(null);
  // Handlers
  function submitHandler(e: any) {
    // Disable default behaviour.
    e.preventDefault();
    // Send data to index.
    dataHandler(genomeString, genomeFile);
    resetInputs();
  }

  //Effects
  useEffect(() => {
    if (genomeFile === null) return;
    // Send data to index after file is uploaded.
    dataHandler(genomeString, genomeFile);
    resetInputs();
  }, [genomeFile]);
  // Functions
  function resetInputs() {
    setGenomeFile(null);
    setGenomeString("");
  }
  return (
    <form onSubmit={submitHandler} className="flex flex-col items-center gap-7">
      {/* Text input. */}
      <TextInput
        value={genomeString}
        handleChange={(newValue: string) => setGenomeString(newValue)}
      />
      <span className="text-xl text-gray-500">lub</span>
      {/* File input. */}
      <FileInput uploadFileHandler={(newFile: any) => setGenomeFile(newFile)} />
    </form>
  );
}
