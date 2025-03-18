import { SignIn } from '@clerk/clerk-react';
import { PageMeta } from '@khulnasoft/design-system';
import AuthLayout from '../../../components/layout/components/AuthLayout';
import { ROUTES } from '../../../constants/routes';

export default function SignInPage() {
  return (
    <AuthLayout>
      <PageMeta title="Sign in" />
      <SignIn path={ROUTES.AUTH_LOGIN} signUpUrl={ROUTES.AUTH_SIGNUP} />
    </AuthLayout>
  );
}
