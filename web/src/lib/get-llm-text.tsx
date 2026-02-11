import type { InferPageType } from "fumadocs-core/source";
import type { source } from "./source";
import { getSection } from "./source/navigation";
import {
  XSCHEMA_GITHUB_ORGANIZATION,
  XSCHEMA_GITHUB_REPO,
  SECTION_TAB_TITLE_MAP,
} from "./constants";

export async function getLLMText(page: InferPageType<typeof source>) {
  const section = getSection(page.slugs[0]);

  const category = SECTION_TAB_TITLE_MAP[section] ?? section;

  const processed = await page.data.getText("processed");

  return `# ${category}: ${page.data.title}
URL: ${page.url}
Source: https://raw.githubusercontent.com/${XSCHEMA_GITHUB_ORGANIZATION}/${XSCHEMA_GITHUB_REPO}/refs/heads/master/web/content/docs/${page.path}

${page.data.description ?? ""}
        
${processed}`;
}
