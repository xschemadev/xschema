import { headingVariants, cardVariants } from '@/components/landing/variants';
import { cn } from '@/lib/cn';
import { Bug, FlaskConicalOff, Replace, Workflow } from 'lucide-react';

const problems = [
  {
    label: 'No shared pipeline',
    description:
      'Each tool re-implements parsing, normalization, and reference resolution from scratch.',
    icon: <Workflow className="size-5" />,
  },
  {
    label: 'No correctness guarantees',
    description:
      'None of them test against the official JSON Schema Test Suite.',
    icon: <FlaskConicalOff className="size-5" />,
  },
  {
    label: 'No transparency',
    description:
      'Unsupported features fail silently or produce wrong output.',
    icon: <Bug className="size-5" />,
  },
  {
    label: 'No portability',
    description:
      'Switching validators means switching converters, configs, and workflows.',
    icon: <Replace className="size-5" />,
  },
];

export function JungleProblem({ className }: { className?: string }) {
  return (
    <section className={cn(className)}>
      <h2 className={cn(headingVariants({ variant: 'h2' }), 'text-brand text-center font-mono font-bold uppercase')}>
        The jungle problem.
      </h2>
      <p className="mt-2 text-fd-muted-foreground text-center">
        Today, every library has its own converter.
      </p>
      <div className="mt-4 space-y-6 leading-relaxed max-w-5xl mx-auto">
        <p className="text-balance font-light max-w-md mx-auto border-dotted border-l-2 border-brand-200 pl-4 italic">
          JSON Schema to Zod? There&apos;s a package for that. Pydantic? Another one. ArkType? You&apos;re on your own.
        </p>
        <p className='text-center text-lg font-semibold'>You end up with</p>
        <ul className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {problems.map((item) => (
            <li key={item.label} className={cn(cardVariants({ variant: 'default' }), "px-0 relative group/feature bg-transparent border-outset")}>
              <div className="pointer-events-none absolute inset-0 z-0 size-full bg-linear-to-t from-fd-card to-transparent rounded-2xl opacity-0 transition duration-200 group-hover/feature:opacity-100" />
              <div className="relative px-6 mb-3 z-10 text-muted-foreground">{item.icon}</div>
              <div className="relative z-10 mb-2 px-6 text-lg font-bold">
                <div className="absolute inset-y-0 left-0 h-6 w-1 origin-center rounded-br-full rounded-tr-full bg-fd-border transition-all duration-200 group-hover/feature:h-8 group-hover/feature:bg-brand" />
                <span className="text-brand-200 inline-block transition duration-200 group-hover/feature:translate-x-2 group-hover/feature:text-brand">
                  {item.label}
                </span>
              </div>
              <p className='relative text-base z-10 text-balance px-6'>{item.description}</p>
            </li>
          ))}
        </ul>
      </div>
    </section>
  );
}
