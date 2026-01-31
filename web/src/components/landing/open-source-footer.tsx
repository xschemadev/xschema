import { Link } from '@tanstack/react-router';
import { cn } from '@/lib/cn';
import { headingVariants, buttonVariants } from '@/components/landing/variants';
import { XSCHEMA_GITHUB_URL } from '@/lib/constants';

export function OpenSourceFooter() {
  return (
    <section className="col-span-full mt-16 flex flex-col items-center text-center">
      <h2
        className={cn(
          headingVariants({ variant: 'h2' }),
          'text-brand font-mono font-bold uppercase',
        )}
      >
        Open source. Open process.
      </h2>
      <p className="mt-4 max-w-xl text-sm leading-relaxed text-fd-muted-foreground">
        xschema is open source. Contributions, adapter proposals, and bug
        reports are greatly appreciated.
      </p>
      <div className="mt-8 flex flex-row items-center gap-4 flex-wrap justify-center">
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
          GitHub
        </a>
      </div>
    </section>
  );
}
