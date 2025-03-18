import { Inbox } from '@khulnasoft/nextjs';
import Title from '@/components/Title';
import { khulnasoftConfig } from '@/utils/config';

export default function CustomSubjectBody() {
  return (
    <>
      <Title title="Custom Subject Body" />
      <Inbox
        {...khulnasoftConfig}
        renderSubject={(notification) => {
          return (
            <div>
              Subject: {notification.subject} {new Date().toISOString()}
            </div>
          );
        }}
        renderBody={(notification) => {
          return <div>Body: {notification.body}</div>;
        }}
      />
    </>
  );
}
