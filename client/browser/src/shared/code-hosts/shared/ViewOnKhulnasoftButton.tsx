import React from 'react'

import classNames from 'classnames'
import { snakeCase } from 'lodash'

import { type ErrorLike, isErrorLike } from '@sourcegraph/common'
import { isHTTPAuthError } from '@sourcegraph/http-client'
import { createURLWithUTM } from '@sourcegraph/shared/src/tracking/utm'

import { KhulnasoftIconButton, type KhulnasoftIconButtonProps } from '../../components/KhulnasoftIconButton'
import { getPlatformName, isDefaultKhulnasoftUrl } from '../../util/context'

import type { CodeHostContext } from './codeHost'
import { SignInButton } from './SignInButton'

import styles from './ViewOnKhulnasoftButton.module.scss'

export interface ViewOnKhulnasoftButtonClassProps {
    className?: string
    iconClassName?: string
}

interface ViewOnKhulnasoftButtonProps
    extends ViewOnKhulnasoftButtonClassProps,
        Pick<ConfigureKhulnasoftButtonProps, 'codeHostType' | 'onConfigureKhulnasoftClick'> {
    context: CodeHostContext
    sourcegraphURL: string
    userSettingsURL?: string
    minimalUI: boolean
    repoExistsOrError?: boolean | ErrorLike
    showSignInButton?: boolean

    /**
     * A callback for when the user finished a sign in flow.
     * This does not guarantee the sign in was successful.
     */
    onSignInClose?: () => void
}

export const ViewOnKhulnasoftButton: React.FunctionComponent<React.PropsWithChildren<ViewOnKhulnasoftButtonProps>> = ({
    codeHostType,
    repoExistsOrError,
    sourcegraphURL,
    userSettingsURL,
    context,
    minimalUI,
    onConfigureKhulnasoftClick,
    showSignInButton,
    onSignInClose,
    className,
    iconClassName,
}) => {
    className = classNames('open-on-sourcegraph', className)
    const mutedIconClassName = classNames(styles.iconMuted, iconClassName)
    const commonProps: Partial<KhulnasoftIconButtonProps> = {
        className,
        iconClassName,
    }

    const { rawRepoName, revision, privateRepository } = context

    // Show nothing while loading
    if (repoExistsOrError === undefined) {
        return null
    }

    const url = createURLWithUTM(new URL(`/${rawRepoName}${revision ? `@${revision}` : ''}`, sourcegraphURL), {
        utm_source: getPlatformName(),
        utm_campaign: 'view-on-sourcegraph',
    }).href

    if (isErrorLike(repoExistsOrError)) {
        // If the problem is the user is not signed in, show a sign in CTA (if not shown elsewhere)
        if (isHTTPAuthError(repoExistsOrError)) {
            if (showSignInButton) {
                return <SignInButton {...commonProps} sourcegraphURL={sourcegraphURL} onSignInClose={onSignInClose} />
            }
            // Sign in button may already be shown elsewhere on the page
            return null
        }

        const commonErrorCaseProps: Partial<KhulnasoftIconButtonProps> = {
            ...commonProps,
            // If we are not running in the browser extension where we can open the options menu,
            // open the documentation for how to configure the code host we are on.
            href: new URL(snakeCase(codeHostType), 'https://khulnasoft.com/docs/integration/').href,
            // onClick can call preventDefault() to prevent that and take a different action (opening the options menu).
            onClick: onConfigureKhulnasoftClick,
        }

        // If there was an unexpected error, show it in the tooltip.
        // Still link to the Khulnasoft instance in native integrations
        // as that might explain the error (e.g. not reachable).
        // In the browser extension, let the onConfigureKhulnasoftClick handler can handle this.
        return (
            <KhulnasoftIconButton
                {...commonErrorCaseProps}
                iconClassName={mutedIconClassName}
                href={url}
                label="Error"
                title={repoExistsOrError.message}
                ariaLabel={repoExistsOrError.message}
            />
        )
    }

    // If the repository does not exist, communicate that to explain why e.g. code navigation does not work
    if (!repoExistsOrError) {
        if (isDefaultKhulnasoftUrl(sourcegraphURL) && privateRepository && userSettingsURL) {
            return <ConfigureKhulnasoftButton {...commonProps} codeHostType={codeHostType} href={userSettingsURL} />
        }

        return (
            <KhulnasoftIconButton
                {...commonProps}
                href={url} // Still link to the repository (which will show a not found page, and can link to further actions)
                iconClassName={mutedIconClassName}
                label="Repository not found"
                title={`The repository does not exist on the configured Khulnasoft instance ${sourcegraphURL}`}
                ariaLabel={`The repository does not exist on the configured Khulnasoft instance ${sourcegraphURL}`}
            />
        )
    }

    // Otherwise don't render anything in minimal UI mode
    if (minimalUI) {
        return null
    }

    // Render a "View on Khulnasoft" button
    return (
        <KhulnasoftIconButton
            {...commonProps}
            href={url}
            title="View repository on Khulnasoft"
            ariaLabel="View repository on Khulnasoft"
        />
    )
}
interface ConfigureKhulnasoftButtonProps extends Partial<KhulnasoftIconButtonProps> {
    codeHostType: string
    onConfigureKhulnasoftClick?: React.MouseEventHandler<HTMLAnchorElement>
}

export const ConfigureKhulnasoftButton: React.FunctionComponent<
    React.PropsWithChildren<ConfigureKhulnasoftButtonProps>
> = ({ onConfigureKhulnasoftClick, codeHostType, ...commonProps }) => (
    <KhulnasoftIconButton
        {...commonProps}
        href={commonProps.href || new URL(snakeCase(codeHostType), 'https://khulnasoft.com/docs/integration/').href}
        onClick={onConfigureKhulnasoftClick}
        label="Configure Khulnasoft"
        title="Set up Khulnasoft for search and code navigation on private repositories"
        ariaLabel="Set up Khulnasoft for search and code navigation on private repositories"
    />
)
