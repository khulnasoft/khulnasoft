import { PageMeta } from '@khulnasoft/design-system';
import { SignUpForm } from './components/SignUpForm';
import AuthLayout from '../../components/layout/components/AuthLayout';

const title = 'Sign up';

export default function SignUpPage() {
  return (
    <AuthLayout title={title}>
      <PageMeta title={title} />
      <SignUpForm />
    </AuthLayout>
  );
}
