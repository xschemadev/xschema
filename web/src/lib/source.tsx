import { loader, LoaderPlugin } from "fumadocs-core/source";
import { docs } from "fumadocs-mdx:collections/server";
import { lucideIconsPlugin } from "fumadocs-core/source/lucide-icons";

export const source = loader({
  source: docs.toFumadocsSource(),
  baseUrl: "/docs",
  plugins: [pageTreeCodeTitles(), lucideIconsPlugin()],
});

function pageTreeCodeTitles(): LoaderPlugin {
  return {
    transformPageTree: {
      file(node) {
        if (
          typeof node.name === "string" &&
          (node.name.endsWith("()") || node.name.match(/^<\w+ \/>$/))
        ) {
          return {
            ...node,
            name: (
              <code key="0" className="text-[0.8125rem]">
                {node.name}
              </code>
            ),
          };
        }
        return node;
      },
    },
  };
}
