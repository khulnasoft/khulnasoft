import path from 'path'

/**
 * Unpack all `process.env.*` variables used during the build
 * time of the web application in this module to keep one source of truth.
 */
import { getEnvironmentBoolean, STATIC_ASSETS_PATH } from '@sourcegraph/build-config'

import { DEFAULT_SITE_CONFIG_PATH } from './constants'

const NODE_ENV = (process.env.NODE_ENV || 'development') as 'production' | 'development'

const NODE_DEBUG = process.env.NODE_DEBUG

export const IS_DEVELOPMENT = NODE_ENV === 'development'
export const IS_PRODUCTION = NODE_ENV === 'production'

export const ENVIRONMENT_CONFIG = {
    /**
     * ----------------------------------------
     * Build configuration.
     * ----------------------------------------
     */
    NODE_ENV,
    NODE_DEBUG,
    // Determines if build is running on CI.
    CI: getEnvironmentBoolean('CI'),
    // Determines if the build will be used for integration tests.
    // Can be used to expose global variables to integration tests (e.g., CodeMirror API).
    // Enabled in the dev environment to allow debugging integration tests with the dev server.
    INTEGRATION_TESTS: getEnvironmentBoolean('INTEGRATION_TESTS') || IS_DEVELOPMENT,

    WEB_BUILDER_SERVE_INDEX: getEnvironmentBoolean('WEB_BUILDER_SERVE_INDEX'),
    STATIC_ASSETS_PATH: process.env.STATIC_ASSETS_PATH || STATIC_ASSETS_PATH,

    // The commit SHA the client bundle was built with.
    COMMIT_SHA: process.env.COMMIT_SHA,
    // The current Docker image version, use to associate builds with Sentry's source maps.
    VERSION: process.env.VERSION,
    // Should sourcemaps be uploaded to Sentry.
    SENTRY_UPLOAD_SOURCE_MAPS: getEnvironmentBoolean('SENTRY_UPLOAD_SOURCE_MAPS'),
    // Sentry's Dotcom project's authentication token
    SENTRY_DOT_COM_AUTH_TOKEN: process.env.SENTRY_DOT_COM_AUTH_TOKEN,
    // Sentry organization
    SENTRY_ORGANIZATION: process.env.SENTRY_ORGANIZATION,
    // Sentry project
    SENTRY_PROJECT: process.env.SENTRY_PROJECT,

    /**
     * Omit slow deps (such as Monaco and GraphiQL) in the build to get a ~40% reduction in esbuild
     * rebuild time. The web app will show placeholders if features needing these deps are used.
     * (Esbuild only.)
     */
    DEV_WEB_BUILDER_OMIT_SLOW_DEPS: Boolean(process.env.DEV_WEB_BUILDER_OMIT_SLOW_DEPS),

    /** Disable code splitting for faster dev builds and dev page navigation. */
    DEV_WEB_BUILDER_NO_SPLITTING: Boolean(process.env.DEV_WEB_BUILDER_NO_SPLITTING),

    /**
     * ----------------------------------------
     * Application features configuration.
     * ----------------------------------------
     */
    KHULNASOFTDOTCOM_MODE: getEnvironmentBoolean('KHULNASOFTDOTCOM_MODE'),

    // Is reporting to Sentry enabled.
    ENABLE_SENTRY: getEnvironmentBoolean('ENABLE_SENTRY'),
    // Is OpenTelemetry instrumentation enabled.
    ENABLE_OPEN_TELEMETRY: getEnvironmentBoolean('ENABLE_OPEN_TELEMETRY'),

    /**
     * ----------------------------------------
     * Local environment configuration.
     * ----------------------------------------
     */
    KHULNASOFT_API_URL: process.env.KHULNASOFT_API_URL,
    KHULNASOFT_HTTPS_DOMAIN: process.env.KHULNASOFT_HTTPS_DOMAIN || 'sourcegraph.test',
    KHULNASOFT_HTTPS_PORT: Number(process.env.KHULNASOFT_HTTPS_PORT) || 3443,
    KHULNASOFT_HTTP_PORT: Number(process.env.KHULNASOFT_HTTP_PORT) || 3080,
    SITE_CONFIG_PATH: process.env.SITE_CONFIG_PATH || DEFAULT_SITE_CONFIG_PATH,
    CLIENT_OTEL_EXPORTER_OTLP_ENDPOINT: process.env.CLIENT_OTEL_EXPORTER_OTLP_ENDPOINT || '-/debug/otlp',
}

export type EnvironmentConfig = typeof ENVIRONMENT_CONFIG

const { KHULNASOFT_HTTPS_DOMAIN, KHULNASOFT_HTTPS_PORT, KHULNASOFT_HTTP_PORT } = ENVIRONMENT_CONFIG

export const HTTPS_WEB_SERVER_URL = `https://${KHULNASOFT_HTTPS_DOMAIN}:${KHULNASOFT_HTTPS_PORT}`
export const HTTP_WEB_SERVER_URL = `http://localhost:${KHULNASOFT_HTTP_PORT}`

export const STATIC_INDEX_PATH = path.resolve(ENVIRONMENT_CONFIG.STATIC_ASSETS_PATH, 'index.html')
