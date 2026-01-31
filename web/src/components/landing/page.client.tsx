"use client";

import {
  type ComponentProps,
  Fragment,
  type HTMLAttributes,
  type ReactElement,
  type ReactNode,
  type RefObject,
  useEffect,
  useRef,
  useState,
} from 'react';
import { ArrowRight, TerminalIcon } from 'lucide-react';
import { cn } from '@/lib/cn';
import MainImg from './main.png';
import OpenAPIImg from './openapi.png';
import NotebookImg from './notebook.png';
import { cva } from 'class-variance-authority';
import HeroImage from './hero-preview.jpeg';
import GorillaImage from '@/assets/gorilla.svg';
import { useTheme } from 'next-themes';
import { GrainGradient, Dithering, ImageDithering } from '@paper-design/shaders-react';
import { useIsMobile } from '@/hooks/useMobile';

export function Hero() {
  const { resolvedTheme } = useTheme();
  const ref = useRef<HTMLImageElement | null>(null);
  const visible = useIsVisible(ref);
  const [showShaders, setShowShaders] = useState(false);
  const [imageReady, setImageReady] = useState(false);
  const isMobile = useIsMobile();
  const isMediumScreen = useIsMobile(2000);

  useEffect(() => {
    // apply some delay, otherwise on slower devices, it errors with uniform images not being fully loaded.
    setTimeout(() => {
      setShowShaders(true);
    }, 400);
  }, []);

  // <ImageDithering
  //   width={1400}
  //   height={720}
  //   image="https://paper.design/flowers.webp"
  //   colorBack="#4e5936"
  //   colorFront="#232110"
  //   colorHighlight="#252422"
  //   originalColors={false}
  //   inverted={false}
  //   type="8x8"
  //   size={4}
  //   fit="cover"
  //   className="absolute animate-fd-fade-in duration-400"
  // />

  // <GrainGradient
  //   className="absolute inset-0 animate-fd-fade-in duration-800"
  //   colors={
  //     resolvedTheme === 'dark'
  //       ? ['#39BE1C', '#9c2f05', '#7A2A0000']
  //       : ['#fcfc51', '#ffa057', '#7A2A0020']
  //   }
  //   colorBack="#00000000"
  //   softness={1}
  //   intensity={0.9}
  //   noise={0.5}
  //   speed={visible ? 1 : 0}
  //   shape="corners"
  //   minPixelRatio={1}
  //   maxPixelCount={1920 * 1080}
  // />

  return (
    <>
      {showShaders && (
        <Dithering
          colorBack="#00000000"
          colorFront={resolvedTheme === 'dark' ? '#22231e' : '#ffeeb2'}
          shape="simplex"
          type="4x4"
          size={4}
          speed={0.1}
          className="absolute animate-fd-fade-in duration-400 size-full"
          minPixelRatio={1}
        />
      )}
      {showShaders && (
        <ImageDithering
          image={GorillaImage}
          width={127 * (isMobile ? 2.8 : isMediumScreen ? 3.5 : 4)}
          height={162 * (isMobile ? 2.8 : isMediumScreen ? 3.5 : 4)}
          colorBack="#00000000"
          colorFront={resolvedTheme === 'dark' ? '#6b4b3e' : '#fa8023'}
          originalColors={resolvedTheme === 'light'}
          type="8x8"
          scale={1}
          size={4}
          speed={0}
          frame={5000 * 120}
          className="absolute animate-fd-fade-in duration-400 bottom-0 right-0 sm:bottom-4 sm:right-4 lg:right-20"
          minPixelRatio={1}
        />
      )}
    </>
  );
}

