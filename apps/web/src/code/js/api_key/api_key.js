import { Sandbox } from 'khulnasoft'

const sandbox = await Sandbox.create({ apiKey: 'YOUR_API_KEY' })
await sandbox.close()
