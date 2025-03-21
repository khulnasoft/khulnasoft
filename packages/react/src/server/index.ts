export type * from '../utils/types';
export type { InboxProps } from '../components/Inbox';

/**
 * Exporting all components from the components folder
 * as empty functions to fix build errors in SSR
 * This will be replaced with actual components
 * when we implement the SSR components in @khulnasoft/js/ui
 */
export function Inbox() {}
export function InboxContent() {}
export function Notifications() {}
export function Preferences() {}
export function Bell() {}

// Hooks
export { KhulnasoftProvider } from '../index';
export * from '../hooks';
