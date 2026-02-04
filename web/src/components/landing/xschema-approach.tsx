import { Check, X } from 'lucide-react';
import { headingVariants, cardVariants } from '@/components/landing/variants';
import { cn } from '@/lib/cn';
import {
  ComparisonGrid,
  type ComparisonRow,
} from '@/components/ui/comparison-grid';
import { Arrow } from '../ui/arrow';
import { PlusCard } from '../ui/plus-card';

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
    label: 'Feature support',
    values: [
      'Documented per adapter with exact coverage',
      'Missing or incomplete',
    ],
    icons: [checkIcon, crossIcon],
  },
  {
    label: 'Languages',
    values: [
      'Multi-language: TypeScript, Python, Go, Rust',
      'One tool per language',
    ],
    icons: [checkIcon, crossIcon],
  },
];

export function XschemaApproach({ className }: { className?: string }) {
  return (
    <section className={cn(className)}>
      <p className="text-brand font-mono font-bold mb-10 relative max-w-md mx-auto uppercase">
        The solution
        <Arrow
          padding={40}
          className="absolute z-2"
          fromX={130}
          toX={200}
          curveX={300}
          fromY={12}
          toY={100}
          curveY={20}
          strokeWidth={3}
          tipSize={10}
          color="#c6bb58"
        />
      </p>
      <PlusCard className={cn('py-10 max-w-4xl mx-auto')}>
        <h2 className={cn(headingVariants({ variant: 'h2' }), 'text-center font-mono font-bold')}>
          xschema APPROACH.
        </h2>
        <p className="mt-2 text-fd-muted-foreground text-center">
          One pipeline to rule them all.
        </p>
        <div className="mt-6 space-y-4 text-lg leading-relaxed">
          {/* <p>xschema runs a UNIFIED pipeline for every adapter:</p>
          <ul className="list-disc list-inside">
            <li>The CLI does the heavy lifting.</li>
            <li>Adapters receive clean, self-contained schemas and only handle code generation.</li>
          </ul> */}
          <div className='flex flex-col max-w-md mx-auto'>
            <div className='font-mono w-fit pr-4'>
              <p className="text-base text-brand-secondary">CLI</p>
              <ul className="text-sm border border-brand-secondary border-dashed p-2 [&>li]:ml-3 [&>li]:list-disc">
                <li>parses your config files</li>
                <li>fetches your schemas</li>
                <li>normalizes them</li>
                <li>resolves all references</li>
                <li>validates them</li>
              </ul>
              <span className='text-xs tracking-tight text-brand-secondary'>{"-->"}</span>
            </div>
            <div className='font-mono w-fit self-end text-right pl-4'>
              <p className="text-base text-brand">ADAPTER</p>
              <ul className="text-sm border border-brand border-dashed p-2 [&>li]:ml-3 [&>li]:list-disc">
                <li>handles code generation</li>
              </ul>
              <span className='text-xs tracking-tight text-brand'>{"<--"}</span>
            </div>
            <div className='font-mono w-fit pr-4'>
              <p className="text-base text-brand-secondary">CLI</p>
              <ul className="text-sm border border-brand-secondary border-dashed p-2 [&>li]:ml-3 [&>li]:list-disc">
                <li>injects the schemas in your codebase</li>
              </ul>
            </div>
          </div>
          {/*
          parses your config files
          fetches your schemas
          normalizes them
          resolves all references
          validates them
          hands them off to the adapter for code generation */}
        </div>
      </PlusCard>
      <ComparisonGrid
        className="my-10 max-w-6xl mx-auto"
        columnHeaders={['xschema', 'One-off converters']}
        rows={comparisonRows}
      />
      {/* <ComplianceCallout className="max-w-4xl mx-auto" /> */}
    </section>
  );
}

function ComplianceCallout({ className }: { className?: string }) {
  return (
    <section
      className={cn(
        cardVariants({ variant: 'secondary' }),
        className,
      )}
    >
      <h3
        className={cn(
          headingVariants({ variant: 'h3' }),
        )}
      >
        Built for compliance.
      </h3>
      <ul className="mt-4 text-lg [&>li]:border-b [&>li]:border-brand/30 [&>li]:py-2 [&>li]:last:border-b-0">
        <li>
          Tested against the official JSON Schema Test Suite.
        </li>
        <li>
          Every adapter reports its
          exact coverage. Supported features, unsupported features, and the
          technical reason behind each gap &mdash; all documented, all tested.
        </li>
        <li>
          Not a footnote, a whole documentation page for each adapter.
        </li>
      </ul>
    </section>
  );
}
