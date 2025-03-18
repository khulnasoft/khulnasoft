import dotenv from 'dotenv';
import path from 'node:path';
import { getContextPath, KhulnasoftComponentEnum, getEnvFileNameForNodeEnv } from '@khulnasoft/shared';

dotenv.config({ path: path.join(__dirname, '..', getEnvFileNameForNodeEnv(process.env.NODE_ENV)) });

export const CONTEXT_PATH = getContextPath(KhulnasoftComponentEnum.WEBHOOK);
