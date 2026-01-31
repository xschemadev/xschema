import { baseOptions, linkItems } from "@/lib/layout.shared";
import { Link } from "@tanstack/react-router";
import { HomeLayout } from "fumadocs-ui/layouts/home";
import { Book, Puzzle } from "lucide-react";

export function Error() {
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
                    url: "/dcocs",
                },
                ...linkItems,
            ]}
        >
            <div className="flex flex-col items-center gap-4 justify-center flex-1">
                <h1 className="text-6xl font-bold text-fd-muted-foreground">500</h1>
                <h2 className="text-2xl font-semibold">Internal Server Error</h2>
                <p className="text-fd-muted-foreground max-w-md">
                    An unexpected error occurred on the server.
                </p>
                <Link
                    to="/"
                    className="mt-4 px-4 py-2 rounded-lg bg-fd-primary text-fd-primary-foreground font-medium text-sm hover:opacity-90 transition-opacity"
                >
                    Back to Home
                </Link>
            </div>
        </HomeLayout>
    );
}
