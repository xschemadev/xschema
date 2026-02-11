import { createMiddleware, createStart } from "@tanstack/react-start";
import { rewritePath } from "fumadocs-core/negotiation";

// Rewrites /docs/foo/bar.mdx -> /llms.mdx/docs/foo/bar
const { rewrite: rewriteLLM } = rewritePath(
  "/docs/{*path}.mdx",
  "/llms.mdx/docs/{*path}",
);

const llmRewriteMiddleware = createMiddleware().server(
  async ({ request, next }) => {
    const url = new URL(request.url);
    const result = rewriteLLM(url.pathname);

    if (result) {
      // redirect to the llms endpoint
      return Response.redirect(new URL(result, request.url), 307);
    }

    return next();
  },
);

export const startInstance = createStart(() => ({
  requestMiddleware: [llmRewriteMiddleware],
}));
