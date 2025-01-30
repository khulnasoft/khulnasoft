import * as listen from 'async-listen'
import * as commander from 'commander'
import * as fs from 'fs'
import * as http from 'http'
import * as open from 'open'
import * as path from 'path'
import * as khulnasoft from 'khulnasoft'

import { pkg } from 'src'
import {
  DOCS_BASE,
  getUserConfig,
  USER_CONFIG_PATH,
  UserConfig,
} from 'src/user'
import { asBold, asFormattedConfig, asFormattedError } from 'src/utils/format'
import { connectionConfig } from 'src/api'
import { handleKHULNASOFTRequestError } from '../../utils/errors'

export const loginCommand = new commander.Command('login')
  .description('log in to CLI')
  .action(async () => {
    let userConfig
    try {
      userConfig = getUserConfig()
    } catch (err) {
      console.error(asFormattedError('Failed to read user config', err))
    }
    if (userConfig) {
      console.log(
        `\nAlready logged in. ${asFormattedConfig(
          userConfig
        )}.\n\nIf you want to log in as a different user, log out first by running 'khulnasoft auth logout'.\nTo change the team, run 'khulnasoft auth configure'.\n`
      )
      return
    } else if (!userConfig) {
      console.log('Attempting to log in...')
      userConfig = await signInWithBrowser()
      if (!userConfig) {
        console.info('Login aborted')
        return
      }

      const signal = connectionConfig.getSignal()
      const config = new khulnasoft.ConnectionConfig({
        accessToken: process.env.KHULNASOFT_ACCESS_TOKEN || userConfig?.accessToken,
        apiKey: process.env.KHULNASOFT_API_KEY || userConfig?.teamApiKey,
      })
      const client = new khulnasoft.ApiClient(config)
      const res = await client.api.GET('/teams', { signal })

      handleKHULNASOFTRequestError(res.error, 'Error getting teams')

      const defaultTeam = res.data.find(
        (team: khulnasoft.components['schemas']['Team']) => team.isDefault
      )
      if (!defaultTeam) {
        console.error(
          asFormattedError('No default team found, please contact support')
        )
        process.exit(1)
      }

      userConfig.teamName = defaultTeam.name
      userConfig.teamId = defaultTeam.teamID
      userConfig.teamApiKey = defaultTeam.apiKey
      fs.mkdirSync(path.dirname(USER_CONFIG_PATH), { recursive: true })
      fs.writeFileSync(USER_CONFIG_PATH, JSON.stringify(userConfig, null, 2))
    }

    console.log(
      `Logged in as ${asBold(userConfig.email)} with selected team ${asBold(
        userConfig.teamName
      )}`
    )
    process.exit(0)
  })

async function signInWithBrowser(): Promise<UserConfig> {
  const server = http.createServer()
  const { port } = await listen.default(server, 0, '127.0.0.1')
  const loginUrl = new URL(`${DOCS_BASE}/api/cli`)
  loginUrl.searchParams.set('next', `http://localhost:${port}`)
  loginUrl.searchParams.set('cliVersion', pkg.version)

  return new Promise((resolve, reject) => {
    server.once('request', (req, res) => {
      // Close the HTTP connection to prevent `server.close()` from hanging
      res.setHeader('connection', 'close')
      const followUpUrl = new URL(`${DOCS_BASE}/api/cli`)
      const searchParams = new URL(req.url || '/', 'http://localhost')
        .searchParams
      const searchParamsObj = Object.fromEntries(
        searchParams.entries()
      ) as unknown as UserConfig & {
        error?: string
      }
      const { error } = searchParamsObj
      if (error) {
        reject(new Error(error))
        followUpUrl.searchParams.set('state', 'error')
        followUpUrl.searchParams.set('error', error)
      } else {
        resolve(searchParamsObj)
        followUpUrl.searchParams.set('state', 'success')
        followUpUrl.searchParams.set('email', searchParamsObj.email!)
      }

      res.statusCode = 302
      res.setHeader('location', followUpUrl.href)
      res.end()
    })

    return open.default(loginUrl.toString())
  })
}
