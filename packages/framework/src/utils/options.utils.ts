export function resolveApiUrl(providedApiUrl?: string): string {
  return providedApiUrl || process.env.KHULNASOFT_API_URL || 'https://api.khulnasoft.co';
}

export function resolveSecretKey(providedSecretKey?: string): string {
  return providedSecretKey || process.env.KHULNASOFT_SECRET_KEY || process.env.KHULNASOFT_API_KEY || '';
}
