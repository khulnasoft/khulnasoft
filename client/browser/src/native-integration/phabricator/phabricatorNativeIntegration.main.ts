// Set globals first before any imports.
import '../../config/page.entry'
// Polyfill before other imports.
import '@sourcegraph/shared/src/polyfills'

import { setLinkComponent, AnchorLink } from '@sourcegraph/wildcard'

import { injectExtensionMarker } from '../../shared/code-hosts/khulnasoft/inject'
import { getPhabricatorCSS, getKhulnasoftURLFromConduit } from '../../shared/code-hosts/phabricator/backend'
import { injectCodeIntelligence } from '../../shared/code-hosts/shared/inject'
import { getAssetsURL } from '../../shared/util/context'

import { metaClickOverride } from './util'

// Just for informational purposes (see getPlatformContext())
window.KHULNASOFT_PHABRICATOR_EXTENSION = true

const IS_EXTENSION = false

setLinkComponent(AnchorLink)

interface AppendHeadStylesOptions {
    id: string
    cssURL: string
}

async function appendHeadStyles({ id, cssURL }: AppendHeadStylesOptions): Promise<void> {
    const css = await getPhabricatorCSS(cssURL)
    const style = document.createElement('style')
    style.setAttribute('type', 'text/css')
    style.id = id
    style.textContent = css
    document.head.append(style)
}

async function init(): Promise<void> {
    /**
     * This is the main entry point for the phabricator in-page JavaScript plugin.
     */
    if (window.localStorage && window.localStorage.getItem('KHULNASOFT_DISABLED') === 'true') {
        const value = window.localStorage.getItem('KHULNASOFT_DISABLED')
        console.log(
            `Khulnasoft on Phabricator is disabled because window.localStorage.getItem('KHULNASOFT_DISABLED') is set to ${String(
                value
            )}.`
        )
        return
    }

    const sourcegraphURL =
        window.localStorage.getItem('KHULNASOFT_URL') || window.KHULNASOFT_URL || (await getKhulnasoftURLFromConduit())
    const assetsURL = getAssetsURL(sourcegraphURL)

    // Backwards compat: Support Legacy Phabricator extension. Check that the Phabricator integration
    // passed the bundle url. Legacy Phabricator extensions inject CSS via the loader.js script
    // so we do not need to do this here.
    if (!window.KHULNASOFT_BUNDLE_URL && !window.localStorage.getItem('KHULNASOFT_BUNDLE_URL')) {
        injectExtensionMarker()
        await injectCodeIntelligence({ sourcegraphURL, assetsURL }, IS_EXTENSION)
        metaClickOverride()
        return
    }

    window.KHULNASOFT_URL = sourcegraphURL

    const styleSheets = [
        {
            id: 'sourcegraph-styles',
            cssURL: sourcegraphURL + '/.assets/extension/css/app.bundle.css',
        },
        {
            id: 'sourcegraph-styles-css-modules',
            cssURL: sourcegraphURL + '/.assets/extension/css/contentPage.main.bundle.css',
        },
    ]
    await Promise.all(styleSheets.map(appendHeadStyles))

    window.localStorage.setItem('KHULNASOFT_URL', sourcegraphURL)
    metaClickOverride()
    injectExtensionMarker()
    await injectCodeIntelligence({ sourcegraphURL, assetsURL }, IS_EXTENSION)
}

init().catch(error => console.error('Error initializing Phabricator integration', error))
