import "@/styles/globals.css";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  const request = new Request("");
  // const requestHeaders = new Headers(Request.headers);

  // Add new request headers
  //console.log(request);
  return <Component {...pageProps} />;
}
