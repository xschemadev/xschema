import { Terminal, Blocks, Package, FileJson } from 'lucide-react';
import { FeatureGrid } from '@/components/feature-grid';
import type { FeatureGridItem } from '@/components/feature-grid';

const componentCards: FeatureGridItem[] = [
  {
    title: 'CLI',
    description:
      'Orchestrates the full pipeline. Parses configs, fetches schemas, normalizes drafts, resolves references, and delegates generation to adapters.',
    icon: <Terminal className="size-5" />,
    detail: 'npx xschema generate',
  },
  {
    title: 'Adapters',
    description:
      'Each adapter targets one validation library and converts normalized schemas into native validators. One adapter per library, one protocol for all.',
    icon: <Blocks className="size-5" />,
    detail: 'Zod \u00b7 ArkType \u00b7 Effect Schema \u00b7 Valibot \u00b7 Pydantic \u00b7 and more',
  },
  {
    title: 'Client',
    description:
      'Language-specific packages that expose generated schemas with full type inference and autocomplete.',
    icon: <Package className="size-5" />,
    detail: '@xschemadev/client',
  },
  {
    title: 'Config',
    description:
      'Declarative config files define which schemas to process, where to fetch them, and which adapter to use.',
    icon: <FileJson className="size-5" />,
    detail: '*.xschema.jsonc',
  },
];

export function ComponentCards() {
  return (
    <div className="col-span-full">
      <FeatureGrid items={componentCards} columns={2} />
    </div>
  );
}
