import * as boxen from 'boxen'
import * as khulnasoft from 'khulnasoft'

import { getUserConfig, UserConfig } from './user'
import { asBold, asPrimary } from './utils/format'

export let apiKey = process.env.KHULNASOFT_API_KEY
export let accessToken = process.env.KHULNASOFT_ACCESS_TOKEN

const authErrorBox = boxen.default(
  `You must be logged in to use this command. Run ${asBold('khulnasoft auth login')}.

If you are seeing this message in CI/CD you may need to set the ${asBold(
    'KHULNASOFT_ACCESS_TOKEN'
  )} environment variable.
Visit ${asPrimary(
    'https://khulnasoft.com/docs/getting-started/api-key'
  )} to get the access token.`,
  {
    width: 70,
    float: 'center',
    padding: 0.5,
    margin: 1,
    borderStyle: 'round',
    borderColor: 'redBright',
  }
)

export function ensureAPIKey() {
  // If apiKey is not already set (either from env var or from user config), try to get it from config file
  if (!apiKey) {
    const userConfig = getUserConfig()
    apiKey = userConfig?.teamApiKey || userConfig?.defaultTeamApiKey
  }

  if (!apiKey) {
    console.error(authErrorBox)
    process.exit(1)
  } else {
    return apiKey
  }
}

export function ensureUserConfig(): UserConfig {
  const userConfig = getUserConfig()
  if (!userConfig) {
    console.error('No user config found, run `khulnasoft auth login` to log in first.')
    process.exit(1)
  }
  return userConfig
}

export function ensureAccessToken() {
  // If accessToken is not already set (either from env var or from user config), try to get it from config file
  if (!accessToken) {
    const userConfig = getUserConfig()
    accessToken = userConfig?.accessToken
  }

  if (!accessToken) {
    console.error(authErrorBox)
    process.exit(1)
  } else {
    return accessToken
  }
}

const userConfig = getUserConfig()

export const connectionConfig = new khulnasoft.ConnectionConfig({
  accessToken: process.env.KHULNASOFT_ACCESS_TOKEN || userConfig?.accessToken,
  apiKey: process.env.KHULNASOFT_API_KEY || userConfig?.teamApiKey,
})
export const client = new khulnasoft.ApiClient(connectionConfig)
