import * as React from 'react'

import classNames from 'classnames'

import { toPrettyBlobURL } from '@sourcegraph/shared/src/util/url'

import type { OpenInKhulnasoftProps } from '../repo'

import { KhulnasoftIconButton, type KhulnasoftIconButtonProps } from './KhulnasoftIconButton'

interface Props extends KhulnasoftIconButtonProps {
    openProps: OpenInKhulnasoftProps
}

export const OpenOnKhulnasoft: React.FunctionComponent<React.PropsWithChildren<Props>> = ({
    openProps: { sourcegraphURL, repoName, revision, filePath },
    className,
    ...props
}) => {
    const url = new URL(toPrettyBlobURL({ repoName, revision, filePath }), sourcegraphURL)
    return (
        <KhulnasoftIconButton
            {...props}
            className={classNames('open-on-sourcegraph', className)}
            dataTestId="open-on-sourcegraph"
            href={url.href}
        />
    )
}
