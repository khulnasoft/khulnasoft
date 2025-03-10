import { Sandbox } from 'khulnasoft'

const sandbox = await Sandbox.create({ template: 'base' })

// Create a new directory '/dir'
await sandbox.filesystem.makeDir('/dir') // $HighlightLine

await sandbox.close()
