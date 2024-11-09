import React from "react";

export type Children = { children: React.ReactNode };

export type StrictOmit<T, K extends keyof T> = Omit<T, K>;
