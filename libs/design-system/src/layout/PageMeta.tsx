import { Helmet } from 'react-helmet-async';

type Props = {
  title?: string;
};

export function PageMeta({ title }: Props) {
  return (
    <Helmet>
      <title>{title ? `${title} | ` : ``}Khulnasoft Cloud Dashboard</title>
    </Helmet>
  );
}
