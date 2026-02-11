import type { ReactNode } from "react";
import { cn } from "@/lib/cn";

export interface ComparisonRow {
  label: string;
  values: [string, string];
  icons?: [ReactNode, ReactNode];
}

export interface ComparisonGridProps {
  columnHeaders: [string, string];
  rows: ComparisonRow[];
  className?: string;
}

export function ComparisonGrid({
  columnHeaders,
  rows,
  className,
}: ComparisonGridProps) {
  return (
    <div className={cn("w-full overflow-x-auto", className)}>
      <table className="w-full border-collapse">
        <thead>
          <tr className="border-b-4 border-fd-border">
            <th className="py-3 px-4 text-left font-medium text-fd-muted-foreground w-[30%]" />
            <th className="py-3 px-4 text-left text-xl font-medium text-brand">
              {columnHeaders[0]}
            </th>
            <th className="py-3 px-4 text-left text-xl font-medium text-fd-muted-foreground">
              {columnHeaders[1]}
            </th>
          </tr>
        </thead>
        <tbody>
          {rows.map((row) => (
            <tr
              key={row.label}
              className="border-b border-fd-border/50 last:border-b-0"
            >
              <td className="py-3 px-4 font-medium text-fd-foreground">
                {row.label}
              </td>
              <td className="py-3 px-4">
                <span className="flex items-center gap-2">
                  {row.icons?.[0] && (
                    <span className="shrink-0">{row.icons[0]}</span>
                  )}
                  <span className="text-fd-foreground">{row.values[0]}</span>
                </span>
              </td>
              <td className="py-3 px-4">
                <span className="flex items-center gap-2">
                  {row.icons?.[1] && (
                    <span className="shrink-0">{row.icons[1]}</span>
                  )}
                  <span className="text-fd-muted-foreground">
                    {row.values[1]}
                  </span>
                </span>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
