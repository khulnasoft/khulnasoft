import { createPortal } from 'react-dom';

import LogoKhulnasoft from 'src/components/logo-khulnasoft.svg?react';

export function LogoLoading() {
  return createPortal(
    <div className="col fixed inset-0 z-50 items-center justify-center bg-neutral">
      <LogoKhulnasoft className="max-h-24 animate-pulse" />
    </div>,
    document.getElementById('root') as HTMLElement,
  );
}
