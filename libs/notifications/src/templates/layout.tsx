import { Body, Container, Head, Html, Img, Preview, Tailwind } from '@react-email/components';
import React from 'react';

interface IBaseEmailLayoutProps {
  previewText: string;
  children: React.ReactNode;
}

export function EmailLayout({ previewText, children }: IBaseEmailLayoutProps) {
  return (
    <Html>
      <Head />
      <Preview>{previewText}</Preview>
      <Tailwind>
        <Body className="mx-auto my-auto bg-white px-2 font-sans">
          <Container className="mx-auto my-[40px] max-w-[465px] rounded border border-solid border-[#eaeaea] p-[20px]">
            <Img
              src={`https://dashboard.khulnasoft.co/static/images/khulnasoft-colored-text.png`}
              width="100"
              height="37"
              alt="Khulnasoft"
              className="mx-auto my-[32px]"
            />
            {children}
          </Container>
        </Body>
      </Tailwind>
    </Html>
  );
}
