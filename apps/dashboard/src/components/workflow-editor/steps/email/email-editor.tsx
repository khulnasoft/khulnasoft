import { Separator } from '@/components/primitives/separator';
import { getComponentByType } from '@/components/workflow-editor/steps/component-utils';
import { EmailPreviewHeader } from '@/components/workflow-editor/steps/email/email-preview';
import { EmailTabsSection } from '@/components/workflow-editor/steps/email/email-tabs-section';
import { UiSchemaGroupEnum, type UiSchema } from '@khulnasoft/shared';

type EmailEditorProps = { uiSchema: UiSchema };

export const EmailEditor = (props: EmailEditorProps) => {
  const { uiSchema } = props;

  if (uiSchema.group !== UiSchemaGroupEnum.EMAIL) {
    return null;
  }

  const { body, subject } = uiSchema.properties ?? {};

  return (
    <div className="flex h-full flex-col">
      <EmailTabsSection>
        <EmailPreviewHeader />
      </EmailTabsSection>
      <EmailTabsSection className="-mx-[2px] -my-[3px] px-7 py-2">
        {getComponentByType({ component: subject.component })}
      </EmailTabsSection>
      <Separator className="before:bg-neutral-100" />
      {/* extra padding on the left to account for the drag handle */}
      <EmailTabsSection className="basis-full bg-neutral-50 pl-16 pr-16 pt-5">
        {getComponentByType({ component: body.component })}
      </EmailTabsSection>
    </div>
  );
};
