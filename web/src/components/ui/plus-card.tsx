import { cn } from "@/lib/cn";

const PlusCard = ({
  children,
  className,
}: {
  children?: React.ReactNode;
  className?: string;
}) => {
  return (
    <div
      className={cn(
        "relative min-h-32 rounded-2xl border border-dashed p-6",
        "flex flex-col justify-between",
        className,
      )}
    >
      {children}
      <PlusIcon className="absolute -top-3 -left-3" />
      <PlusIcon className="absolute -top-3 -right-3" />
      <PlusIcon className="absolute -bottom-3 -left-3" />
      <PlusIcon className="absolute -right-3 -bottom-3" />
    </div>
  );
};

const PlusIcon = ({ className }: { className?: string }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    width={24}
    height={24}
    strokeWidth="2"
    stroke="currentColor"
    className={`text-brand size-6 ${className}`}
  >
    <path strokeLinecap="round" strokeLinejoin="round" d="M12 6v12m6-6H6" />
  </svg>
);

export { PlusCard };
