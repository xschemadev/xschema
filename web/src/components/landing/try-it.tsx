import { cn } from "@/lib/cn";
import { headingVariants } from "@/components/landing/variants";
import { motion } from "motion/react";
import { BorderBeam } from "../ui/border-beam";
import { CliAnimation } from "./cli-animation";
import { Arrow } from "../ui/arrow";
import {
  fadeUp,
  slideFromLeft,
  slideFromRight,
  defaultViewport,
} from "@/components/landing/animation-variants";

const notesSchemaConfig = `{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "Note",
      "sourceType": "url",
      "source": "https://your.schemastore.org/note.json",
      "adapter": "@xschemadev/zod"
    },
    {
      "id": "Tag",
      "sourceType": "json",
      "source": {
        "type": "string",
        "enum": ["work", "hobby", "other"]
      },
      "adapter": "@xschemadev/zod"
    }
  ]
}`;

const usersSchemaConfig = `{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "User",
      "sourceType": "file",
      "source": "./user.schema.json"
      "adapter": "@xschemadev/effect"
    }
  ]
}`;

export function TryIt({ className }: { className?: string }) {
  return (
    <>
      <motion.div
        className="col-span-full"
        variants={fadeUp}
        initial="hidden"
        whileInView="visible"
        viewport={defaultViewport}
      >
        <h2
          className={cn(
            headingVariants({ variant: "h2" }),
            "text-brand font-mono font-bold uppercase",
          )}
        >
          Try it.
        </h2>
        <p className="text-lg text-fd-muted-foreground max-w-xl">
          Write a schema. Pick an adapter. See the output.
        </p>
      </motion.div>
      <motion.div
        className="flex flex-col gap-4"
        variants={slideFromLeft}
        initial="hidden"
        whileInView="visible"
        viewport={defaultViewport}
      >
        <pre className="overflow-x-auto border border-dashed border-brand max-w-fit p-4 text-xs font-mono text-fd-foreground">
          <code>{usersSchemaConfig}</code>
        </pre>
        <pre className="overflow-x-auto border border-dashed border-brand-secondary max-w-fit p-4 text-xs font-mono text-fd-foreground">
          <code>{notesSchemaConfig}</code>
        </pre>
      </motion.div>
      <motion.div
        className="flex flex-col gap-8 relative max-w-3xl w-full mx-auto h-fit min-h-32 justify-between overflow-visible rounded-3xl border border-dashed p-2"
        variants={slideFromRight}
        initial="hidden"
        whileInView="visible"
        viewport={defaultViewport}
      >
        <BorderBeam
          colorFrom="var(--color-brand)"
          colorTo="var(--color-brand-200)"
          size={100}
          borderWidth={2}
          duration={15}
          className="absolute z-40"
        />
        <div className="relative">
          <Arrow
            padding={40}
            className="absolute z-2"
            fromX={150}
            toX={170}
            curveX={180}
            fromY={1}
            toY={45}
            curveY={-50}
            strokeWidth={3}
            tipSize={10}
            color="#c6bb58"
          />
          <CliAnimation className="w-full" />
        </div>
      </motion.div>

      <p className="text-xs text-fd-muted-foreground italic">
        StackBlitz embed coming soon.
      </p>
    </>
  );
}
