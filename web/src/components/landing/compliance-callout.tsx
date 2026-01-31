import { cn } from '@/lib/cn';
import { headingVariants } from '@/components/landing/variants';

export function ComplianceCallout() {
  return (
    <section
      className={cn(
        'col-span-full rounded-2xl border border-brand/30 bg-brand/5 p-6 shadow-lg',
        'relative overflow-hidden',
      )}
    >
      <div className="absolute inset-y-0 left-0 w-1 bg-brand" />
      <h3
        className={cn(
          headingVariants({ variant: 'h3' }),
          'text-brand pl-4',
        )}
      >
        Built for transparency.
      </h3>
      <p className="mt-4 pl-4 text-sm leading-relaxed">
        <span className="font-semibold">Compliance report</span>: Tested
        against the official JSON Schema Test Suite. Every adapter reports its
        exact coverage. Supported features, unsupported features, and the
        technical reason behind each gap &mdash; all documented, all tested. Not
        a footnote, a whole documentation page for each adapter.
      </p>
    </section>
  );
}
