import { createContext, useContext, useMemo } from 'react'

import type * as Comlink from 'comlink'
import { noop } from 'lodash'

import type { AuthenticatedUser } from '@sourcegraph/shared/src/auth'
import type { RepositoryMatch } from '@sourcegraph/shared/src/search/stream'

import type { ExtensionCoreAPI } from '../../contract'
import { KhulnasoftUri, type KhulnasoftUriOptionals } from '../../file-system/KhulnasoftUri'
import type { VSCodePlatformContext } from '../platform/context'

type MinimalRepositoryMatch = Pick<RepositoryMatch, 'repository' | 'branches' | 'description'>

export interface MatchHandlersContext {
    openRepo: (repository: MinimalRepositoryMatch) => void
    openFile: (repositoryName: string, optional?: KhulnasoftUriOptionals) => void
    openSymbol: (symbolUrl: string) => void
    openCommit: (commitUrl: string) => void
    instanceURL: string
}

export const MatchHandlersContext = createContext<MatchHandlersContext>({
    // Initialize in `SearchResultsView` (via `useMatchHandlers`)
    openRepo: noop,
    openFile: noop,
    openSymbol: noop,
    openCommit: noop,
    instanceURL: '',
})

export function useMatchHandlers({
    platformContext,
    extensionCoreAPI,
    onRepoSelected,
    authenticatedUser,
    instanceURL,
}: {
    platformContext: VSCodePlatformContext
    extensionCoreAPI: Comlink.Remote<ExtensionCoreAPI>
    onRepoSelected: (repositoryMatch: MinimalRepositoryMatch) => void
    authenticatedUser: AuthenticatedUser | null
    instanceURL: string
}): Omit<MatchHandlersContext, 'instanceURL'> {
    const host = useMemo(() => new URL(instanceURL).host, [instanceURL])

    const matchHandlers: Omit<MatchHandlersContext, 'instanceURL'> = useMemo(
        () => ({
            openRepo: repositoryMatch => {
                // noop, implementation in SearchResultsView component since the repo page depends on its state.
                // nvm, pass "onRepoSelected" prop
                onRepoSelected(repositoryMatch)

                extensionCoreAPI
                    .openKhulnasoftFile(`sourcegraph://${host}/${repositoryMatch.repository}`)
                    .catch(error => {
                        console.error('Error opening Khulnasoft repository', error)
                    })
                // Log View Event to sync search history
                // URL must be provided to render Recent Searches on Web
                platformContext.telemetryService.logPageView(
                    'Repository',
                    null,
                    authenticatedUser !== null,
                    `https://${host}/${repositoryMatch.repository}`
                )
            },
            openFile: (repositoryName, optionals) => {
                // Create sourcegraph URI
                const sourcegraphUri = KhulnasoftUri.fromParts(host, repositoryName, optionals)

                const uri = sourcegraphUri.uri + sourcegraphUri.positionSuffix()

                // Log View Event to sync search history
                platformContext.telemetryService.logPageView(
                    'Blob',
                    null,
                    authenticatedUser !== null,
                    sourcegraphUri.uri.replace('sourcegraph://', 'https://')
                )

                extensionCoreAPI
                    .openKhulnasoftFile(uri)
                    .catch(error => console.error('Error opening Khulnasoft file', error))
            },
            openSymbol: (symbolUrl: string) => {
                const {
                    path,
                    position,
                    revision,
                    repositoryName,
                    host: codeHost,
                } = KhulnasoftUri.parse(`https:/${symbolUrl}`, window.URL)
                const sourcegraphUri = KhulnasoftUri.fromParts(host, `${codeHost}/${repositoryName}`, {
                    revision,
                    path,
                    position: position
                        ? {
                              line: position.line - 1, // Convert to 1-based
                              character: position.character - 1,
                          }
                        : undefined,
                })
                const uri = sourcegraphUri.uri + sourcegraphUri.positionSuffix()

                extensionCoreAPI.openKhulnasoftFile(uri).catch(error => {
                    console.error('Error opening Khulnasoft file', error)
                })
            },
            openCommit: commitUrl => {
                const commitURL = new URL(commitUrl, instanceURL)
                extensionCoreAPI.openLink(commitURL.href).catch(error => {
                    console.error('Error opening commit in browser', error)
                })

                // Roadmap: open diff in VS Code instead of Khulnasoft Web.
            },
        }),
        [extensionCoreAPI, platformContext, authenticatedUser, onRepoSelected, host, instanceURL]
    )

    return matchHandlers
}

export function useOpenSearchResultsContext(): MatchHandlersContext {
    return useContext(MatchHandlersContext)
}
