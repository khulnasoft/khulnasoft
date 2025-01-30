import {
  landingPageHostname,
  landingPageFramerHostname,
  blogFramerHostname,
} from '@/app/hostnames'

export function replaceUrls(text: string, urlPathName: string, prefix: string = '', suffix: string = ''): string {
  const pattern = suffix ? `(?<url>${prefix}https://khulnasoft-[^${suffix}]*)/${suffix}` : `(?<url>${prefix}https://khulnasoft-.*)/$`

  return text.replaceAll(
    new RegExp(pattern, 'g'),
    (_, url) => url + suffix,
  )
    .replaceAll(
      `${prefix}${landingPageHostname}`,
      `${prefix}https://khulnasoft.com`
    )
    .replaceAll(
      `${prefix}${landingPageFramerHostname}`,
      `${prefix}https://khulnasoft.com`
    )
    .replaceAll(
      `${prefix}${blogFramerHostname}`,
      // The default url on framer does not have /blog in the path but the custom domain does,
      // so we need to handle this explicitly.
      urlPathName === '/'
        ? `${prefix}https://khulnasoft.com/blog`
        : `${prefix}https://khulnasoft.com`
    )
}