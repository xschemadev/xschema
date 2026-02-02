import { cn } from '@/lib/cn';
import type { ReactNode } from 'react';

export interface FeatureGridItem {
  title: string;
  description: string;
  icon: ReactNode;
  detail?: string;
}

export interface FeatureGridProps {
  items: FeatureGridItem[];
  columns?: 2 | 3 | 4;
  className?: string;
}

const columnClasses = {
  2: 'md:grid-cols-2',
  3: 'md:grid-cols-2 lg:grid-cols-3',
  4: 'md:grid-cols-2 lg:grid-cols-4',
} as const;

export function FeatureGrid({ items, columns = 2, className }: FeatureGridProps) {
  return (
    <div className={cn('relative z-10 grid grid-cols-1', columnClasses[columns], className)}>
      {items.map((item, index) => (
        <FeatureGridCard key={item.title} item={item} index={index} total={items.length} columns={columns} />
      ))}
    </div>
  );
}

function FeatureGridCard({
  item,
  index,
  total,
  columns,
}: {
  item: FeatureGridItem;
  index: number;
  total: number;
  columns: number;
}) {
  const isFirstInRow = index % columns === 0;
  const isInTopHalf = index < columns;

  return (
    <div
      className={cn(
        'group/feature relative flex flex-col border-fd-border py-8 lg:border-r',
        isFirstInRow && 'lg:border-l',
        isInTopHalf && total > columns && 'lg:border-b',
      )}
    >
      {isInTopHalf ? (
        <div className="pointer-events-none absolute inset-0 size-full bg-gradient-to-t from-fd-card to-transparent opacity-0 transition duration-200 group-hover/feature:opacity-100" />
      ) : (
        <div className="pointer-events-none absolute inset-0 size-full bg-gradient-to-b from-fd-card to-transparent opacity-0 transition duration-200 group-hover/feature:opacity-100" />
      )}
      <div className="relative z-10 mb-4 px-8 text-fd-muted-foreground">
        {item.icon}
      </div>
      <div className="relative z-10 mb-2 px-8 text-lg font-bold">
        <div className="absolute inset-y-0 left-0 h-6 w-1 origin-center rounded-br-full rounded-tr-full bg-fd-border transition-all duration-200 group-hover/feature:h-8 group-hover/feature:bg-brand" />
        <span className="inline-block text-fd-foreground transition duration-200 group-hover/feature:translate-x-2">
          {item.title}
        </span>
      </div>
      <p className="relative z-10 max-w-xs px-8 text-sm text-fd-muted-foreground">
        {item.description}
      </p>
      {item.detail && (
        <div className="relative z-10 mt-auto px-8 pt-4">
          <span className="inline-block rounded border border-fd-border bg-fd-background px-2 py-1 font-mono text-xs text-fd-muted-foreground">
            {item.detail}
          </span>
        </div>
      )}
    </div>
  );
}
