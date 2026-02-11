import { headingVariants, cardVariants } from "@/components/landing/variants";
import { cn } from "@/lib/cn";
import { motion } from "motion/react";
import { PlusCard } from "../ui/plus-card";
import {
  fadeUp,
  slideFromLeft,
  slideFromRight,
  defaultViewport,
} from "@/components/landing/animation-variants";

const frameworkBullets = [
  "Full type inference and autocomplete",
  "Schemas from files, URLs, or inline definitions",
  "No runtime conversion cost",
];

const runtimeBullets = [
  "Programmatic convert() API",
  "Dynamic schema sources",
  "Per-schema conversion cost",
];

function AudienceCard({
  title,
  bodyParagraphs,
  mode,
}: {
  title: string;
  bodyParagraphs: string[];
  mode: {
    heading: string;
    badge: string;
    body: string;
    bullets: string[];
  };
}) {
  return (
    <div className={cn(cardVariants(), "flex flex-col p-4 sm:p-6")}>
      <h3 className="text-2xl text-center font-semibold">{title}</h3>
      <div className="mt-6 space-y-3 text-fd-muted-foreground">
        {bodyParagraphs.map((p, i) => (
          <p
            key={i}
            className="text-balance font-light border-dotted border-l-2 border-brand-200 pl-4 italic"
          >
            {p}
          </p>
        ))}
      </div>
      <div className="mt-6 h-full rounded-xl border border-fd-border/50 bg-fd-background/50 py-4 px-6">
        <div className="flex items-center gap-3 flex-wrap font-mono tracking-tight">
          <h4 className="text-lg font-semibold">{mode.heading}</h4>
          <span className="rounded-full bg-brand/10 px-2.5 py-0.5 text-xs font-medium text-brand">
            {mode.badge}
          </span>
        </div>
        <p className="mt-2 text-sm leading-relaxed text-fd-muted-foreground">
          {mode.body}
        </p>
        <ul className="mt-3 space-y-1.5">
          {mode.bullets.map((item) => (
            <li key={item} className="flex items-start gap-2 text-sm">
              <span className="mt-1.5 size-1.5 shrink-0 rounded-full bg-brand" />
              <span>{item}</span>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export function BuiltFor({ className }: { className?: string }) {
  return (
    <section className={cn(className)}>
      <motion.h2
        className={cn(
          headingVariants({ variant: "h2" }),
          "text-brand text-center font-mono font-bold uppercase",
        )}
        variants={fadeUp}
        initial="hidden"
        whileInView="visible"
        viewport={defaultViewport}
      >
        Built for.
      </motion.h2>
      <PlusCard className="mt-10 grid grid-cols-1 gap-6 lg:grid-cols-2 rounded-5xl p-0 sm:p-6 max-sm:border-none">
        <motion.div
          className="h-full"
          variants={slideFromLeft}
          initial="hidden"
          whileInView="visible"
          viewport={defaultViewport}
        >
          <AudienceCard
          title="Developers shipping real projects."
          bodyParagraphs={[
            "You have JSON Schemas \u2014 from an API contract, an OpenAPI spec, a schema registry, or a local definition \u2014 and you need native validators in your codebase.",
          ]}
          mode={{
            heading: "Framework Mode",
            badge: "Build-time",
            body: "Schemas are declared in config files and converted at build time.",
            bullets: frameworkBullets,
          }}
        />
        </motion.div>
        <motion.div
          className="h-full"
          variants={slideFromRight}
          initial="hidden"
          whileInView="visible"
          viewport={defaultViewport}
        >
          <AudienceCard
            title="Library maintainers who need codegen."
            bodyParagraphs={[
              "You maintain a library or framework that needs to generate validators from JSON Schema programmatically, but you don\u2019t want to own the conversion logic.",
            ]}
            mode={{
              heading: "Runtime Mode",
              badge: "Programmatic",
              body: "Convert schemas programmatically at runtime. When codegen is part of the workflow.",
              bullets: runtimeBullets,
            }}
          />
        </motion.div>
      </PlusCard>
    </section>
  );
}
