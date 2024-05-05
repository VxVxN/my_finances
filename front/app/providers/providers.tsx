"use client";

import { NextUIProvider } from "@nextui-org/system";
import { useRouter } from "next/navigation";
import React, { FC } from "react";
import { ThemeProvider as NextThemesProvider } from "next-themes";
import { ThemeProviderProps } from "next-themes/dist/types";

interface ProvidersProps {
  children: React.ReactNode;
  themeProps?: any;
}

const Providers: FC<ProvidersProps> = ({ children, themeProps }) => {
  const router = useRouter();

  return (
    <NextUIProvider navigate={router.push}>
      <NextThemesProvider {...themeProps}>
        {children}
        </NextThemesProvider>
    </NextUIProvider>
  );
};

export default Providers;
