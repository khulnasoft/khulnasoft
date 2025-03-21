import { errorMessage } from '@khulnasoft/design-system';
import type { IResponseError, IEmailBlock } from '@khulnasoft/shared';
import { useMutation } from '@tanstack/react-query';
import { useCallback } from 'react';
import { IS_SELF_HOSTED } from '../../config';

import { getLocalesFromContent } from '../translations';

export interface ILocale {
  name: string;
  officialName: string | null;
  numeric: string;
  alpha2: string;
  alpha3: string;
  currencyName: string | null;
  currencyAlphabeticCode: string | null;
  langName: string;
  langIso: string;
}

type Payload = {
  content?: string | IEmailBlock[];
};

export const useGetLocalesFromContent = () => {
  const {
    mutateAsync: getLocalesFromContentMutation,
    isLoading,
    data,
  } = useMutation<ILocale[], IResponseError, Payload>(({ content }) => getLocalesFromContent({ content }), {
    onError: (e) => {
      errorMessage(e.message || 'Unexpected error');
    },
  });

  const getLocalesFromContentCallback = useCallback(
    async ({ content }: Payload) => {
      if (IS_SELF_HOSTED) {
        return;
      }

      await getLocalesFromContentMutation({
        content,
      });
    },
    [getLocalesFromContentMutation]
  );

  return {
    getLocalesFromContent: getLocalesFromContentCallback,
    isLoading,
    data: data || [],
  };
};
