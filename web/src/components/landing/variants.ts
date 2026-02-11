import { cva } from "class-variance-authority";

export const headingVariants = cva("font-medium tracking-tight", {
  variants: {
    variant: {
      h2: "text-3xl lg:text-4xl",
      h3: "text-xl lg:text-2xl",
    },
  },
});

export const buttonVariants = cva(
  "inline-flex justify-center px-5 py-3 rounded-full font-medium tracking-tight transition-colors",
  {
    variants: {
      variant: {
        primary: "bg-brand text-brand-foreground hover:bg-brand-200",
        secondary:
          "border bg-fd-secondary text-fd-secondary-foreground hover:bg-fd-accent",
      },
    },
    defaultVariants: {
      variant: "primary",
    },
  },
);

export const cardVariants = cva(
  "rounded-2xl text-sm p-6 bg-origin-border shadow-lg",
  {
    variants: {
      variant: {
        secondary: "bg-brand-secondary text-brand-secondary-foreground",
        default: "border bg-fd-card",
      },
    },
    defaultVariants: {
      variant: "default",
    },
  },
);
