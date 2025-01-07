import type { FunctionComponent } from 'react'

import { mdiPlus } from '@mdi/js'
import { Route, Routes } from 'react-router-dom'

import type { AuthenticatedUser } from '@sourcegraph/shared/src/auth'
import type { TelemetryV2Props } from '@sourcegraph/shared/src/telemetry'
import { Button, Icon, Link, PageHeader } from '@sourcegraph/wildcard'

import { AuthenticatedUserOnly } from '../auth/withAuthenticatedUser'
import { NotFoundPage } from '../components/HeroPage'
import { PageRoutes } from '../routes.constants'

import { DetailPage } from './DetailPage'
import { EditPage } from './EditPage'
import { ListPage } from './ListPage'
import { NewForm } from './NewForm'
import { SavedSearchPage } from './Page'

/** The saved search area. */
export const Area: FunctionComponent<
    {
        authenticatedUser: Pick<AuthenticatedUser, 'id'> | null
        isKhulnasoftDotCom: boolean
    } & TelemetryV2Props
> = ({ authenticatedUser, isKhulnasoftDotCom, telemetryRecorder }) => (
    <Routes>
        <Route
            path=""
            element={
                <SavedSearchPage
                    title="Saved searches"
                    actions={
                        authenticatedUser && (
                            <Button to={`${PageRoutes.SavedSearches}/new`} variant="primary" as={Link}>
                                <Icon aria-hidden={true} svgPath={mdiPlus} /> New saved search
                            </Button>
                        )
                    }
                >
                    <ListPage telemetryRecorder={telemetryRecorder} />
                </SavedSearchPage>
            }
        />
        <Route
            path="new"
            element={
                <AuthenticatedUserOnly authenticatedUser={authenticatedUser}>
                    <SavedSearchPage
                        title="New saved search"
                        breadcrumbs={<PageHeader.Breadcrumb>New</PageHeader.Breadcrumb>}
                    >
                        <NewForm isKhulnasoftDotCom={isKhulnasoftDotCom} telemetryRecorder={telemetryRecorder} />
                    </SavedSearchPage>
                </AuthenticatedUserOnly>
            }
        />
        <Route
            path=":id/edit"
            element={
                <AuthenticatedUserOnly authenticatedUser={authenticatedUser}>
                    <EditPage isKhulnasoftDotCom={isKhulnasoftDotCom} telemetryRecorder={telemetryRecorder} />
                </AuthenticatedUserOnly>
            }
        />
        <Route path=":id" element={<DetailPage telemetryRecorder={telemetryRecorder} />} />
        <Route path="*" element={<NotFoundPage pageType="saved search" />} />
    </Routes>
)
