import { ActionButton, IconThumbDownAlt, IconThumbUpAlt } from '@khulnasoft/design-system';
import { css } from '@khulnasoft/khulnasofti/css';
import { Flex, styled } from '@khulnasoft/khulnasofti/jsx';
import { text } from '@khulnasoft/khulnasofti/recipes';

const Text = styled('p', text);

export type Voting = 'up' | 'down';

const voteButtonClass = css({
  _selected: {
    '& svg': {
      color: 'typography.text.main !important',
    },
  },
});

export const VotingWidget = ({
  voted,
  onVoteClick,
}: {
  voted: Voting | undefined;
  onVoteClick: (vote: Voting) => () => void;
}) => {
  return (
    <Flex
      className={css({
        marginTop: '250',
      })}
      align="center"
      gap="125"
    >
      <Text>Did you find it useful?</Text>
      <Flex gap="50" align="center">
        <ActionButton
          className={voteButtonClass}
          aria-selected={voted === 'up'}
          onClick={onVoteClick('up')}
          Icon={() => <IconThumbUpAlt size={20} />}
        />
        <ActionButton
          className={voteButtonClass}
          aria-selected={voted === 'down'}
          onClick={onVoteClick('down')}
          Icon={() => <IconThumbDownAlt size={20} />}
        />
      </Flex>
    </Flex>
  );
};
