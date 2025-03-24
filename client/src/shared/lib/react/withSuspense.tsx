import { Suspense, SuspenseProps } from "react";

export const withSuspense = (
  component: SuspenseProps["children"],
  fallback: SuspenseProps["fallback"],
) => {
  return <Suspense fallback={fallback}>{component}</Suspense>;
};
