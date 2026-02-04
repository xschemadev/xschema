import { createFileRoute } from "@tanstack/react-router";
import { getLLMText } from "@/lib/get-llm-text";
import { source } from "@/lib/source";

export const Route = createFileRoute("/llms-full.txt")({
  server: {
    handlers: {
      GET: async () => {
        const pages = source.getPages();
        const texts = await Promise.all(pages.map((page) => getLLMText(page)));
        const content = texts.join("\n\n---\n\n");

        return new Response(content, {
          headers: {
            "Content-Type": "text/plain; charset=utf-8",
          },
        });
      },
    },
  },
});
