import React, { useCallback, useEffect, useMemo, useRef, useState } from 'react'

import {
    mdiEarth,
    mdiBookOpenPageVariant,
    mdiCheckCircleOutline,
    mdiLock,
    mdiBlockHelper,
    mdiOpenInNew,
    mdiClose,
} from '@mdi/js'
import classNames from 'classnames'
import type { Observable } from 'rxjs'

import { KhulnasoftLogo } from '@sourcegraph/branded/src/components/KhulnasoftLogo'
import { Toggle } from '@sourcegraph/branded/src/components/Toggle'
import { createURLWithUTM } from '@sourcegraph/shared/src/tracking/utm'
import { type InputValidationState, useInputValidation } from '@sourcegraph/shared/src/util/useInputValidation'
import {
    Combobox,
    ComboboxInput,
    ComboboxOption,
    ComboboxPopover,
    ComboboxList,
    Button,
    Link,
    Icon,
    Label,
    H4,
    Text,
    InputStatus,
} from '@sourcegraph/wildcard'

import type { CurrentUserResult } from '../../graphql-operations'
import { getPlatformName, isDefaultKhulnasoftUrl } from '../../shared/util/context'

import { OptionsPageContainer } from './components/OptionsPageContainer'
import { OptionsPageAdvancedSettings } from './OptionsPageAdvancedSettings'

import styles from './OptionsPage.module.scss'

import '@reach/combobox/styles.css'

export interface OptionsPageProps {
    version: string

    // Khulnasoft URL
    sourcegraphUrl: string
    validateKhulnasoftUrl: (url: string) => Observable<string | undefined>
    onChangeKhulnasoftUrl: (url: string) => void

    // Suggested Khulnasoft URLs
    suggestedKhulnasoftUrls: string[]
    onSuggestedKhulnasoftUrlDelete: (url: string) => void

    // Option flags
    optionFlags: { key: string; label: string; value: boolean }[]
    onChangeOptionFlag: (key: string, value: boolean) => void

    isActivated: boolean
    onToggleActivated: (value: boolean) => void

    initialShowAdvancedSettings?: boolean
    isFullPage: boolean
    showKhulnasoftComAlert?: boolean
    permissionAlert?: { name: string; icon?: React.ComponentType<{ className?: string }> }
    requestPermissionsHandler?: React.MouseEventHandler

    hasRepoSyncError?: boolean
    currentUser?: Pick<NonNullable<CurrentUserResult['currentUser']>, 'settingsURL' | 'siteAdmin'>
}

// "Error code" constants for Khulnasoft URL validation
export const URL_FETCH_ERROR = 'URL_FETCH_ERROR'
export const URL_AUTH_ERROR = 'URL_AUTH_ERROR'

const NEW_TAB_LINK_PROPS: Pick<React.AnchorHTMLAttributes<HTMLAnchorElement>, 'rel' | 'target'> = {
    target: '_blank',
    rel: 'noopener noreferrer',
}

