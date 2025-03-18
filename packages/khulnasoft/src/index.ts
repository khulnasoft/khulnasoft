#!/usr/bin/env node

import { Command } from 'commander';
import { v4 as uuidv4 } from 'uuid';
import { devCommand, DevCommandOptions } from './commands';
import { sync } from './commands/sync';
import { AnalyticService, ConfigService } from './services';
import { IInitCommandOptions, init } from './commands/init';
import { KHULNASOFT_API_URL, KHULNASOFT_SECRET_KEY } from './constants';

const analytics = new AnalyticService();
export const config = new ConfigService();
if (process.env.NODE_ENV === 'development') {
  config.clearStore();
}
const anonymousIdLocalState = config.getValue('anonymousId');
const anonymousId = anonymousIdLocalState || uuidv4();
const program = new Command();

program.name('khulnasoft').description(`A CLI tool to interact with Khulnasoft Cloud`);

program
  .command('sync')
  .description(
    `Sync your state with Khulnasoft Cloud

  Specifying the Bridge URL and Secret Key:
  (e.g., npx khulnasoft@latest sync -b https://acme.org/api/khulnasoft -s KHULNASOFT_SECRET_KEY)

  Sync with Khulnasoft Cloud in Europe:
  (e.g., npx khulnasoft@latest sync -b https://acme.org/api/khulnasoft -s KHULNASOFT_SECRET_KEY -a https://eu.api.khulnasoft.co)`
  )
  .usage('-b <url> -s <secret-key> [-a <url>]')
  .option('-a, --api-url <url>', 'The Khulnasoft Cloud API URL', KHULNASOFT_API_URL || 'https://api.khulnasoft.co')
  .requiredOption(
    '-b, --bridge-url <url>',
    'The Khulnasoft endpoint URL hosted in the Bridge application, by convention ends in /api/khulnasoft'
  )
  .requiredOption(
    '-s, --secret-key <secret-key>',
    'The Khulnasoft Secret Key. Obtainable at https://dashboard.khulnasoft.com/api-keys',
    KHULNASOFT_SECRET_KEY || ''
  )
  .action(async (options) => {
    analytics.track({
      identity: {
        anonymousId,
      },
      data: {},
      event: 'Sync Khulnasoft Endpoint State',
    });
    await sync(options.bridgeUrl, options.secretKey, options.apiUrl);
  });

program
  .command('dev')
  .description(
    `Start Khulnasoft Studio and a local tunnel

  Running the Bridge application on port 4000: 
  (e.g., npx khulnasoft@latest dev -p 4000)

  Running the Bridge application on a different route: 
  (e.g., npx khulnasoft@latest dev -r /v1/api/khulnasoft)
  
  Running with a custom tunnel:
  (e.g., npx khulnasoft@latest dev --tunnel https://my-tunnel.ngrok.app)`
  )
  .usage('[-p <port>] [-r <route>] [-o <origin>] [-d <dashboard-url>] [-sp <studio-port>] [-t <url>] [-H]')
  .option('-p, --port <port>', 'The local Bridge endpoint port', '4000')
  .option('-r, --route <route>', 'The Bridge endpoint route', '/api/khulnasoft')
  .option('-o, --origin <origin>', 'The Bridge endpoint origin')
  .option('-d, --dashboard-url <url>', 'The Khulnasoft Cloud Dashboard URL', 'https://dashboard.khulnasoft.co')
  .option('-sp, --studio-port <port>', 'The Local Studio server port', '2022')
  .option('-sh, --studio-host <host>', 'The Local Studio server host', 'localhost')
  .option('-t, --tunnel <url>', 'Self hosted tunnel. e.g. https://my-tunnel.ngrok.app')
  .option('-H, --headless', 'Run the Bridge in headless mode without opening the browser', false)
  .action(async (options: DevCommandOptions) => {
    analytics.track({
      identity: {
        anonymousId,
      },
      data: {},
      event: 'Open Dev Server',
    });

    return await devCommand(options, anonymousId);
  });

program
  .command('init')
  .description(`Create a new Khulnasoft application`)
  .option(
    '-s, --secret-key <secret-key>',
    `The Khulnasoft development environment Secret Key. Note that your Khulnasoft app won't work outside of local mode without it.`
  )
  .option('-a, --api-url <url>', 'The Khulnasoft Cloud API URL', 'https://api.khulnasoft.co')
  .action(async (options: IInitCommandOptions) => {
    return await init(options, anonymousId);
  });

program.parse(process.argv);
