import { Link } from "@tanstack/react-router";
import { cn } from "@/lib/cn";
import { motion } from "motion/react";
import { headingVariants, buttonVariants } from "@/components/landing/variants";
import { XSCHEMA_GITHUB_URL } from "@/lib/constants";
import GorillaImage from "@/assets/gorilla-square.svg";
import { fadeUp, defaultViewport } from "@/components/landing/animation-variants";

export function OpenSourceFooter({ className }: { className?: string }) {
  return (
    <motion.section
      className={cn(
        "flex flex-col md:flex-row items-center gap-8 justify-evenly text-center overflow-y-hidden",
        className,
      )}
      variants={fadeUp}
      initial="hidden"
      whileInView="visible"
      viewport={defaultViewport}
    >
      <div>
        <h2
          className={cn(
            headingVariants({ variant: "h2" }),
            "text-brand font-mono font-bold uppercase text-balance",
          )}
        >
          Open source. Open process.
        </h2>
        <p className="mt-4 max-w-xl leading-relaxed text-fd-muted-foreground text-balance">
          xschema is open source. Contributions, adapter proposals and bug
          reports are greatly appreciated.
        </p>
        <div className="mt-8 flex flex-row items-center gap-4 flex-wrap justify-center">
          <Link to="/docs/$" className={cn(buttonVariants(), "max-sm:text-sm")}>
            Read the docs
          </Link>
          <a
            href={XSCHEMA_GITHUB_URL}
            target="_blank"
            rel="noreferrer noopener"
            className={cn(
              buttonVariants({ variant: "secondary" }),
              "max-sm:text-sm",
            )}
          >
            View on GitHub
          </a>
        </div>
      </div>
      <img
        src={GorillaImage}
        alt="Gorilla"
        className="object-cover h-40 md:h-80 rounded-4xl"
      />
    </motion.section>
  );
}
