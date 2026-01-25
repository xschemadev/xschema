import { createFileRoute, Link } from "@tanstack/react-router";
import { HomeLayout } from "fumadocs-ui/layouts/home";
import { baseOptions, linkItems } from "@/lib/layout.shared";
import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/cn";
import { Book, Puzzle } from "lucide-react";

export const Route = createFileRoute("/")({
  component: Home,
});

function Home() {
  return (
    <HomeLayout
      {...baseOptions()}
      links={[
        {
          type: "menu",
          on: "menu",
          text: "Documentation",
          items: [
            {
              text: "Getting Started",
              url: "/docs",
              icon: <Book />,
            },
            {
              text: "Adapters",
              url: "/docs/adapters",
              icon: <Puzzle />,
            },
          ],
        },
        {
          type: "main",
          on: "nav",
          text: "Documentation",
          url: "/docs",
        },
        ...linkItems,
      ]}
    >
      <div className="flex flex-col flex-1 justify-center px-4 py-8 text-center">
        <h1 className="font-medium text-xl mb-4">
          Fumadocs on Tanstack Start.
        </h1>
        <Link
          to="/docs/$"
          params={{
            _splat: "",
          }}
          className={cn(buttonVariants({ variant: "primary" }), "mx-auto")}
        >
          Open Docs
        </Link>
      </div>
    </HomeLayout>
  );
}
