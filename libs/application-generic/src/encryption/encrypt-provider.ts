import { EncryptedSecret, ICredentialsDto, KHULNASOFT_ENCRYPTION_SUB_MASK, secureCredentials } from '@khulnasoft/shared';

import { decrypt, encrypt } from './cipher';

export function encryptSecret(text: string): EncryptedSecret {
  const encrypted = encrypt(text);

  return `${KHULNASOFT_ENCRYPTION_SUB_MASK}${encrypted}`;
}

export function decryptSecret(text: string | EncryptedSecret): string {
  let encryptedSecret = text;

  if (isKhulnasoftEncrypted(text)) {
    encryptedSecret = text.slice(KHULNASOFT_ENCRYPTION_SUB_MASK.length);
  }

  return decrypt(encryptedSecret);
}

export function encryptCredentials(credentials: ICredentialsDto): ICredentialsDto {
  const encryptedCredentials: ICredentialsDto = {};

  // eslint-disable-next-line guard-for-in
  for (const key in credentials) {
    encryptedCredentials[key] = isCredentialEncryptionRequired(key)
      ? encryptSecret(credentials[key])
      : credentials[key];
  }

  return encryptedCredentials;
}

export function decryptCredentials(credentials: ICredentialsDto): ICredentialsDto {
  const decryptedCredentials: ICredentialsDto = {};

  // eslint-disable-next-line guard-for-in
  for (const key in credentials) {
    decryptedCredentials[key] =
      typeof credentials[key] === 'string' && isKhulnasoftEncrypted(credentials[key])
        ? decryptSecret(credentials[key])
        : credentials[key];
  }

  return decryptedCredentials;
}

export function encryptApiKey(apiKey: string): EncryptedSecret {
  if (isKhulnasoftEncrypted(apiKey)) {
    return apiKey;
  }

  return encryptSecret(apiKey);
}

export function decryptApiKey(apiKey: string): string {
  if (isKhulnasoftEncrypted(apiKey)) {
    return decryptSecret(apiKey);
  }

  return apiKey;
}

function isKhulnasoftEncrypted(text: string): text is EncryptedSecret {
  return text.startsWith(KHULNASOFT_ENCRYPTION_SUB_MASK);
}

function isCredentialEncryptionRequired(key: string): boolean {
  return secureCredentials.some((secureCred) => secureCred === key);
}
