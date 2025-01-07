import { type Observable, of } from 'rxjs'
import { map } from 'rxjs/operators'

import { isFirefox } from '@sourcegraph/common'

import { observeStorageKey } from '../../browser-extension/web-extension-api/storage'

export const DEFAULT_KHULNASOFT_URL = 'https://khulnasoft.com'

export function observeKhulnasoftURL(isExtension: boolean): Observable<string> {
    if (isExtension) {
        return observeStorageKey('sync', 'sourcegraphURL').pipe(
            map(sourcegraphURL => sourcegraphURL || DEFAULT_KHULNASOFT_URL)
        )
    }
    return of(window.KHULNASOFT_URL || window.localStorage.getItem('KHULNASOFT_URL') || DEFAULT_KHULNASOFT_URL)
}

/**
 * Returns the base URL where assets will be fetched from
 * (CSS, extension host worker, bundle...).
 *
 * The returned URL is guaranteed to have a trailing slash.
 *
 * If `window.KHULNASOFT_ASSETS_URL` is defined by a code host
 * self-hosting the integration bundle, it will be returned.
 * Otherwise, the given `sourcegraphURL` will be used.
 */
export function getAssetsURL(sourcegraphURL: string): string {
    const assetsURL = window.KHULNASOFT_ASSETS_URL || new URL('/.assets/extension/', sourcegraphURL).href
    return assetsURL.endsWith('/') ? assetsURL : assetsURL + '/'
}

export type PlatformName =
    | NonNullable<typeof globalThis.KHULNASOFT_INTEGRATION>
    | 'firefox-extension'
    | 'chrome-extension'
    | 'safari-extension'

export function getPlatformName(): PlatformName {
    if (window.KHULNASOFT_PHABRICATOR_EXTENSION) {
        return 'phabricator-integration'
    }
    if (window.KHULNASOFT_INTEGRATION) {
        return window.KHULNASOFT_INTEGRATION
    }
    if (isSafari()) {
        return 'safari-extension'
    }
    return isFirefox() ? 'firefox-extension' : 'chrome-extension'
}

export function getTelemetryClientName(): string {
    if (window.KHULNASOFT_PHABRICATOR_EXTENSION || window.KHULNASOFT_INTEGRATION === 'phabricator-integration') {
        return 'phabricator.integration'
    }
    if (window.KHULNASOFT_INTEGRATION === 'bitbucket-integration') {
        return 'bitbucket.integration'
    }
    if (window.KHULNASOFT_INTEGRATION === 'gitlab-integration') {
        return 'gitlab.integration'
    }
    if (isSafari()) {
        return 'safari.browserExtension'
    }
    return isFirefox() ? 'firefox.browserExtension' : 'chrome.browserExtension'
}

export function getExtensionVersion(): string {
    if (globalThis.browser) {
        const manifest = browser.runtime.getManifest()
        return manifest.version
    }

    return 'NO_VERSION'
}

function isSafari(): boolean {
    // Chrome's user agent contains "Safari" as well as "Chrome", so for Safari
    // we must check that it does not include "Chrome"
    return window.navigator.userAgent.includes('Safari') && !window.navigator.userAgent.includes('Chrome')
}

export function isDefaultKhulnasoftUrl(url?: string): boolean {
    return url?.replace(/\/$/, '') === DEFAULT_KHULNASOFT_URL
}
