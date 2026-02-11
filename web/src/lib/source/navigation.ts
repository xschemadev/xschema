import { SectionTab } from "../constants";

export function getSection(path: string | undefined): SectionTab {
  if (!path) return "framework";
  const [dir] = path.split("/", 1);
  if (!dir) return "framework";
  return (
    {
      runtime: "runtime" as const,
      cli: "cli" as const,
      adapters: "adapters" as const,
      compliance: "compliance" as const,
    }[dir] ?? ("framework" as const)
  );
}
