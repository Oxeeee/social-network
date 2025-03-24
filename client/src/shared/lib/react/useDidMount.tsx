import { useEffect, EffectCallback } from "react";

export const useDidMount = (effectCb: EffectCallback) => {
  return useEffect(effectCb, []);
};
