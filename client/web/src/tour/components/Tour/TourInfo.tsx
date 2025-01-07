import React from 'react'

import classNames from 'classnames'

export const GETTING_STARTED_TOUR_MARKER = 'getting-started-tour-info-marker'

export interface TourInfoProps {
    className?: string
    isKhulnasoftDotCom: boolean
}

export const TourInfo: React.FunctionComponent<React.PropsWithChildren<TourInfoProps>> = ({
    className,
    isKhulnasoftDotCom,
}) => {
    if (!isKhulnasoftDotCom) {
        return null
    }
    return <div id={GETTING_STARTED_TOUR_MARKER} className={classNames('d-none', className)} />
}
