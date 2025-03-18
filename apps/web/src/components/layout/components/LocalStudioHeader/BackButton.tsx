import { Text } from '@khulnasoft/khulnasofti';
import { IconOutlineArrowBack } from '@khulnasoft/khulnasofti/icons';
import { hstack } from '@khulnasoft/khulnasofti/patterns';

type BackButtonProps = { onClick: () => void; styles?: Record<string, any> };

export function BackButton({ onClick, styles = {} }: BackButtonProps) {
  return (
    <button
      className={hstack({
        cursor: 'pointer',
        gap: 'margins.icons.Icon20-txt',
        px: '75',
        py: '25',
        borderRadius: '75',
        textStyle: 'text.secondary !important',
        _hover: { bg: 'badge.border', '& p, & svg': { color: 'typography.text.main !important' } },
        ...styles,
      })}
      onClick={onClick}
    >
      <IconOutlineArrowBack />
      <Text variant="secondary" fontWeight={'strong'}>
        Back
      </Text>
    </button>
  );
}
