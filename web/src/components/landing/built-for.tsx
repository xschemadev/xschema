import { headingVariants, cardVariants } from '@/components/landing/variants';
import { cn } from '@/lib/cn';

const frameworkBullets = [
  'Full type inference and autocomplete',
  'Tree-shakeable output',
  'No runtime conversion cost',
  'Schemas from files, URLs, or inline definitions',
];

const runtimeBullets = [
  'Programmatic convert() API',
  'Dynamic schema sources',
  'Per-schema conversion cost',
];

function ModeSection({
  heading,
  badge,
  body,
  bullets,
}: {
  heading: string;
  badge: string;
  body: string;
  bullets: string[];
}) {
  return (
    <div className="mt-6 rounded-xl border border-fd-border/50 bg-fd-background/50 p-4">
      <div className="flex items-center gap-2">
        <h4 className="text-sm font-semibold">{heading}</h4>
        <span className="rounded-full bg-brand/10 px-2.5 py-0.5 text-xs font-medium text-brand">
          {badge}
        </span>
      </div>
      <p className="mt-2 text-sm leading-relaxed text-fd-muted-foreground">
        {body}
      </p>
      <ul className="mt-3 space-y-1.5">
        {bullets.map((item) => (
          <li key={item} className="flex items-start gap-2 text-sm">
            <span className="mt-1.5 size-1.5 shrink-0 rounded-full bg-brand" />
            <span>{item}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}

function AudienceCard({
  title,
  bodyParagraphs,
  keyPoint,
  mode,
}: {
  title: string;
  bodyParagraphs: string[];
  keyPoint: string;
  mode: {
    heading: string;
    badge: string;
    body: string;
    bullets: string[];
  };
}) {
  return (
    <div className={cn(cardVariants(), 'flex flex-col')}>
      <h3 className="text-base font-semibold">{title}</h3>
      <div className="mt-4 space-y-3 text-sm leading-relaxed">
        {bodyParagraphs.map((p, i) => (
          <p key={i}>{p}</p>
        ))}
      </div>
      <ModeSection
        heading={mode.heading}
        badge={mode.badge}
        body={mode.body}
        bullets={mode.bullets}
      />
      <div className="mt-auto border-t border-fd-border/50 pt-4 mt-6">
        <p className="text-sm font-medium text-brand">{keyPoint}</p>
      </div>
    </div>
  );
}

export function BuiltFor({ className }: { className?: string }) {
  return (
    <section className={cn('col-span-full', className)}>
      <h2
        className={cn(
          headingVariants({ variant: 'h2' }),
          'text-brand text-center font-mono font-bold uppercase',
        )}
      >
        Built for.
      </h2>
      <div className="mt-10 grid grid-cols-1 gap-6 lg:grid-cols-2">
        <AudienceCard
          title="Developers shipping real projects."
          bodyParagraphs={[
            'You have JSON Schemas \u2014 from an API contract, an OpenAPI spec, a schema registry, or a local definition \u2014 and you need native validators in your codebase.',
            'xschema converts them to Zod, ArkType, Valibot, Pydantic, and others. Define your schemas once, run xschema generate, use the output directly. TypeScript, Python, Go, Rust \u2014 one unified workflow.',
          ]}
          keyPoint="No manual translation. No drift between schema and validator."
          mode={{
            heading: 'Framework Mode',
            badge: 'Build-time',
            body: 'Schemas are declared in config files and converted at build time. Run xschema generate to produce native validators with full static type inference. Zero runtime cost.',
            bullets: frameworkBullets,
          }}
        />
        <AudienceCard
          title="Library maintainers who need codegen."
          bodyParagraphs={[
            'You maintain a library or framework that needs to generate validators from JSON Schema, but you don\u2019t want to own the conversion logic.',
            'xschema adapters handle it. Each one is tested against the official JSON Schema Test Suite, with clear reporting of what is supported and what isn\u2019t. Use Framework Mode to generate at build time, or Runtime Mode to convert programmatically inside your library.',
          ]}
          keyPoint="Offload the conversion. Focus on your domain."
          mode={{
            heading: 'Runtime Mode',
            badge: 'Programmatic',
            body: 'Convert schemas programmatically at runtime. When codegen is part of the workflow.',
            bullets: runtimeBullets,
          }}
        />
      </div>
    </section>
  );
}