export function CliAnimation() {
  const installCmd = '➜ bun xschema generate';
  const tickTime = 100;
  const timeCommandEnter = installCmd.length;
  const timeCommandRun = timeCommandEnter + 9;
  const timeCommandEnd = timeCommandRun + 9;
  const timeWindowOpen = timeCommandEnd + 1;
  const timeEnd = timeWindowOpen + 1;

  const [tick, setTick] = useState(timeEnd);

  useEffect(() => {
    const timer = setInterval(() => {
      setTick((prev) => (prev >= timeEnd ? prev : prev + 1));
    }, tickTime);

    return () => {
      clearInterval(timer);
    };
  }, [timeEnd]);

  const lines: ReactElement[] = [];

  lines.push(
    <span key="command_type">
      {installCmd.substring(0, tick)}
      {tick < timeCommandEnter && <div className="inline-block h-3 w-1 animate-pulse bg-white" />}
    </span>,
  );

  if (tick >= timeCommandEnter) {
    lines.push(<span key="space"> </span>);
  }

  if (tick > timeCommandRun)
    lines.push(
      <Fragment key="command_response">
        {tick > timeCommandRun + 1 && (
          <>
            <span className="font-bold">{"[1/5] Scanning for xschema config files"}</span>
            <span>{"  → Found 2 config files, 3 schemas (typescript)"}</span>
          </>
        )}
        {tick > timeCommandRun + 2 && (
          <>
            <span className="font-bold">{"[2/5] Fetching schemas"}</span>
            <span>{"✓ Fetching schemas..."}</span>
          </>
        )}
        {tick > timeCommandRun + 3 && (
          <>
            <span>{"  → notes:Note from @xschemadev/zod"}</span>
            <span>{"  → notes:Tag from @xschemadev/arktype"}</span>
            <span>{"  → users:Profile from @xschemadev/effect"}</span>
            <span>{"✓ Fetched 3 schemas"}</span>
          </>
        )}
        {tick > timeCommandRun + 4 && (
          <>
            <span className="font-bold">{"[3/5] Processing schemas"}</span>
            <span>{"✓ Processing schemas..."}</span>
          </>
        )}
        {tick > timeCommandRun + 5 && (
          <>
            <span className="font-bold">{"✓ Processed 3 schemas"}</span>
          </>
        )}
        {tick > timeCommandRun + 6 && (
          <>
            <span className="font-bold">{"[4/5] Generating validators"}</span>
            <span>{"✓ Running adapters..."}</span>
          </>
        )}
        {tick > timeCommandRun + 7 && (
          <>
            <span className="font-bold">{"[5/5] Writing output files"}</span>
            <span>{""}</span>
            <span>{"✓ Generation complete (1.3s)"}</span>
            <span>{""}</span>
          </>
        )}
        {tick > timeCommandRun + 8 && (
          <>
            <span className="font-bold">{"  Schemas generated:"}</span>
            <span>{"    notes"}</span>
            <span>{"      • Note"}</span>
            <span>{"      • Tag"}</span>
            <span>{"    users"}</span>
            <span>{"      • Profile"}</span>
          </>
        )}
      </Fragment>,
    );

  return (
    <div
      className="relative mt-4 w-full mx-auto max-w-[800px]"
      onMouseEnter={() => {
        if (tick >= timeEnd) {
          setTick(0);
        }
      }}
    >
      {tick > timeWindowOpen && (
        <NotificationWindow className="hidden sm:block absolute bottom-5 right-4 z-10 animate-in fade-in slide-in-from-top-10" />
      )}
      <pre className="overflow-hidden rounded-xl border text-sm shadow-lg bg-fd-card">
        <div className="flex flex-row items-center gap-2 border-b px-4 py-2">
          <TerminalIcon className="size-4" /> <span className="font-bold">Terminal</span>
          <div className="grow" />
          <div className="size-2 rounded-full bg-red-400" />
        </div>
        <div className="min-h-128 overflow-x-auto p-4 w-full">
          <code className="grid w-fit">{lines}</code>
        </div>
      </pre>
    </div>
  );
}

function NotificationWindow(props: HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      {...props}
      className={cn(
        'overflow-hidden rounded-md border bg-fd-background shadow-xl',
        props.className,
      )}
    >
      <div className="relative flex h-6 flex-row items-center border-b bg-fd-muted px-4 text-xs text-fd-muted-foreground">
        <p className="absolute inset-x-0 text-center">xschema</p>
      </div>
      <div className="p-4 text-sm">Generation completed!</div>
    </div>
  );
}

const previewButtonVariants = cva('w-20 h-8 text-sm font-medium transition-colors rounded-full', {
  variants: {
    active: {
      true: 'text-fd-primary-foreground',
      false: 'text-fd-muted-foreground',
    },
  },
});

