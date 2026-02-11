import { cn } from "@/lib/cn";
import { getSection } from "@/lib/source/navigation";
import { createFileRoute, Outlet, useLocation } from "@tanstack/react-router";
import { RootProvider } from "fumadocs-ui/provider/tanstack";
import { ReactNode } from "react";

export const Route = createFileRoute("/_layout")({
  component: RouteComponent,
});

export function Body({
  children,
}: {
  children: ReactNode;
}): React.ReactElement {
  const mode = useMode();

  return (
    <div className={cn(mode, "relative flex min-h-screen flex-col")}>
      {children}
    </div>
  );
}

export function useMode(): string | undefined {
  const { pathname } = useLocation();
  const slug = pathname.split("/").slice(1);
  if (Array.isArray(slug)) return getSection(slug[1]);
  return undefined;
}

function RouteComponent() {
  return (
    <Body>
      <RootProvider>
        <Outlet />
      </RootProvider>
    </Body>
  );
}
