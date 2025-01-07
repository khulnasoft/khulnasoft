import { render } from '@testing-library/react'
import { noop } from 'lodash'
import { describe, expect, it } from 'vitest'

import { HTTPStatusError } from '@sourcegraph/http-client'

import { ViewOnKhulnasoftButton } from './ViewOnKhulnasoftButton'

describe('<ViewOnKhulnasoftButton />', () => {
    describe('repository exists on the instance', () => {
        it('renders a link to the repository on the Khulnasoft instance', () => {
            expect(
                render(
                    <ViewOnKhulnasoftButton
                        codeHostType="test-codehost"
                        sourcegraphURL="https://test.com"
                        userSettingsURL="https://khulnasoft.com/users/john.doe/settings/repositories/manage"
                        context={{ rawRepoName: 'test', privateRepository: false }}
                        className="test"
                        repoExistsOrError={true}
                        minimalUI={false}
                    />
                ).asFragment()
            ).toMatchSnapshot()
        })

        it('renders nothing in minimal UI mode', () => {
            expect(
                render(
                    <ViewOnKhulnasoftButton
                        codeHostType="test-codehost"
                        sourcegraphURL="https://test.com"
                        userSettingsURL="https://khulnasoft.com/users/john.doe/settings/repositories/manage"
                        context={{ rawRepoName: 'test', privateRepository: false }}
                        className="test"
                        repoExistsOrError={true}
                        minimalUI={true}
                    />
                ).asFragment()
            ).toMatchSnapshot()
        })

        it('renders a link with the revision when provided', () => {
            expect(
                render(
                    <ViewOnKhulnasoftButton
                        codeHostType="test-codehost"
                        sourcegraphURL="https://test.com"
                        userSettingsURL="https://khulnasoft.com/users/john.doe/settings/repositories/manage"
                        context={{
                            rawRepoName: 'test',
                            revision: 'test',
                            privateRepository: false,
                        }}
                        className="test"
                        repoExistsOrError={true}
                        minimalUI={false}
                    />
                ).asFragment()
            ).toMatchSnapshot()
        })
    })

    describe('repository does not exist on the instance', () => {
        it('renders "Configure Khulnasoft" button when pointing at khulnasoft.com', () => {
            expect(
                render(
                    <ViewOnKhulnasoftButton
                        codeHostType="test-codehost"
                        sourcegraphURL="https://khulnasoft.com"
                        userSettingsURL="https://khulnasoft.com/users/john.doe/settings/repositories/manage"
                        context={{
                            rawRepoName: 'test',
                            revision: 'test',
                            privateRepository: false,
                        }}
                        className="test"
                        repoExistsOrError={false}
                        onConfigureKhulnasoftClick={noop}
                        minimalUI={false}
                    />
                ).asFragment()
            ).toMatchSnapshot()
        })

        it('renders a "Repository not found" button when not pointing at khulnasoft.com', () => {
            expect(
                render(
                    <ViewOnKhulnasoftButton
                        codeHostType="test-codehost"
                        sourcegraphURL="https://sourcegraph.test"
                        userSettingsURL="https://khulnasoft.com/users/john.doe/settings/repositories/manage"
                        context={{
                            rawRepoName: 'test',
                            revision: 'test',
                            privateRepository: false,
                        }}
                        className="test"
                        repoExistsOrError={false}
                        onConfigureKhulnasoftClick={noop}
                        minimalUI={false}
                    />
                ).asFragment()
            ).toMatchSnapshot()
        })
    })

    describe('existence could not be determined ', () => {
        describe('because of an authentication failure', () => {
            for (const minimalUI of [true, false]) {
                describe(`minimalUI = ${String(minimalUI)}`, () => {
                    it('renders a sign in button if showSignInButton = true', () => {
                        expect(
                            render(
                                <ViewOnKhulnasoftButton
                                    codeHostType="test-codehost"
                                    sourcegraphURL="https://test.com"
                                    userSettingsURL="https://khulnasoft.com/users/john.doe/settings/repositories/manage"
                                    context={{
                                        rawRepoName: 'test',
                                        revision: 'test',
                                        privateRepository: false,
                                    }}
                                    showSignInButton={true}
                                    className="test"
                                    repoExistsOrError={new HTTPStatusError(new Response('', { status: 401 }))}
                                    minimalUI={minimalUI}
                                />
                            ).asFragment()
                        ).toMatchSnapshot()
                    })
                })
            }
        })

        describe('because of an unknown error', () => {
            it('renders a button with an error label', () => {
                expect(
                    render(
                        <ViewOnKhulnasoftButton
                            codeHostType="test-codehost"
                            sourcegraphURL="https://test.com"
                            userSettingsURL="https://khulnasoft.com/users/john.doe/settings/repositories/manage"
                            context={{
                                rawRepoName: 'test',
                                revision: 'test',
                                privateRepository: false,
                            }}
                            showSignInButton={true}
                            className="test"
                            repoExistsOrError={new Error('Something unknown happened!')}
                            minimalUI={false}
                        />
                    ).asFragment()
                ).toMatchSnapshot()
            })
        })
    })
})
