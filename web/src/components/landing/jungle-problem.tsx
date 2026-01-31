import { headingVariants, cardVariants } from '@/components/landing/variants';
import { cn } from '@/lib/cn';

const problems = [
  {
    label: 'No shared pipeline',
    description:
      'each tool re-implements parsing, normalization, and reference resolution from scratch.',
  },
  {
    label: 'No correctness guarantees',
    description:
      'none of them test against the official JSON Schema Test Suite.',
  },
  {
    label: 'No transparency',
    description:
      'unsupported features fail silently or produce wrong output.',
  },
  {
    label: 'No portability',
    description:
      'switching validators means switching converters, configs, and workflows.',
  },
];

export function JungleProblem() {
  return (
    <section className={cn(cardVariants(), 'col-span-full')}>
      <h3 className={cn(headingVariants({ variant: 'h3' }), 'text-brand')}>
        The jungle problem.
      </h3>
      <p className="mt-2 text-fd-muted-foreground">
        Today, every library has its own converter.
      </p>
      <div className="mt-6 space-y-4 text-sm leading-relaxed">
        <p>
          Need to convert JSON Schema to a Zod validator? There&apos;s a package
          for that. Pydantic? Another one. ArkType? You&apos;re on your own.
        </p>
        <p>
          Each converter follows a different flow, supports a different subset of
          the spec, and there&apos;s no standard way to verify correctness. You
          end up with:
        </p>
        <ul className="space-y-3">
          {problems.map((item) => (
            <li key={item.label} className="flex items-start gap-2">
              <span className="mt-1.5 size-1.5 shrink-0 rounded-full bg-brand" />
              <span>
                <strong>{item.label}</strong> — {item.description}
              </span>
            </li>
          ))}
        </ul>
      </div>
    </section>
  );
}
