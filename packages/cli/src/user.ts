import * as os from 'os'
import * as path from 'path'
import * as fs from 'fs'

export interface UserConfig {
  email: string
  accessToken: string
  defaultTeamApiKey?: string
  defaultTeamId?: string
  teamName: string
  teamId: string
  teamApiKey: string
  dockerProxySet?: boolean
}

export const USER_CONFIG_PATH = path.join(os.homedir(), '.khulnasoft', 'config.json') // TODO: Keep in Keychain
export const DOCS_BASE = process.env.KHULNASOFT_DOCS_BASE || 'https://khulnasoft.com/docs'

export function getUserConfig(): UserConfig | null {
  if (!fs.existsSync(USER_CONFIG_PATH)) return null
  return JSON.parse(fs.readFileSync(USER_CONFIG_PATH, 'utf8'))
}
