import { RiArrowRightSLine, RiGuideFill } from 'react-icons/ri';
import { RQBJsonLogic } from 'react-querybuilder';
import { Link } from 'react-router-dom';
import { StepResponseDto, WorkflowOriginEnum } from '@khulnasoft/shared';

import { Button } from '@/components/primitives/button';
import { Separator } from '@/components/primitives/separator';
import { SidebarContent } from '@/components/side-navigation/sidebar';
import { useConditionsCount } from '@/hooks/use-conditions-count';

export function SkipConditionsButton({
  origin,
  step,
  inSidebar = false,
}: {
  origin: WorkflowOriginEnum;
  step: StepResponseDto;
  inSidebar?: boolean;
}) {
  const canEditStepConditions = origin === WorkflowOriginEnum.KHULNASOFT_CLOUD;
  const uiSchema = step.controls.uiSchema;
  const skip = uiSchema?.properties?.skip;

  const conditionsCount = useConditionsCount(step.controls.values.skip as RQBJsonLogic);

  const button = (
    <Link to={'./conditions'} relative="path" state={{ stepType: step.type }}>
      <Button variant="secondary" mode="outline" className="flex w-full justify-start gap-1.5 text-xs font-medium">
        <RiGuideFill className="h-4 w-4 text-neutral-600" />
        Step Conditions
        {conditionsCount > 0 && (
          <span className="ml-auto flex items-center gap-0.5">
            <span>{conditionsCount}</span>
            <RiArrowRightSLine className="ml-auto h-4 w-4 text-neutral-600" />
          </span>
        )}
      </Button>
    </Link>
  );

  if (!skip || !canEditStepConditions) {
    return null;
  }

  if (!inSidebar) {
    return button;
  }

  return (
    <>
      <SidebarContent>{button}</SidebarContent>
      <Separator />
    </>
  );
}
