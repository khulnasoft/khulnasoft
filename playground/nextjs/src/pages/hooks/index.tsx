import React from 'react';
import { KhulnasoftProvider } from '@khulnasoft/nextjs';
import { NotionTheme } from './_components/notion-theme';
import { khulnasoftConfig } from '../../utils/config';
import { StatusProvider } from './_components/status-context';

const Page = () => {
  return (
    <KhulnasoftProvider {...khulnasoftConfig}>
      <StatusProvider>
        <NotionTheme />
      </StatusProvider>
    </KhulnasoftProvider>
  );
};

export default Page;
