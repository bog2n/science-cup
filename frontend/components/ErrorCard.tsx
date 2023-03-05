import React from "react";

export default function ErrorCard({ error }: any) {
  return (
    <div className={`ease-in fixed bg-slate-400 rounded-full px-5 top-4 left-4 p-3 duration-150 text-red-500 text-lg font-semibold ${error ? "opacity-100" : "opacity-0"}`}>
      {error}
    </div>
  );
}
