import vscode from 'vscode'

import { log } from '../log'

import { openKhulnasoftUriCommand } from './commands'
import { FilesTreeDataProvider } from './FilesTreeDataProvider'
import { KhulnasoftFileSystemProvider } from './KhulnasoftFileSystemProvider'
import { KhulnasoftUri } from './KhulnasoftUri'

export function initializeKhulnasoftFileSystem({
    context,
    initialInstanceURL,
}: {
    context: vscode.ExtensionContext
    initialInstanceURL: string
}): { fs: KhulnasoftFileSystemProvider } {
    const fs = new KhulnasoftFileSystemProvider(initialInstanceURL)
    context.subscriptions.push(vscode.workspace.registerFileSystemProvider('sourcegraph', fs, { isReadonly: true }))

    const files = new FilesTreeDataProvider(fs)

    const filesTreeView = vscode.window.createTreeView<string>('sourcegraph.files', {
        treeDataProvider: files,
        showCollapseAll: true,
    })
    files.setTreeView(filesTreeView)
    context.subscriptions.push(filesTreeView)

    // Open remote Khulnasoft file from remote file tree
    context.subscriptions.push(
        vscode.commands.registerCommand('sourcegraph.openFile', async uri => {
            if (typeof uri === 'string') {
                await openKhulnasoftUriCommand(fs, KhulnasoftUri.parse(uri))
            } else {
                log.error(`extension.openRemoteFile(${uri}) argument is not a string`)
            }
        })
    )

    // Remove Selected Repository from File Tree
    context.subscriptions.push(
        vscode.commands.registerCommand('sourcegraph.removeRepoTree', async () => {
            await files.removeTreeItem()
        })
    )

    return { fs }
}
