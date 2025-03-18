import Title from '@/components/Title';
import { khulnasoftConfig } from '@/utils/config';
import { Inbox, Preferences } from '@khulnasoft/nextjs';

export default function Home() {
  return (
    <>
      <Title title="Preferences Component" />
      <div className="h-[600px] w-96 overflow-y-auto">
        <Inbox {...khulnasoftConfig}>
          <Preferences />
        </Inbox>
      </div>
    </>
  );
}
