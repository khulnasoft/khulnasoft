import { Sandbox } from 'khulnasoft'

const sandbox = await Sandbox.create({
  template: 'base',
  envVars: { FOO: 'Hello' }, // $HighlightLine
})

await sandbox.close()
