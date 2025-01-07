import { InfoTooltip, Input, InputEnd } from 'design-system';
import { CopyIconButton } from 'src/application/copy-icon-button';
import { Translate } from 'src/intl/translate';

import { ExternalLink } from './link';

const T = Translate.prefix('deployToKhulnasoftButton');

export function DeployToKhulnasoftButton({ deployUrl }: { deployUrl?: string }) {
  if (deployUrl === undefined) {
    return null;
  }

  const markdown = `[![Deploy to Khulnasoft](https://www.khulnasoft.com/static/images/deploy/button.svg)](${deployUrl})`;

  return (
    <div className="card col gap-4 p-4">
      <div className="row items-center gap-2 font-medium">
        <T id="title" />
        <InfoTooltip content={<T id="tooltip" />} />
      </div>

      <Input
        readOnly
        value={markdown}
        inputClassName="text-xs"
        end={
          <InputEnd>
            <CopyIconButton text={markdown} className="size-4" />
          </InputEnd>
        }
      />

      <ExternalLink openInNewTab href={deployUrl}>
        <img src="https://www.khulnasoft.com/static/images/deploy/button.svg" className="h-8" />
      </ExternalLink>
    </div>
  );
}
