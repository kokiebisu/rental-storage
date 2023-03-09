import { MantineProvider } from "@mantine/core";
import type { AppProps } from "next/app";

import { ContextProvider } from "@/context";
import "@/styles/globals.css";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <MantineProvider
      withGlobalStyles
      withNormalizeCSS
      theme={{
        colorScheme: "light",
      }}
    >
      <ContextProvider>
        <Component {...pageProps} />
      </ContextProvider>
    </MantineProvider>
  );
}
