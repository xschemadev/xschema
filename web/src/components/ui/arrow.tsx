import { useEffect, useRef } from "react";

type ArrowOptions = {
    fromX: number;
    fromY: number;
    toX: number;
    toY: number;
    curveX?: number;
    curveY?: number;
    color?: string;
    strokeWidth?: number;
    tipSize?: number;
    tipAngle?: number;
    tipOffset?: number;
    cap?: "round" | "square" | "butt";
    padding?: number;
    className?: string;
    style?: React.CSSProperties;
};

const DEFAULT_PADDING = 16;
const DEFAULT_TIP_SIZE = 12;
const DEFAULT_TIP_OFFSET = 0;
const DEFAULT_STROKE_WIDTH = 1;
const DEFAULT_COLOR = "black";
const DEFAULT_CAP = "round";
const DEFAULT_TIP_ANGLE = Math.PI / 4 + Math.PI / 2;

function Arrow(props: ArrowOptions) {
    const canvasRef = useRef<HTMLCanvasElement>(null);

    useEffect(() => {
        if (!canvasRef.current) return;

        handleCanvas(canvasRef.current, props);
    }, [props]);

    return (
        <canvas ref={canvasRef} className={props.className} style={props.style} />
    );
}

function getCanvasDpr(_: CanvasRenderingContext2D) {
    const devicePixelRatio = window.devicePixelRatio || 1;

    const canvasPixelRatio =
        // @ts-expect-error
        _.webkitBackingStorePixelRatio ||
        // @ts-expect-error
        _.mozBackingStorePixelRatio ||
        // @ts-expect-error
        _.msBackingStorePixelRatio ||
        // @ts-expect-error
        _.oBackingStorePixelRatio ||
        // @ts-expect-error
        _.backingStorePixelRatio ||
        1;

    return devicePixelRatio / canvasPixelRatio;
}

function handleCanvas(canvas: HTMLCanvasElement, options: ArrowOptions) {
    const _ = canvas.getContext("2d")!;
    const dpr = getCanvasDpr(_);

    const minX = Math.min(options.fromX, options.toX);
    const maxX = Math.max(options.fromX, options.toX);
    const minY = Math.min(options.fromY, options.toY);
    const maxY = Math.max(options.fromY, options.toY);

    const padding = options.padding ?? DEFAULT_PADDING;
    const strokeWidth = options.strokeWidth ?? DEFAULT_STROKE_WIDTH;
    const tipSize = options.tipSize ?? DEFAULT_TIP_SIZE;
    const tipAngle = options.tipAngle ?? DEFAULT_TIP_ANGLE;
    const tipOffset = options.tipOffset ?? DEFAULT_TIP_OFFSET;
    const cap = options.cap ?? DEFAULT_CAP;
    const offset = padding - strokeWidth / 2;
    const hasCurve =
        typeof options.curveX === "number" && typeof options.curveY === "number";

    const width = maxX - minX + 2 * padding;
    const height = maxY - minY + 2 * padding;
    const aX = options.fromX - minX + offset;
    const aY = options.fromY - minY + offset;
    const bX = options.toX - minX + offset;
    const bY = options.toY - minY + offset;
    const curveX = hasCurve ? options.curveX! - minX + offset : aX;
    const curveY = hasCurve ? options.curveY! - minY + offset : aY;
    // Arrow tip
    const angle = Math.atan2(bY - curveY, bX - curveX);
    const cX = bX + tipSize * Math.cos(angle + tipAngle + tipOffset);
    const cY = bY + tipSize * Math.sin(angle + tipAngle + tipOffset);
    const dX = bX + tipSize * Math.cos(angle - tipAngle + tipOffset);
    const dY = bY + tipSize * Math.sin(angle - tipAngle + tipOffset);

    canvas.width = width * dpr;
    canvas.height = height * dpr;
    canvas.style.position = "absolute";
    canvas.style.top = `${minY - padding}px`;
    canvas.style.left = `${minX - padding}px`;
    canvas.style.width = `${width}px`;
    canvas.style.height = `${height}px`;
    // canvas.style.border = '1px solid lightgrey'

    _.scale(dpr, dpr);

    /* ---
    Draw
    --- */
    _.strokeStyle = options.color ?? DEFAULT_COLOR;
    _.lineWidth = strokeWidth;
    _.lineCap = cap;

    // Arrow body
    _.beginPath();
    _.moveTo(aX, aY);
    if (hasCurve) _.quadraticCurveTo(curveX, curveY, bX, bY);
    else _.lineTo(bX, bY);
    _.stroke();
    // Arrow tip
    _.beginPath();
    _.moveTo(bX, bY);
    _.lineTo(cX, cY);
    _.stroke();
    _.beginPath();
    _.moveTo(bX, bY);
    _.lineTo(dX, dY);
    _.stroke();
}

export { Arrow };