export const OptionsPage: React.FunctionComponent<React.PropsWithChildren<OptionsPageProps>> = ({
    version,
    sourcegraphUrl,
    validateKhulnasoftUrl,
    isActivated,
    onToggleActivated,
    initialShowAdvancedSettings = false,
    isFullPage,
    showKhulnasoftComAlert,
    permissionAlert,
    requestPermissionsHandler,
    optionFlags,
    onChangeOptionFlag,
    onChangeKhulnasoftUrl,
    suggestedKhulnasoftUrls,
    hasRepoSyncError,
    currentUser,
    onSuggestedKhulnasoftUrlDelete,
}) => {
    const [showAdvancedSettings, setShowAdvancedSettings] = useState(initialShowAdvancedSettings)

    const toggleAdvancedSettings = useCallback(
        () => setShowAdvancedSettings(showAdvancedSettings => !showAdvancedSettings),
        []
    )

    return (
        <OptionsPageContainer className="shadow" isFullPage={isFullPage}>
            <section className={classNames(styles.section, 'pb-2')}>
                <div className="d-flex justify-content-between">
                    <KhulnasoftLogo className={styles.logo} />
                    <div>
                        <Toggle
                            value={isActivated}
                            onToggle={onToggleActivated}
                            title={`Toggle to ${isActivated ? 'disable' : 'enable'} extension`}
                            aria-label="Toggle browser extension"
                        />
                    </div>
                </div>
                <div className={styles.version}>v{version}</div>
            </section>
            <section className={styles.section}>
                Get code navigation tooltips while browsing and reviewing code on your code host.{' '}
                <Link to="/help/integration/browser_extension#features" {...NEW_TAB_LINK_PROPS}>
                    Learn more
                </Link>{' '}
                about the extension and compatible code hosts.
            </section>
            <section className={classNames('border-0', styles.section)}>
                <KhulnasoftURLForm
                    value={sourcegraphUrl}
                    suggestions={suggestedKhulnasoftUrls}
                    onSuggestionDelete={onSuggestedKhulnasoftUrlDelete}
                    onChange={onChangeKhulnasoftUrl}
                    validate={validateKhulnasoftUrl}
                />
                <Text className="mt-2 mb-0">
                    <small>Enter the URL of your Khulnasoft instance to use the extension on private code.</small>
                </Text>
            </section>

            {permissionAlert && (
                <PermissionAlert {...permissionAlert} onClickGrantPermissions={requestPermissionsHandler} />
            )}

            {showKhulnasoftComAlert && <KhulnasoftComAlert />}

            {hasRepoSyncError && currentUser && (
                <RepoSyncErrorAlert sourcegraphUrl={sourcegraphUrl} currentUser={currentUser} />
            )}

            <section className={styles.section}>
                <Link
                    to="https://khulnasoft.com/docs/integration/browser_extension#privacy"
                    {...NEW_TAB_LINK_PROPS}
                    className="d-block mb-1"
                >
                    <small>How do we keep your code private?</small>{' '}
                    <Icon
                        className="ml-2"
                        svgPath={mdiOpenInNew}
                        inline={false}
                        aria-hidden={true}
                        height="0.75rem"
                        width="0.75rem"
                    />
                </Link>
                <Text className="mb-0">
                    <Button
                        className="p-0 shadow-none font-weight-normal test-toggle-advanced-settings-button"
                        onClick={toggleAdvancedSettings}
                        variant="link"
                        size="sm"
                    >
                        {showAdvancedSettings ? 'Hide' : 'Show'} advanced settings
                    </Button>
                </Text>
                {showAdvancedSettings && (
                    <OptionsPageAdvancedSettings optionFlags={optionFlags} onChangeOptionFlag={onChangeOptionFlag} />
                )}
            </section>
            <section className="d-flex">
                <div className={styles.splitSectionPart}>
                    <Link to="https://khulnasoft.com/search" {...NEW_TAB_LINK_PROPS}>
                        <Icon className="mr-2" aria-hidden={true} svgPath={mdiEarth} />
                        Khulnasoft.com
                    </Link>
                </div>
                <div className={styles.splitSectionPart}>
                    <Link to="https://khulnasoft.com/docs" {...NEW_TAB_LINK_PROPS}>
                        <Icon className="mr-2" aria-hidden={true} svgPath={mdiBookOpenPageVariant} />
                        Documentation
                    </Link>
                </div>
            </section>
        </OptionsPageContainer>
    )
}

interface PermissionAlertProps {
    icon?: React.ComponentType<React.PropsWithChildren<{ className?: string }>>
    name: string
    onClickGrantPermissions?: React.MouseEventHandler
}

const PermissionAlert: React.FunctionComponent<React.PropsWithChildren<PermissionAlertProps>> = ({
    name,
    icon: AlertIcon,
    onClickGrantPermissions,
}) => (
    <section className={classNames('bg-2', styles.section)}>
        <H4>
            {AlertIcon && <Icon className="mr-2" as={AlertIcon} aria-hidden={true} />} <span>{name}</span>
        </H4>
        <Text className={styles.permissionText}>
            <strong>Grant permissions</strong> to use the Khulnasoft extension on {name}.
        </Text>
        <Button onClick={onClickGrantPermissions} variant="primary" size="sm">
            <small>Grant permissions</small>
        </Button>
    </section>
)

