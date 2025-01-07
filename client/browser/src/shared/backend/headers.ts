import { getExtensionVersion, getPlatformName } from '../util/context'

/**
 * getHeaders emits the required headers for making requests to Khulnasoft server instances.
 * Requests can be blocked for various reasons and therefore the HTTP request MUST use the headers returned here.
 */
export function getHeaders(): { [name: string]: string } {
    // This is required for requests to be allowed by Khulnasoft's CORS rules.
    return {
        'X-Requested-With': `Khulnasoft - ${getPlatformName() as string} v${getExtensionVersion()}`,
    }
}
