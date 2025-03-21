import styled from '@emotion/styled';

import { Text, AngularLogo, IframeLogo, ReactLogo, VueLogo } from '@khulnasoft/design-system';
import { FrameworkEnum } from '../../quick-start/consts';

const KhulnasoftInAppFrameworkHeaderHolder = styled.div`
  display: flex;
  align-items: center;
  gap: 12px;
`;

const FRAMEWORKS_MAP = {
  [FrameworkEnum.REACT]: { logo: ReactLogo, title: 'React integration guide' },
  [FrameworkEnum.ANGULAR]: { logo: AngularLogo, title: 'Angular integration guide' },
  [FrameworkEnum.VUE]: { logo: VueLogo, title: 'Vue integration guide' },
  [FrameworkEnum.JS]: { logo: IframeLogo, title: 'iFrame integration guide' },
};

export const KhulnasoftInAppFrameworkHeader = ({ framework }: { framework: FrameworkEnum | null }) => {
  if (!framework) return null;

  const { logo: Icon, title } = FRAMEWORKS_MAP[framework];

  return (
    <KhulnasoftInAppFrameworkHeaderHolder>
      <Icon />
      <Text style={{ fontSize: 20 }}>{title}</Text>
    </KhulnasoftInAppFrameworkHeaderHolder>
  );
};
