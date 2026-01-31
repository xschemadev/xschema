import useWindowSize from "./useWindowSize";

const MOBILE_BREAKPOINT = 768;

export function useIsMobile(breakpointProp?: number) {
    const breakpoint = breakpointProp || MOBILE_BREAKPOINT;
    const [width] = useWindowSize();
    return width < breakpoint;
}
