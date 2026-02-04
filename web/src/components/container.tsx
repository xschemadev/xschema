import { cn } from "@/lib/cn";

export default function Container({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) {
  return (
    <section
      className={cn(
        "bg-fd-background container mx-auto min-h-[calc(100vh-4rem)] border-x",
        className,
      )}
    >
      {children}
    </section>
  );
}
