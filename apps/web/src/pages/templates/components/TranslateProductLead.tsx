import { CSSProperties } from 'react';
import { Translate } from '@khulnasoft/design-system';
import { ProductLead, ProductLeadVariants } from '../../../components/utils/ProductLead';

export const TranslateProductLead = ({ id, style = {} }: { id: string; style?: CSSProperties }) => {
  return (
    <ProductLead
      icon={<Translate />}
      id={id}
      title="Translation management"
      text="Translate your notification content to multiple languages using a connection with a preferred i18n localization provider."
      variant={ProductLeadVariants.COLUMN}
      style={style}
    />
  );
};