export function PreviewImages(props: ComponentProps<'div'>) {
  const [active, setActive] = useState(0);
  const previews = [
    {
      image: MainImg,
      name: 'Docs',
    },
    {
      image: NotebookImg,
      name: 'Notebook',
    },
    {
      image: OpenAPIImg,
      name: 'OpenAPI',
    },
  ];

  return (
    <div {...props} className={cn('relative grid', props.className)}>
      <div className="absolute flex flex-row left-1/2 -translate-1/2 bottom-0 z-2 p-0.5 rounded-full bg-fd-card border shadow-xl">
        <div
          role="none"
          className="absolute bg-fd-primary rounded-full w-20 h-8 transition-transform z-[-1]"
          style={{
            transform: `translateX(calc(var(--spacing) * 20 * ${active}))`,
          }}
        />
        {previews.map((item, i) => (
          <button
            key={i}
            className={cn(previewButtonVariants({ active: active === i }))}
            onClick={() => setActive(i)}
          >
            {item.name}
          </button>
        ))}
      </div>
      {previews.map((item, i) => (
        <img
          key={i}
          src={item.image}
          alt="preview"
          className={cn(
            'col-start-1 row-start-1 select-none',
            active === i ? 'animate-in fade-in slide-in-from-bottom-12 duration-800' : 'invisible',
          )}
        />
      ))}
    </div>
  );
}

const WritingTabs = [
  {
    name: 'Writer',
    value: 'writer',
  },
  {
    name: 'Developer',
    value: 'developer',
  },
  {
    name: 'Automation',
    value: 'automation',
  },
] as const;

export function Writing({
  tabs: tabContents,
}: {
  tabs: Record<(typeof WritingTabs)[number]['value'], ReactNode>;
}) {
  const [tab, setTab] = useState<(typeof WritingTabs)[number]['value']>('writer');

  return (
    <div className="col-span-full my-20">
      <h2 className="text-4xl text-brand mb-8 text-center font-medium tracking-tight">
        Anybody can write.
      </h2>
      <p className="text-center mb-8 mx-auto w-full max-w-[800px]">
        Native support for Markdown & MDX, offering intuitive, convenient and extensive syntax for
        non-dev writers, developers, and AI agents.
      </p>
      <div className="flex justify-center items-center gap-4 text-fd-muted-foreground mb-6">
        {WritingTabs.map((item) => (
          <Fragment key={item.value}>
            <ArrowRight className="size-4 first:hidden" />
            <button
              className={cn(
                'text-lg font-medium transition-colors',
                item.value === tab && 'text-brand',
              )}
              onClick={() => setTab(item.value)}
            >
              {item.name}
            </button>
          </Fragment>
        ))}
      </div>
      {Object.entries(tabContents).map(([key, value]) => (
        <div
          key={key}
          aria-hidden={key !== tab}
          className={cn('animate-fd-fade-in', key !== tab && 'hidden')}
        >
          {value}
        </div>
      ))}
    </div>
  );
}

export function AgnosticBackground() {
  const { resolvedTheme } = useTheme();
  const ref = useRef<HTMLDivElement>(null);
  const visible = useIsVisible(ref);

  return (
    <div
      ref={ref}
      className="absolute inset-0 -z-1 mask-[linear-gradient(to_top,white_30%,transparent_calc(100%-120px))]"
    >
      <Dithering
        colorBack="#00000000"
        colorFront={resolvedTheme === 'dark' ? '#fc7744' : '#c6bb58'}
        shape="warp"
        type="4x4"
        speed={visible ? 0.4 : 0}
        className="size-full"
        minPixelRatio={1}
      />
    </div>
  );
}

export function ContentAdoptionBackground(props: ComponentProps<typeof GrainGradient>) {
  const { resolvedTheme } = useTheme();

  return (
    <GrainGradient
      colors={
        resolvedTheme === 'dark'
          ? ['#39BE1C', '#9c2f05', '#7A2A0000']
          : ['#DF3F00', '#fcfc51', '#ffa057', '#7A2A0020']
      }
      speed={0}
      colorBack="#1D1004"
      shape="sphere"
      {...props}
    />
  );
}

let observer: IntersectionObserver;
const observerTargets = new WeakMap<Element, (entry: IntersectionObserverEntry) => void>();

function useIsVisible(ref: RefObject<HTMLElement | null>) {
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    observer ??= new IntersectionObserver((entries) => {
      for (const entry of entries) {
        observerTargets.get(entry.target)?.(entry);
      }
    });

    const element = ref.current;
    if (!element) return;
    observerTargets.set(element, (entry) => {
      setVisible(entry.isIntersecting);
    });
    observer.observe(element);

    return () => {
      observer.unobserve(element);
      observerTargets.delete(element);
    };
  }, [ref]);

  return visible;
}
