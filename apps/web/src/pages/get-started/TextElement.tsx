import { Text } from '@khulnasoft/khulnasofti';
import { css } from '@khulnasoft/khulnasofti/css';

export function TextElement({ children }) {
  return <Text className={css({ color: 'typography.text.secondary' })}>{children}</Text>;
}