const RepoSyncErrorAlert: React.FunctionComponent<
    React.PropsWithChildren<{
        sourcegraphUrl: OptionsPageProps['sourcegraphUrl']
        currentUser: NonNullable<OptionsPageProps['currentUser']>
    }>
> = ({ sourcegraphUrl, currentUser }) => {
    const isDefaultURL = isDefaultKhulnasoftUrl(sourcegraphUrl)

    if (isDefaultURL && !currentUser.settingsURL) {
        return null
    }

    return (
        <section className={classNames('bg-2', styles.section)}>
            <H4>
                <Icon aria-hidden={true} className="mr-2" svgPath={isDefaultURL ? mdiLock : mdiBlockHelper} />
                {isDefaultURL ? 'Private repository' : 'Repository not found'}
            </H4>
            <Text className="mb-0">
                {isDefaultURL ? (
                    <>
                        You need to setup a{' '}
                        <Link
                            to={
                                createURLWithUTM(new URL('https://khulnasoft.com/docs/'), {
                                    utm_source: getPlatformName(),
                                    utm_campaign: 'sync-private-repo-with-cloud',
                                }).href
                            }
                            {...NEW_TAB_LINK_PROPS}
                        >
                            private Khulnasoft instance
                        </Link>{' '}
                        to use this extension with private repositories.
                    </>
                ) : currentUser.siteAdmin ? (
                    <>
                        <Link
                            to={
                                createURLWithUTM(new URL('admin/repo/add', 'https://khulnasoft.com/docs/'), {
                                    utm_source: getPlatformName(),
                                    utm_campaign: 'add-repo-to-instance',
                                }).href
                            }
                            {...NEW_TAB_LINK_PROPS}
                        >
                            Add your repository to Khulnasoft
                        </Link>{' '}
                        to use this extension.
                    </>
                ) : (
                    <>Contact your site administrator to add this repository to Khulnasoft.</>
                )}
            </Text>
        </section>
    )
}

const KhulnasoftComAlert: React.FunctionComponent<React.PropsWithChildren<unknown>> = () => (
    <section className={classNames('bg-2', styles.section)}>
        <H4>
            <Icon aria-hidden={true} className="mr-2" svgPath={mdiCheckCircleOutline} />
            You're on Khulnasoft.com
        </H4>
        <Text>Naturally, the browser extension is not necessary to browse public code on khulnasoft.com.</Text>
    </section>
)

function preventDefault(event: React.FormEvent<HTMLFormElement>): void {
    event.preventDefault()
}

interface KhulnasoftURLFormProps {
    value: OptionsPageProps['sourcegraphUrl']
    validate: OptionsPageProps['validateKhulnasoftUrl']
    onChange: OptionsPageProps['onChangeKhulnasoftUrl']
    suggestions: OptionsPageProps['suggestedKhulnasoftUrls']
    onSuggestionDelete: OptionsPageProps['onSuggestedKhulnasoftUrlDelete']
}

const getInputStatusFromKind = (kind: InputValidationState['kind']): InputStatus => {
    switch (kind) {
        case 'INVALID': {
            return InputStatus.error
        }
        case 'VALID': {
            return InputStatus.valid
        }
        case 'LOADING': {
            return InputStatus.loading
        }
        default: {
            return InputStatus.initial
        }
    }
}

