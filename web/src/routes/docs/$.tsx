import { createFileRoute, notFound } from "@tanstack/react-router";
import { DocsLayout } from "fumadocs-ui/layouts/docs";
import { createServerFn } from "@tanstack/react-start";
import { source } from "@/lib/source";
import browserCollections from "fumadocs-mdx:collections/browser";
import {
  DocsBody,
  DocsDescription,
  DocsPage,
  DocsTitle,
} from "fumadocs-ui/layouts/docs/page";
import defaultMdxComponents from "fumadocs-ui/mdx";
import { Accordion, Accordions } from "fumadocs-ui/components/accordion";
import { Tab, Tabs } from "fumadocs-ui/components/tabs";
import * as Twoslash from "fumadocs-twoslash/ui";
import { Mermaid } from "@/components/mdx/mermaid";
import { baseOptions } from "@/lib/layout.shared";
import { useFumadocsLoader } from "fumadocs-core/source/client";
import { LLMCopyButton, ViewOptions } from "@/components/page-actions";

export const Route = createFileRoute("/docs/$")({
  component: Page,
  loader: async ({ params }) => {
    const slugs = params._splat?.split("/") ?? [];
    const data = await serverLoader({ data: slugs });
    await clientLoader.preload(data.path);
    return data;
  },
});

const serverLoader = createServerFn({
  method: "GET",
})
  .inputValidator((slugs: string[]) => slugs)
  .handler(async ({ data: slugs }) => {
    const page = source.getPage(slugs);
    if (!page) throw notFound();

    return {
      path: page.path,
      pageUrl: page.url,
      pagePath: page.path,
      pageTree: await source.serializePageTree(source.getPageTree()),
    };
  });

interface PageActionsProps {
  pageUrl: string;
  pagePath: string;
}

const clientLoader = browserCollections.docs.createClientLoader<PageActionsProps>({
  component({ toc, frontmatter, default: MDX }, { pageUrl, pagePath }) {
    const markdownUrl = `${pageUrl}.mdx`;
    // pagePath already includes .mdx extension, so don't add it again
    const githubUrl = `https://github.com/xschemadev/xschema/blob/main/web/content/docs/${pagePath}`;

    return (
      <DocsPage
        toc={toc}
        tableOfContent={{ style: "clerk" }}
        tableOfContentPopover={{ style: "clerk" }}
      >
        <DocsTitle>{frontmatter.title}</DocsTitle>
        <DocsDescription>{frontmatter.description}</DocsDescription>
        <div className="flex flex-row gap-2 items-center -mt-6 border-b pb-6">
          <LLMCopyButton markdownUrl={markdownUrl} />
          <ViewOptions markdownUrl={markdownUrl} githubUrl={githubUrl} />
        </div>
        <DocsBody>
          <MDX
            components={{
              ...defaultMdxComponents,
              ...Twoslash,
              Accordion,
              Accordions,
              Tab,
              Tabs,
              Mermaid,
            }}
          />
        </DocsBody>
      </DocsPage>
    );
  },
});

function Page() {
  const data = Route.useLoaderData();
  const { pageTree } = useFumadocsLoader(data);
  const Content = clientLoader.getComponent(data.path);

  return (
    <DocsLayout {...baseOptions()} tree={pageTree}>
      <Content pageUrl={data.pageUrl} pagePath={data.pagePath} />
    </DocsLayout>
  );
}
