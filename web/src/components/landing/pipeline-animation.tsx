"use client";

import { useState, type ReactNode } from "react";
import { cn } from "@/lib/cn";
import { headingVariants } from "@/components/landing/variants";

interface PipelineStage {
  id: string;
  label: string;
  description: string;
}

const stages: PipelineStage[] = [
  {
    id: "parse",
    label: "Parse",
    description:
      "Discovers config files and extracts schema declarations.",
  },
  {
    id: "retrieve",
    label: "Retrieve",
    description:
      "Fetches schemas from files, URLs, or inline definitions.",
  },
  {
    id: "process",
    label: "Process",
    description:
      "Normalizes, validates, resolves references, bundles.",
  },
  {
    id: "generate",
    label: "Generate",
    description:
      "Delegates code generation to the target adapter.",
  },
  {
    id: "inject",
    label: "Inject",
    description:
      "Writes generated validators and types to your project.",
  },
];

const PIPELINE_ID = "pipeline";

// SVG layout constants for the pipeline
const SVG_WIDTH = 760;
const SVG_HEIGHT = 100;
const NODE_WIDTH = 110;
const NODE_HEIGHT = 40;
const NODE_RX = 8;
const NODE_Y = (SVG_HEIGHT - NODE_HEIGHT) / 2;
const NODE_SPACING = (SVG_WIDTH - stages.length * NODE_WIDTH) / (stages.length - 1);

function getNodeX(index: number): number {
  return index * (NODE_WIDTH + NODE_SPACING);
}

function PipelineSvg({
  activeIndex,
  onStageHover,
  onStageClick,
}: {
  activeIndex: number;
  onStageHover: (index: number) => void;
  onStageClick: (index: number) => void;
}) {
  // Build connector paths between consecutive nodes
  const connectors: { d: string; index: number }[] = [];
  for (let i = 0; i < stages.length - 1; i++) {
    const x1 = getNodeX(i) + NODE_WIDTH;
    const x2 = getNodeX(i + 1);
    const cy = SVG_HEIGHT / 2;
    connectors.push({
      d: `M ${x1} ${cy} L ${x2} ${cy}`,
      index: i,
    });
  }

  return (
    <svg
      viewBox={`0 0 ${SVG_WIDTH} ${SVG_HEIGHT}`}
      className="w-full h-auto"
      preserveAspectRatio="xMidYMid meet"
    >
      <defs>
        <radialGradient id={`${PIPELINE_ID}-light-grad`} fx="0.5" fy="0.5">
          <stop offset="0%" stopColor="var(--color-brand)" stopOpacity="0.9" />
          <stop offset="100%" stopColor="var(--color-brand)" stopOpacity="0" />
        </radialGradient>
        {/* Masks for each connector path so lights follow the line */}
        {connectors.map((c, i) => (
          <mask key={`mask-${i}`} id={`${PIPELINE_ID}-mask-${i}`}>
            <path d={c.d} stroke="white" strokeWidth="14" fill="none" />
          </mask>
        ))}
      </defs>

      {/* Connector lines */}
      <g
        stroke="var(--color-fd-border)"
        strokeWidth="2"
        fill="none"
        strokeDasharray="6 4"
      >
        {connectors.map((c, i) => (
          <path key={`conn-${i}`} d={c.d} />
        ))}
      </g>

      {/* Animated lights traveling along connectors */}
      {connectors.map((c, i) => (
        <g key={`light-${i}`} mask={`url(#${PIPELINE_ID}-mask-${i})`}>
          <circle r="18" fill={`url(#${PIPELINE_ID}-light-grad)`}>
            <animateMotion
              dur={`${1.8 + i * 0.3}s`}
              repeatCount="indefinite"
              path={c.d}
              keyPoints="0;1"
              keyTimes="0;1"
              calcMode="spline"
              keySplines="0.42 0 0.58 1"
            />
          </circle>
        </g>
      ))}

      {/* Stage nodes */}
      {stages.map((stage, i) => {
        const x = getNodeX(i);
        const isActive = i === activeIndex;
        return (
          <g
            key={stage.id}
            className="cursor-pointer"
            onMouseEnter={() => onStageHover(i)}
            onClick={() => onStageClick(i)}
            role="button"
            tabIndex={0}
            onKeyDown={(e) => {
              if (e.key === "Enter" || e.key === " ") onStageClick(i);
            }}
          >
            {/* Active glow */}
            {isActive && (
              <rect
                x={x - 3}
                y={NODE_Y - 3}
                width={NODE_WIDTH + 6}
                height={NODE_HEIGHT + 6}
                rx={NODE_RX + 2}
                fill="none"
                stroke="var(--color-brand)"
                strokeWidth="2"
                opacity="0.5"
              >
                <animate
                  attributeName="opacity"
                  values="0.3;0.6;0.3"
                  dur="2s"
                  repeatCount="indefinite"
                />
              </rect>
            )}
            {/* Node background */}
            <rect
              x={x}
              y={NODE_Y}
              width={NODE_WIDTH}
              height={NODE_HEIGHT}
              rx={NODE_RX}
              fill={isActive ? "var(--color-brand)" : "var(--color-fd-card)"}
              stroke={isActive ? "var(--color-brand)" : "var(--color-fd-border)"}
              strokeWidth="1.5"
              className="transition-colors duration-200"
            />
            {/* Node label */}
            <text
              x={x + NODE_WIDTH / 2}
              y={NODE_Y + NODE_HEIGHT / 2}
              textAnchor="middle"
              dominantBaseline="central"
              fill={
                isActive
                  ? "var(--color-brand-foreground)"
                  : "var(--color-fd-foreground)"
              }
              fontSize="14"
              fontWeight="600"
              fontFamily="var(--font-mono, ui-monospace, monospace)"
              className="pointer-events-none select-none"
            >
              {stage.label}
            </text>
            {/* Step number */}
            <text
              x={x + NODE_WIDTH / 2}
              y={NODE_Y - 10}
              textAnchor="middle"
              fill="var(--color-fd-muted-foreground)"
              fontSize="10"
              fontFamily="var(--font-mono, ui-monospace, monospace)"
              className="pointer-events-none select-none"
            >
              {i + 1}
            </text>
          </g>
        );
      })}
    </svg>
  );
}

