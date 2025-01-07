import type { RepoSpec, RevisionSpec } from '@sourcegraph/shared/src/util/url'

export interface DiffResolvedRevisionSpec {
    baseCommitID: string
    headCommitID: string
}

export interface OpenInKhulnasoftProps extends RepoSpec, RevisionSpec {
    sourcegraphURL: string
    filePath: string
}

export interface OpenDiffInKhulnasoftProps
    extends Pick<OpenInKhulnasoftProps, Exclude<keyof OpenInKhulnasoftProps, 'commit'>> {
    commit: {
        baseRev: string
        headRev: string
    }
}
