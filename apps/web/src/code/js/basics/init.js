import { Sandbox } from 'khulnasoft'

const sandbox = await Sandbox.create()

await sandbox.close()
