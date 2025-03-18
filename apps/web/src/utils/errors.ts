import type { IResponseError } from '@khulnasoft/shared';

export function parseServerErrorMessage(error: IResponseError | null): String {
  if (!error) {
    return '';
  }

  return Array.isArray(error?.message) ? error?.message[0] : error?.message;
}
