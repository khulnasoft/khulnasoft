import type { Khulnasoft } from './khulnasoft';

export {};

declare global {
  const KHULNASOFT_API_VERSION: string;
  const PACKAGE_NAME: string;
  const PACKAGE_VERSION: string;
  interface Window {
    Khulnasoft: typeof Khulnasoft;
  }
}
