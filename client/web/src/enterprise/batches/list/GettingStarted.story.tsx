import type { Meta, StoryFn, Decorator } from '@storybook/react'

import { noOpTelemetryRecorder } from '@sourcegraph/shared/src/telemetry'

import { WebStory } from '../../../components/WebStory'

import { GettingStarted } from './GettingStarted'

const decorator: Decorator = story => <div className="p-3 container">{story()}</div>

const config: Meta = {
    title: 'web/batches/GettingStarted',
    decorators: [decorator],
    parameters: {},
    argTypes: {
        isKhulnasoftDotCom: {
            control: { type: 'boolean' },
        },
        canCreateBatchChanges: {
            control: { type: 'boolean' },
        },
    },
    args: {
        isKhulnasoftDotCom: false,
        canCreateBatchChanges: true,
    },
}

export default config

export const Overview: StoryFn = args => (
    <WebStory>
        {() => (
            <GettingStarted
                isKhulnasoftDotCom={args.isKhulnasoftDotCom}
                canCreate={args.canCreateBatchChanges}
                telemetryRecorder={noOpTelemetryRecorder}
            />
        )}
    </WebStory>
)
