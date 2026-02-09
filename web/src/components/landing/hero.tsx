"use client";

import { useEffect, useState } from "react";
import { useTheme } from "next-themes";
import { motion } from "motion/react";
import { Dithering, ImageDithering } from "@paper-design/shaders-react";
import { useIsMobile } from "@/hooks/useMobile";
import GorillaImage from "@/assets/gorilla.svg";

const shaderFade = {
  initial: { opacity: 0 },
  animate: { opacity: 1, transition: { duration: 0.4, ease: "easeOut" as const } },
};

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
        <motion.div className="absolute size-full" {...shaderFade}>
          <Dithering
            colorBack="#00000000"
            colorFront={resolvedTheme === "dark" ? "#22231e" : "#ffeeb2"}
            shape="simplex"
            type="4x4"
            size={4}
            speed={0.1}
            className="absolute size-full"
            minPixelRatio={1}
          />
        </motion.div>
      )}
      {showShaders && (
        <motion.div
          className="absolute bottom-0 right-0 sm:bottom-4 sm:right-4 lg:right-20"
          {...shaderFade}
        >
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
            style={{
              opacity: isMediumScreen && resolvedTheme === "light" ? 0.3 : 1,
            }}
            speed={0}
            frame={5000 * 120}
            minPixelRatio={1}
          />
        </motion.div>
      )}
    </>
  );
}
