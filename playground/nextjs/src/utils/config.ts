export const khulnasoftConfig = {
  applicationIdentifier: process.env.NEXT_PUBLIC_KHULNASOFT_APP_ID ?? '',
  subscriberId: process.env.NEXT_PUBLIC_KHULNASOFT_SUBSCRIBER_ID ?? '',
  backendUrl: process.env.NEXT_PUBLIC_KHULNASOFT_BACKEND_URL ?? 'http://localhost:3000',
  socketUrl: process.env.NEXT_PUBLIC_KHULNASOFT_SOCKET_URL ?? 'http://localhost:3002',
};
