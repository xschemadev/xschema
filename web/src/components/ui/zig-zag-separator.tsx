import { cn } from "@/lib/cn";

const ZigzagSeparator = ({ className }: { className?: string }) => (
    <div
        className={cn("bg-primary h-3 w-full", className)}
        style={{
            maskImage: `url("data:image/svg+xml,%3Csvg width='24' height='12' viewBox='0 0 24 12' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M0 12 L12 0 L24 12' fill='none' stroke='black' stroke-width='2'/%3E%3C/svg%3E")`,
            maskRepeat: "repeat-x",
            WebkitMaskImage: `url("data:image/svg+xml,%3Csvg width='24' height='12' viewBox='0 0 24 12' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M0 12 L12 0 L24 12' fill='none' stroke='black' stroke-width='2'/%3E%3C/svg%3E")`,
            WebkitMaskRepeat: "repeat-x",
        }}
    />
);

export { ZigzagSeparator };