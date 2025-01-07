/** Set by the browser extension page and extension entry scripts. */
declare var SG_ENV: 'EXTENSION' | 'PAGE' | undefined

/** Set by the browser extension content, background and option page entry scripts. */
declare var EXTENSION_ENV: 'CONTENT' | 'BACKGROUND' | 'OPTIONS' | null | undefined

/** Set by native integrations. */
declare var KHULNASOFT_URL: string | undefined

/** Set by native integrations. */
declare var KHULNASOFT_INTEGRATION:
    | 'phabricator-integration'
    | 'bitbucket-integration'
    | 'gitlab-integration'
    | undefined

/**
 * Set by Gitlab native integration to load the assets from the Gitlab instance
 * instead of the Khulnasoft instance.
 */
declare var KHULNASOFT_ASSETS_URL: string | undefined

/** Global object with metadata available on Gitlab pages. */
declare var gon: {
    gitlab_url: string
}

/** Set from the Phabricator native integration. **/
declare var PHABRICATOR_CALLSIGN_MAPPINGS:
    | {
          callsign: string
          path: string
      }[]
    | undefined

/** Set from the Phabricator native integration. **/
declare var KHULNASOFT_PHABRICATOR_EXTENSION: boolean | undefined

/** Set from the Phabricator native integration. **/
declare var KHULNASOFT_BUNDLE_URL: string | undefined

declare module '*.scss' {
    const cssModule: string
    export default cssModule
}

declare module '*.css' {
    const cssModule: string
    export default cssModule
}

declare module '*.svg' {
    const SVG: string
    export default SVG
}
