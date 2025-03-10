import { Sandbox } from 'khulnasoft'

const sandbox = await Sandbox.create({
  template: 'base',
  cwd: '/code', // $HighlightLine
})

// You can also change the cwd of an existing sandbox
sandbox.cwd = '/home' // $HighlightLine

await sandbox.close()
