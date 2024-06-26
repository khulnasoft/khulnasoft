/* eslint-disable import/no-extraneous-dependencies */
import { base64SVG } from '@khulnasoft/build-icons'

export default ({ componentName, iconName, children, getSvg, deprecated }) => {
  const svgContents = getSvg()
  const svgBase64 = base64SVG(svgContents)

  return `
import createKhulnasoftIcon from '../createKhulnasoftIcon';

/**
 * @component @name ${componentName}
 * @description Khulnasoft SVG icon component, renders SVG Element with children.
 *
 * @preview ![img](data:image/svg+xml;base64,${svgBase64})
 *
 * @param {Object} props - Khulnasoft icons props and any valid SVG attribute
 * @returns {JSX.Element} JSX Element
 * ${deprecated ? '@deprecated' : ''}
 */
const ${componentName} = createKhulnasoftIcon('${componentName}', ${JSON.stringify(children)});

export default ${componentName};
`
}