function StageDescription({
  stage,
  index,
}: {
  stage: PipelineStage;
  index: number;
}) {
  return (
    <div className="flex flex-col gap-2">
      <div className="flex items-center gap-3">
        <span className="flex size-7 shrink-0 items-center justify-center rounded-full bg-brand text-brand-foreground text-xs font-bold">
          {index + 1}
        </span>
        <h4 className="text-lg font-semibold font-mono text-fd-foreground">
          {stage.label}
        </h4>
      </div>
      <p className="text-fd-muted-foreground text-sm leading-relaxed pl-10">
        {stage.description}
      </p>
    </div>
  );
}

export function PipelineAnimation({
  children,
}: {
  children?: ReactNode;
}) {
  const [activeIndex, setActiveIndex] = useState(0);

  return (
    <div className="col-span-full flex flex-col gap-8">
      {/* Section heading */}
      <h2
        className={cn(
          headingVariants({ variant: "h2" }),
          "text-brand text-center font-mono font-bold uppercase"
        )}
      >
        Under the hood.
      </h2>

      {/* Pipeline layout: SVG on top / left, description on right */}
      <div className="flex flex-col gap-8 lg:flex-row lg:items-start lg:gap-12">
        {/* Pipeline SVG */}
        <div className="flex-1 min-w-0">
          <div className="overflow-x-auto rounded-2xl border border-fd-border bg-fd-card/50 p-4 sm:p-6">
            <div className="min-w-[600px]">
              <PipelineSvg
                activeIndex={activeIndex}
                onStageHover={setActiveIndex}
                onStageClick={setActiveIndex}
              />
            </div>
          </div>
        </div>

        {/* Description panel */}
        <div className="flex flex-col gap-4 lg:w-[280px] lg:shrink-0 lg:pt-2">
          <StageDescription
            stage={stages[activeIndex]}
            index={activeIndex}
          />
          {/* Stage list for context */}
          <div className="flex flex-col gap-1 mt-2">
            {stages.map((stage, i) => (
              <button
                key={stage.id}
                type="button"
                className={cn(
                  "flex items-center gap-2 rounded-lg px-3 py-1.5 text-left text-sm transition-colors",
                  i === activeIndex
                    ? "bg-brand/10 text-brand font-medium"
                    : "text-fd-muted-foreground hover:text-fd-foreground hover:bg-fd-muted/50"
                )}
                onMouseEnter={() => setActiveIndex(i)}
                onClick={() => setActiveIndex(i)}
              >
                <span
                  className={cn(
                    "flex size-5 shrink-0 items-center justify-center rounded-full text-[10px] font-bold",
                    i === activeIndex
                      ? "bg-brand text-brand-foreground"
                      : "bg-fd-muted text-fd-muted-foreground"
                  )}
                >
                  {i + 1}
                </span>
                <span className="font-mono text-xs">{stage.label}</span>
              </button>
            ))}
          </div>
        </div>
      </div>

      {/* Caption */}
      <p className="text-center text-sm text-fd-muted-foreground max-w-2xl mx-auto">
        The CLI does the heavy lifting. Adapters receive clean, self-contained
        schemas and only handle code generation.
      </p>

      {/* CliAnimation slot */}
      {children}
    </div>
  );
}
