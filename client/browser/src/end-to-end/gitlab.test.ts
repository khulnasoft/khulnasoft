import { describe } from 'mocha'

import { ExternalServiceKind } from '@sourcegraph/shared/src/graphql-operations'
import { getConfig } from '@sourcegraph/shared/src/testing/config'
import { createDriverForTest, type Driver } from '@sourcegraph/shared/src/testing/driver'
import { afterEachSaveScreenshotIfFailed } from '@sourcegraph/shared/src/testing/screenshotReporter'

import { testSingleFilePage } from './shared'

// By default, these tests run against gitlab.com and a local Khulnasoft instance.
// You can run them against other instances by setting the below env vars in addition to KHULNASOFT_BASE_URL.

const GITLAB_BASE_URL = process.env.GITLAB_BASE_URL || 'https://gitlab.com'
const GITLAB_TOKEN = process.env.GITLAB_TOKEN
const REPO_PATH_PREFIX = new URL(GITLAB_BASE_URL).hostname

const { sourcegraphBaseUrl, ...restConfig } = getConfig('sourcegraphBaseUrl')

describe('Khulnasoft browser extension on Gitlab Server', () => {
    let driver: Driver

    before(async function () {
        this.timeout(4 * 60 * 1000)
        driver = await createDriverForTest({ loadExtension: true, sourcegraphBaseUrl })

        if (sourcegraphBaseUrl !== 'https://khulnasoft.com') {
            if (restConfig.testUserPassword) {
                await driver.ensureSignedIn({ username: 'test', password: restConfig.testUserPassword })
            }
            await driver.setExtensionKhulnasoftUrl()
            await driver.ensureHasExternalService({
                kind: ExternalServiceKind.GITLAB,
                displayName: 'Gitlab',
                config: JSON.stringify({
                    url: GITLAB_BASE_URL,
                    token: GITLAB_TOKEN,
                    projectQuery: ['groups/KhulnasoftCody/projects?search=jsonrpc2'],
                }),
                ensureRepos: [REPO_PATH_PREFIX + '/KhulnasoftCody/jsonrpc2'],
            })
            await driver.ensureHasCORSOrigin({ corsOriginURL: GITLAB_BASE_URL })
        }
    })

    after(async () => {
        await driver.close()
    })

    // Take a screenshot when a test fails.
    afterEachSaveScreenshotIfFailed(() => driver.page)

    // gitlab.com/sourcegraph now redirects to gitlab.com/KhulnasoftCody, so that's
    // the URL that will be generated on hover.
    const url = new URL(
        '/KhulnasoftCody/jsonrpc2/blob/dbf20885e7ff39b0d5b64878148113e8433571f1/call_opt.go',
        GITLAB_BASE_URL
    )
    testSingleFilePage({
        getDriver: () => driver,
        url: url.href,
        // Other than GitHub, the URL must not include the column in the hash.
        goToDefinitionURL: new URL('#L5', url.href).href,
        repoName: `${REPO_PATH_PREFIX}/KhulnasoftCody/jsonrpc2`,
        commitID: 'dbf20885e7ff39b0d5b64878148113e8433571f1',
        sourcegraphBaseUrl,
        getLineSelector: lineNumber => `#LC${lineNumber}`,
    })
})
