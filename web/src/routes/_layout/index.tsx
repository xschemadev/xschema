import { createFileRoute, Link } from "@tanstack/react-router";
import { HomeLayout } from "fumadocs-ui/layouts/home";
import { baseOptions, linkItems } from "@/lib/layout.shared";
import { Book, Puzzle } from "lucide-react";
import { cn } from '@/lib/cn';
import {
  Hero,
  CliAnimation,
} from '@/components/landing/page.client';
import { buttonVariants, headingVariants } from '@/components/landing/variants';
import { JungleProblem } from '@/components/landing/jungle-problem';
import { XschemaApproach } from '@/components/landing/xschema-approach';
import { ComplianceCallout } from '@/components/landing/compliance-callout';
import { BuiltFor } from '@/components/landing/built-for';
import { OpenSourceFooter } from '@/components/landing/open-source-footer';
import { ComponentCards } from '@/components/landing/component-cards';
import { PipelineAnimation } from '@/components/landing/pipeline-animation';
import { TryIt } from '@/components/landing/try-it';
import { MOTTO, XSCHEMA_GITHUB_URL } from "@/lib/constants";
import Container from "@/components/container";

export const Route = createFileRoute("/_layout/")({
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
      <Page />
    </HomeLayout>
  );
}

function Page() {
  return (
    <main className="text-landing-foreground bg-[repeating-linear-gradient(45deg,color-mix(in_oklab,var(--border)40%,transparent)0,color-mix(in_oklab,var(--border)40%,transparent)1px,transparent_0,transparent_50%)] dark:bg-[repeating-linear-gradient(45deg,color-mix(in_oklab,var(--border)10%,transparent)0,color-mix(in_oklab,var(--border)10%,transparent)1px,transparent_0,transparent_50%)] bg-size-[12px_12px] bg-fixed">
      <Container className="pt-4 pb-6 md:pb-12 ">
        <div className="relative flex min-h-[600px] h-[70vh] max-h-[900px] border rounded-2xl overflow-hidden mx-auto w-full max-w-[1400px] bg-origin-border">
          <Hero />
          <div className="flex flex-col z-2 px-4 size-full md:p-12 max-md:items-center max-md:text-center">
            <p className="mt-12 text-sm font-mono text-brand font-medium rounded-full py-2 px-4 border border-brand/50 w-fit">
              {MOTTO}.
            </p>
            <h1 className="text-4xl my-8 leading-tight font-medium xl:text-5xl xl:mb-12">
              Cross-language type-safety,
              <br />but it's <span className="underline decoration-wavy decoration-brand decoration-3 underline-offset-4">not</span> a
              <span className="italic font-thin"> f@&#ing</span> <span className="text-brand">jungle</span>.
            </h1>
            <div className="flex flex-row items-center justify-center gap-4 flex-wrap w-fit">
              <Link to="/docs/$" className={cn(buttonVariants(), 'max-sm:text-sm')}>
                Getting Started
              </Link>
              <a
                href={XSCHEMA_GITHUB_URL}
                target="_blank"
                rel="noreferrer noopener"
                className={cn(buttonVariants({ variant: 'secondary' }), 'max-sm:text-sm')}
              >
                Open in StackBlitz
              </a>
            </div>
          </div>
        </div>
        <div className="grid grid-cols-1 gap-10 mt-12 px-6 mx-auto w-full md:px-12 lg:grid-cols-2">
          <p className="text-2xl tracking-tight leading-snug font-light col-span-full md:text-3xl xl:text-4xl">
            <span className="underline decoration-wavy decoration-brand decoration-3 underline-offset-4">xschema</span> is an ecosystem of tools to bring <span className="text-brand font-medium">cross-language type-safety validation </span>
            in your codebases. It leverages the power of <span className="text-brand font-medium">JSON Schema</span> to validate your data across languages.
          </p>
          <Separator className="col-span-full" />
          <h2
            className={cn(
              headingVariants({ variant: 'h2' }),
              'col-span-full text-brand text-center font-mono font-bold uppercase',
            )}
          >
            One pipeline to rule them all.
          </h2>
          <JungleProblem />
          <XschemaApproach />
          <ComplianceCallout />
          <BuiltFor />
          <Separator className="col-span-full" />
          <PipelineAnimation>
            <CliAnimation />
          </PipelineAnimation>
          <Separator className="col-span-full" />
          <ComponentCards />
          <Separator className="col-span-full" />
          <TryIt />
          <Separator className="col-span-full" />
          <OpenSourceFooter />
        </div>
      </Container>
    </main>
  );
}

const Separator = ({ className }: { className?: string }) => {
  return (
    <div className={cn("border rounded-xl bg-[repeating-linear-gradient(45deg,color-mix(in_oklab,var(--border)40%,transparent)0,color-mix(in_oklab,var(--border)40%,transparent)1px,transparent_0,transparent_50%)] dark:bg-[repeating-linear-gradient(45deg,color-mix(in_oklab,var(--border)10%,transparent)0,color-mix(in_oklab,var(--border)10%,transparent)1px,transparent_0,transparent_50%)] bg-size-[12px_12px] h-10 w-full", className)} />
  );
};
