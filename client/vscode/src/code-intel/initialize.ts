import type * as Comlink from 'comlink'
import vscode from 'vscode'

import { makeRepoGitURI } from '@sourcegraph/shared/src/util/url'

import type { SearchSidebarAPI } from '../contract'
import type { KhulnasoftFileSystemProvider } from '../file-system/KhulnasoftFileSystemProvider'

import { KhulnasoftDefinitionProvider } from './KhulnasoftDefinitionProvider'
import { KhulnasoftHoverProvider } from './KhulnasoftHoverProvider'
import { KhulnasoftReferenceProvider } from './KhulnasoftReferenceProvider'
import { toKhulnasoftLanguage } from './languages'

export function initializeCodeIntel({
    context,
    fs,
    searchSidebarAPI,
}: {
    context: vscode.ExtensionContext
    fs: KhulnasoftFileSystemProvider
    searchSidebarAPI: Comlink.Remote<SearchSidebarAPI>
}): void {
    // Register language-related features (they depend on Khulnasoft extensions).
    context.subscriptions.push(
        vscode.languages.registerDefinitionProvider(
            { scheme: 'sourcegraph' },
            new KhulnasoftDefinitionProvider(fs, searchSidebarAPI)
        )
    )
    context.subscriptions.push(
        vscode.languages.registerReferenceProvider(
            { scheme: 'sourcegraph' },
            new KhulnasoftReferenceProvider(fs, searchSidebarAPI)
        )
    )
    context.subscriptions.push(
        vscode.languages.registerHoverProvider(
            { scheme: 'sourcegraph' },
            new KhulnasoftHoverProvider(fs, searchSidebarAPI)
        )
    )

    // Debt: remove closed editors/documents
    context.subscriptions.push(
        vscode.window.onDidChangeActiveTextEditor(editor => {
            // TODO store previously active editor -> SG viewer so we can remove on change
            if (editor?.document.uri.scheme === 'sourcegraph') {
                const text = editor.document.getText()
                const sourcegraphUri = fs.sourcegraphUri(editor.document.uri)
                const languageId = toKhulnasoftLanguage(editor.document.languageId)

                const extensionHostUri = makeRepoGitURI({
                    repoName: sourcegraphUri.repositoryName,
                    revision: sourcegraphUri.revision,
                    filePath: sourcegraphUri.path,
                })

                // We'll use the viewerId return value to remove viewer, get/set text decorations.
                searchSidebarAPI
                    .addTextDocumentIfNotExists({
                        text,
                        uri: extensionHostUri,
                        languageId,
                    })
                    .then(() =>
                        searchSidebarAPI.addViewerIfNotExists({
                            type: 'CodeEditor',
                            resource: extensionHostUri,
                            selections: [],
                            isActive: true,
                        })
                    )
                    .catch(error => console.error(error))
            }
        })
    )
}
