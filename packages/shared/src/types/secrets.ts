export const KHULNASOFT_ENCRYPTION_SUB_MASK = 'nvsk.';

export type EncryptedSecret = `${typeof KHULNASOFT_ENCRYPTION_SUB_MASK}${string}`;
