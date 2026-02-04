"use client";

import { useEffect, useState } from "react";
import { useTheme } from "next-themes";
import { Dithering, ImageDithering } from "@paper-design/shaders-react";
import { useIsMobile } from "@/hooks/useMobile";
import GorillaImage from "@/assets/gorilla.svg";

export function Hero() {
  const { resolvedTheme } = useTheme();
  const [showShaders, setShowShaders] = useState(false);
  const isMobile = useIsMobile();
  const isMediumScreen = useIsMobile(2000);

  useEffect(() => {
    // apply some delay, otherwise on slower devices, it errors with uniform images not being fully loaded.
    setTimeout(() => {
      setShowShaders(true);
    }, 400);
  }, []);

  return (
    <>
      {showShaders && (
        <Dithering
          colorBack="#00000000"
          colorFront={resolvedTheme === "dark" ? "#22231e" : "#ffeeb2"}
          shape="simplex"
          type="4x4"
          size={4}
          speed={0.1}
          className="absolute animate-fd-fade-in duration-400 size-full"
          minPixelRatio={1}
        />
      )}
      {showShaders && (
        <ImageDithering
          image={GorillaImage}
          width={127 * (isMobile ? 2.8 : isMediumScreen ? 3.5 : 4)}
          height={162 * (isMobile ? 2.8 : isMediumScreen ? 3.5 : 4)}
          colorBack="#00000000"
          colorFront={resolvedTheme === "dark" ? "#6b4b3e" : "#fa8023"}
          originalColors={resolvedTheme === "light"}
          type="8x8"
          scale={1}
          size={4}
          speed={0}
          frame={5000 * 120}
          className="absolute animate-fd-fade-in duration-400 bottom-0 right-0 sm:bottom-4 sm:right-4 lg:right-20"
          minPixelRatio={1}
        />
      )}
    </>
  );
}
