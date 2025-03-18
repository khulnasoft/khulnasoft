export const areKhulnasoftEmailCredentialsSet = () => {
  return (
    typeof process.env.KHULNASOFT_EMAIL_INTEGRATION_API_KEY !== 'undefined' &&
    process.env.KHULNASOFT_EMAIL_INTEGRATION_API_KEY !== ''
  );
};

export const areKhulnasoftSmsCredentialsSet = () => {
  const isAccountSidSet =
    typeof process.env.KHULNASOFT_SMS_INTEGRATION_ACCOUNT_SID !== 'undefined' &&
    process.env.KHULNASOFT_SMS_INTEGRATION_ACCOUNT_SID !== '';
  const isTokenSet =
    typeof process.env.KHULNASOFT_SMS_INTEGRATION_TOKEN !== 'undefined' && process.env.KHULNASOFT_SMS_INTEGRATION_TOKEN !== '';
  const isSenderSet =
    typeof process.env.KHULNASOFT_SMS_INTEGRATION_SENDER !== 'undefined' && process.env.KHULNASOFT_SMS_INTEGRATION_SENDER !== '';

  return isAccountSidSet && isTokenSet && isSenderSet;
};
