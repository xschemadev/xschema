export const XSCHEMA_NAME = "xschema";
export const XSCHEMA_GITHUB_ORGANIZATION = "xschemadev";
export const XSCHEMA_GITHUB_REPO = "xschema";
export const XSCHEMA_GITHUB_URL = `https://github.com/${XSCHEMA_GITHUB_ORGANIZATION}/${XSCHEMA_GITHUB_REPO}`;

export const sectionTabsList = [
  "framework",
  "runtime",
  "cli",
  "adapters",
  "compliance",
] as const;
export type SectionTab = (typeof sectionTabsList)[number];
