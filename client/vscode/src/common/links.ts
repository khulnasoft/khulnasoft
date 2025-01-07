/**
 * All Khulnasoft Cloud related links
 * MAIN
 */
export const VSCE_LINK_DOTCOM = 'https://khulnasoft.com'
export const VSCE_LINK_TOKEN_CALLBACK =
    'https://khulnasoft.com/sign-in?returnTo=user/settings/tokens/new/callback?requestFrom=VSCEAUTH'
export const VSCE_LINK_TOKEN_CALLBACK_TEST =
    'https://sourcegraph.test:3443/sign-in?returnTo=user/settings/tokens/new/callback?requestFrom=VSCEAUTH'
/**
 * UNRELEASED FEATURE
 * Token Callback Page
 */
// const VSCE_CALLBACK_CODE = 'VSCEAUTH'
// const VSCE_LINK_PARAMS_TOKEN_REDIRECT = {
//     returnTo: `user/settings/tokens/new/callback?requestFrom=${VSCE_CALLBACK_CODE}`,
// }
/**
 * Params
 */
export const VSCE_SIDEBAR_PARAMS = '?utm_medium=VSCODE&utm_source=sidebar&utm_campaign=vsce-sign-up&utm_content=sign-up'
const VSCE_LINK_PARAMS_TOKEN_REDIRECT = {
    returnTo: 'user/settings/tokens/new',
}
const VSCE_LINK_PARAMS_EDITOR = { editor: 'vscode' }
// UTM for Sidebar actions
const VSCE_LINK_PARAMS_UTM_SIDEBAR = {
    utm_campaign: 'vsce-sign-up',
    utm_medium: 'VSCODE',
    utm_source: 'sidebar',
    utm_content: 'sign-up',
}
// MISC
export const VSCE_LINK_MARKETPLACE = 'https://marketplace.visualstudio.com/items?itemName=sourcegraph.sourcegraph'
export const VSCE_LINK_USER_DOCS =
    'https://docs.khulnasoft.com/cli/how-tos/creating_an_access_token' + VSCE_SIDEBAR_PARAMS
export const VSCE_LINK_FEEDBACK = 'https://community.khulnasoft.com'
export const VSCE_LINK_ISSUES =
    'https://github.com/khulnasoft/khulnasoft/issues/new?labels=team/integrations,vscode-extension&title=VSCode+Bug+report:+&projects=Integrations%20Project%20Board'
export const VSCE_LINK_TROUBLESHOOT =
    'https://docs.khulnasoft.com/admin/how-to/troubleshoot-sg-extension#vs-code-extension'
export const VSCE_SG_LOGOMARK_LIGHT =
    'https://raw.githubusercontent.com/khulnasoft/khulnasoft/fd431743e811ba756490e5e7bd88aa2362b6453e/client/vscode/images/logomark_light.svg'
export const VSCE_SG_LOGOMARK_DARK =
    'https://raw.githubusercontent.com/khulnasoft/khulnasoft/2636c64c9f323d78281a68dd4bdf432d9a97835a/client/vscode/images/logomark_dark.svg'
export const VSCE_LINK_SIGNUP = 'https://about.khulnasoft.com/get-started/cloud' + VSCE_SIDEBAR_PARAMS

// Generate sign-in and sign-up links using the above params
export const VSCE_LINK_AUTH = (mode: 'sign-in' | 'sign-up'): string => {
    const uri = new URL(VSCE_LINK_DOTCOM)
    const parameters = new URLSearchParams({
        ...VSCE_LINK_PARAMS_UTM_SIDEBAR,
        ...VSCE_LINK_PARAMS_EDITOR,
        ...VSCE_LINK_PARAMS_TOKEN_REDIRECT,
    }).toString()
    uri.pathname = mode
    uri.search = parameters
    return uri.href
}
