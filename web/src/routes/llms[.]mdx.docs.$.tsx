import { createFileRoute, notFound } from "@tanstack/react-router";
import { getLLMText } from "@/lib/get-llm-text";
import { source } from "@/lib/source";

export const Route = createFileRoute("/llms.mdx/docs/$")({
  server: {
    handlers: {
      GET: async ({ params }) => {
        const slugs = params._splat?.split("/") ?? [];
        const page = source.getPage(slugs);

        if (!page) {
          throw notFound();
        }

        const content = await getLLMText(page);

        return new Response(content, {
          headers: {
            "Content-Type": "text/markdown; charset=utf-8",
          },
        });
      },
    },
  },
});
