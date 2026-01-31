export const XSCHEMA_NAME = "xschema";
export const XSCHEMA_GITHUB_ORGANIZATION = "xschemadev";
export const XSCHEMA_GITHUB_REPO = "xschema";
export const XSCHEMA_GITHUB_URL = `https://github.com/${XSCHEMA_GITHUB_ORGANIZATION}/${XSCHEMA_GITHUB_REPO}`;
export const MOTTO = "Bring your JSON Schemas to life";

export const sectionTabsList = [
  "framework",
  "runtime",
  "cli",
  "adapters",
  "compliance",
] as const;
export type SectionTab = (typeof sectionTabsList)[number];

export const SECTION_TAB_TITLE_MAP: Record<SectionTab, string> = {
  framework: `${XSCHEMA_NAME} (Framework Mode)`,
  runtime: `${XSCHEMA_NAME} Runtime (Programmatic API)`,
  cli: `${XSCHEMA_NAME} CLI (Command-line interface)`,
  adapters: `${XSCHEMA_NAME} Adapters (Language-specific code generation)`,
  compliance: `${XSCHEMA_NAME} Compliance (Compliance testing)`,
};
