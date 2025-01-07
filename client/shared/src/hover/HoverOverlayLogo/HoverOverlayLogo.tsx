import * as React from 'react'

import classNames from 'classnames'

import { KhulnasoftIcon } from '@sourcegraph/wildcard'

import styles from './HoverOverlayLogo.module.scss'

export const HoverOverlayLogo: React.FunctionComponent<React.PropsWithChildren<{ className?: string }>> = ({
    className,
}) => (
    <span className={classNames(styles.container, className)}>
        <KhulnasoftIcon className={styles.icon} />
    </span>
)
