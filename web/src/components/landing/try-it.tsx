import { Link } from '@tanstack/react-router';
import { cn } from '@/lib/cn';
import { headingVariants, buttonVariants, cardVariants } from '@/components/landing/variants';
import { XSCHEMA_GITHUB_URL } from '@/lib/constants';

const exampleSchema = `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "id": { "type": "string", "format": "uuid" },
    "name": { "type": "string", "minLength": 1 },
    "email": { "type": "string", "format": "email" },
    "role": { "enum": ["admin", "user", "viewer"] }
  },
  "required": ["id", "name", "email"]
}`;

export function TryIt() {
  return (
    <div className="col-span-full flex flex-col items-center text-center gap-6">
      <h2
        className={cn(
          headingVariants({ variant: 'h2' }),
          'text-brand font-mono font-bold uppercase',
        )}
      >
        Try it.
      </h2>
      <p className="text-lg text-fd-muted-foreground max-w-xl">
        Write a schema. Pick an adapter. See the output.
      </p>
      <div
        className={cn(
          cardVariants(),
          'w-full max-w-2xl text-left overflow-hidden',
        )}
      >
        <p className="text-xs font-mono text-fd-muted-foreground mb-3">
          Example JSON Schema
        </p>
        <pre className="overflow-x-auto rounded-lg border border-fd-border bg-fd-background p-4 text-xs font-mono text-fd-foreground leading-relaxed">
          <code>{exampleSchema}</code>
        </pre>
        <p className="mt-4 text-xs text-fd-muted-foreground italic">
          StackBlitz embed coming soon.
        </p>
      </div>
      <div className="flex flex-row items-center gap-4 flex-wrap justify-center">
        <Link
          to="/docs/$"
          className={cn(buttonVariants(), 'max-sm:text-sm')}
        >
          Read the docs
        </Link>
        <a
          href={XSCHEMA_GITHUB_URL}
          target="_blank"
          rel="noreferrer noopener"
          className={cn(
            buttonVariants({ variant: 'secondary' }),
            'max-sm:text-sm',
          )}
        >
          View on GitHub
        </a>
      </div>
    </div>
  );
}
