"use client";

import {
  Fragment,
  type HTMLAttributes,
  type ReactElement,
  useEffect,
  useState,
} from 'react';
import { TerminalIcon } from 'lucide-react';
import { cn } from '@/lib/cn';
import GorillaImage from '@/assets/gorilla.svg';
import { useTheme } from 'next-themes';
import { Dithering, ImageDithering } from '@paper-design/shaders-react';
import { useIsMobile } from '@/hooks/useMobile';

export function Hero() {
  const { resolvedTheme } = useTheme();
  const [showShaders, setShowShaders] = useState(false);
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

export function CliAnimation({ className }: { className?: string }) {
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
            <span className="text-brand">{"  → users:Profile from @xschemadev/effect"}</span>
            <span className="text-brand-secondary">{"  → notes:Note from @xschemadev/zod"}</span>
            <span className="text-brand-secondary">{"  → notes:Tag from @xschemadev/arktype"}</span>
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
            <span className="text-brand">{"    users"}</span>
            <span className="text-brand">{"      • Profile"}</span>
            <span className="text-brand-secondary">{"    notes"}</span>
            <span className="text-brand-secondary">{"      • Note"}</span>
            <span className="text-brand-secondary">{"      • Tag"}</span>
          </>
        )}
      </Fragment>,
    );

  return (
    <div
      className={cn("relative", className)}
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
