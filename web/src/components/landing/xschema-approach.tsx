import { Check, X } from 'lucide-react';
import { headingVariants, cardVariants } from '@/components/landing/variants';
import { cn } from '@/lib/cn';
import {
  ComparisonGrid,
  type ComparisonRow,
} from '@/components/ui/comparison-grid';

const checkIcon = <Check className="size-4 text-brand" />;
const crossIcon = <X className="size-4 text-fd-muted-foreground/60" />;

const comparisonRows: ComparisonRow[] = [
  {
    label: 'Shared pipeline',
    values: ['One CLI for all adapters', 'Each tool builds its own'],
    icons: [checkIcon, crossIcon],
  },
  {
    label: 'Compliance testing',
    values: [
      'Official JSON Schema Test Suite',
      'Often absent or incomplete',
    ],
    icons: [checkIcon, crossIcon],
  },
  {
    label: 'Supported features',
    values: [
      'Documented per adapter with exact coverage',
      'Undocumented or incomplete',
    ],
    icons: [checkIcon, crossIcon],
  },
  {
    label: 'Unsupported features',
    values: [
      'Explicitly listed with rationale',
      'Silent failures',
    ],
    icons: [checkIcon, crossIcon],
  },
  {
    label: 'Validators',
    values: [
      'Zod, ArkType, Effect, Valibot, Pydantic, more',
      'One converter per library',
    ],
    icons: [checkIcon, crossIcon],
  },
  {
    label: 'Languages',
    values: [
      'TypeScript, Python, Go, Rust',
      'One tool per language',
    ],
    icons: [checkIcon, crossIcon],
  },
];

export function XschemaApproach() {
  return (
    <section className={cn(cardVariants(), 'col-span-full')}>
      <h3 className={cn(headingVariants({ variant: 'h3' }), 'text-brand')}>
        The xschema approach.
      </h3>
      <p className="mt-2 text-fd-muted-foreground">
        One pipeline. Any adapter. Verified.
      </p>
      <div className="mt-6 space-y-4 text-sm leading-relaxed">
        <p>
          xschema runs one pipeline for every adapter: normalize the schema,
          resolve all references, validate it, then hand it off to the adapter
          for code generation. The adapter only does one thing — convert a clean
          schema to native code.
        </p>
        <p>
          Every adapter is tested against the official JSON Schema Test Suite.
        </p>
      </div>
      <ComparisonGrid
        className="mt-8"
        columnHeaders={['xschema', 'One-off converters']}
        rows={comparisonRows}
      />
    </section>
  );
}
