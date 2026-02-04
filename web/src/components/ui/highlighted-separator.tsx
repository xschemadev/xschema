import { cn } from "@/lib/cn";

export const HighlightedSeparator = ({ className }: { className?: string }) => {
    return (
        <div className={cn("border rounded-xl bg-[repeating-linear-gradient(45deg,color-mix(in_oklab,var(--border)40%,transparent)0,color-mix(in_oklab,var(--border)40%,transparent)1px,transparent_0,transparent_50%)] dark:bg-[repeating-linear-gradient(45deg,color-mix(in_oklab,var(--border)10%,transparent)0,color-mix(in_oklab,var(--border)10%,transparent)1px,transparent_0,transparent_50%)] bg-size-[12px_12px] h-10 w-full", className)} />
    );
};