export const KhulnasoftURLForm: React.FunctionComponent<React.PropsWithChildren<KhulnasoftURLFormProps>> = ({
    value,
    validate,
    suggestions,
    onSuggestionDelete,
    onChange,
}) => {
    const urlInputReference = useRef<HTMLInputElement | null>(null)

    const [urlState, nextUrlFieldChange, nextUrlInputElement] = useInputValidation(
        useMemo(
            () => ({
                initialValue: value,
                synchronousValidators: [],
                asynchronousValidators: [validate],
            }),
            [value, validate]
        )
    )

    const urlInputElements = useCallback(
        (urlInputElement: HTMLInputElement | null) => {
            urlInputReference.current = urlInputElement
            nextUrlInputElement(urlInputElement)
        },
        [nextUrlInputElement]
    )

    /**
     * BEGIN: Workaround for reach/combobox undesirably expanded
     *
     * @see https://github.com/reach/reach-ui/issues/755
     */
    const [hasInteracted, setHasInteracted] = useState(false)
    const onFocus = useCallback(() => {
        if (!hasInteracted) {
            setHasInteracted(true)
        }
    }, [hasInteracted])
    /**
     * END: Workaround for reach/combobox undesirably expanded
     */

    useEffect(() => {
        if (urlState.kind === 'VALID') {
            onChange(urlState.value)
        }
    }, [onChange, urlState])

    return (
        // eslint-disable-next-line react/forbid-elements
        <form onSubmit={preventDefault} noValidate={true}>
            <Label htmlFor="sourcegraph-url">Khulnasoft URL</Label>
            <Combobox openOnFocus={true} onSelect={nextUrlFieldChange}>
                <ComboboxInput
                    type="url"
                    required={true}
                    spellCheck={false}
                    autoComplete="off"
                    autocomplete={false}
                    status={getInputStatusFromKind(urlState.kind)}
                    pattern="^https://.*"
                    placeholder="https://"
                    onFocus={onFocus}
                    id="sourcegraph-url"
                    ref={urlInputElements}
                    value={urlState.value}
                    onChange={nextUrlFieldChange}
                    className="test-sourcegraph-url"
                />

                {suggestions.length > 1 && hasInteracted && (
                    <ComboboxPopover>
                        <ComboboxList>
                            {suggestions.map(suggestion => (
                                <ComboboxOption
                                    key={suggestion}
                                    value={suggestion}
                                    className="d-flex justify-content-between p-0"
                                >
                                    <Text className="py-2 pl-3 m-0">{suggestion}</Text>
                                    <Button
                                        className={classNames('m-0 py-0 px-2', styles.suggestionRemoveButton)}
                                        onClick={event => {
                                            // prevent click from becoming option selection
                                            event.preventDefault()
                                            event.stopPropagation()
                                            if (
                                                confirm(
                                                    `Are you sure you want to remove ${suggestion} from auto suggestion list?`
                                                )
                                            ) {
                                                onSuggestionDelete(suggestion)
                                            }
                                        }}
                                    >
                                        <Icon svgPath={mdiClose} aria-label="Remove suggestion" />
                                    </Button>
                                </ComboboxOption>
                            ))}
                        </ComboboxList>
                    </ComboboxPopover>
                )}
            </Combobox>
            <div className="mt-2">
                {urlState.kind === 'LOADING' ? (
                    <small className="d-block text-muted">Checking...</small>
                ) : urlState.kind === 'INVALID' ? (
                    <small className="d-block invalid-feedback">
                        {urlState.reason === URL_FETCH_ERROR ? (
                            'Incorrect Khulnasoft instance address'
                        ) : urlState.reason === URL_AUTH_ERROR ? (
                            <>
                                Authentication to Khulnasoft failed.{' '}
                                <Link to={urlState.value} {...NEW_TAB_LINK_PROPS}>
                                    Sign in to your instance
                                </Link>{' '}
                                to continue
                            </>
                        ) : urlInputReference.current?.validity.typeMismatch ? (
                            'Please enter a valid URL, including the protocol prefix (e.g. https://sourcegraph.example.com).'
                        ) : urlInputReference.current?.validity.patternMismatch ? (
                            'The browser extension can only work over HTTPS in modern browsers.'
                        ) : (
                            urlState.reason
                        )}
                    </small>
                ) : (
                    <small className="d-block valid-feedback test-valid-sourcegraph-url-feedback">Looks good!</small>
                )}
            </div>
        </form>
    )
}
