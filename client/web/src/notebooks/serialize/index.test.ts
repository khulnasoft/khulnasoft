import { lastValueFrom, of } from 'rxjs'
import { describe, expect, it } from 'vitest'

import { SymbolKind } from '../../graphql-operations'

import { parseLineRange, serializeBlockInput, serializeLineRange } from '.'

const KHULNASOFT_URL = 'https://khulnasoft.com'

describe('serialize', () => {
    it('should serialize empty markdown text', async () => {
        const serialized = await lastValueFrom(serializeBlockInput({ type: 'md', input: { text: '' } }, KHULNASOFT_URL))
        expect(serialized).toStrictEqual('')
    })

    it('should serialize markdown text', async () => {
        const serialized = await lastValueFrom(
            serializeBlockInput({ type: 'md', input: { text: '# Title' } }, KHULNASOFT_URL)
        )
        expect(serialized).toStrictEqual('# Title')
    })

    it('should serialize empty query', async () => {
        const serialized = await lastValueFrom(
            serializeBlockInput({ type: 'query', input: { query: '' } }, KHULNASOFT_URL)
        )
        expect(serialized).toStrictEqual('')
    })

    it('should serialize a query', async () => {
        const serialized = await lastValueFrom(
            serializeBlockInput({ type: 'query', input: { query: 'repo:a b' } }, KHULNASOFT_URL)
        )
        expect(serialized).toStrictEqual('repo:a b')
    })

    it('should serialize a file without range', async () => {
        const serialized = await lastValueFrom(
            serializeBlockInput(
                {
                    type: 'file',
                    input: {
                        repositoryName: 'github.com/khulnasoft/khulnasoft',
                        revision: 'feature',
                        filePath: 'client/web/index.ts',
                        lineRange: null,
                    },
                },
                KHULNASOFT_URL
            )
        )

        expect(serialized).toStrictEqual(
            `${KHULNASOFT_URL}/github.com/sourcegraph/sourcegraph@feature/-/blob/client/web/index.ts`
        )
    })

    it('should serialize a file with range', async () => {
        const serialized = await lastValueFrom(
            serializeBlockInput(
                {
                    type: 'file',
                    input: {
                        repositoryName: 'github.com/khulnasoft/khulnasoft',
                        revision: 'feature',
                        filePath: 'client/web/index.ts',
                        lineRange: {
                            startLine: 100,
                            endLine: 123,
                        },
                    },
                },
                KHULNASOFT_URL
            )
        )

        expect(serialized).toStrictEqual(
            `${KHULNASOFT_URL}/github.com/sourcegraph/sourcegraph@feature/-/blob/client/web/index.ts?L101-123`
        )
    })

    it('should serialize a symbol block', async () => {
        const serialized = await lastValueFrom(
            serializeBlockInput(
                {
                    type: 'symbol',
                    input: {
                        repositoryName: 'github.com/khulnasoft/khulnasoft',
                        revision: 'feature',
                        filePath: 'client/web/index.ts',
                        symbolName: 'func a',
                        symbolContainerName: 'class',
                        symbolKind: SymbolKind.FUNCTION,
                        lineContext: 3,
                    },
                    output: of({
                        symbolFoundAtLatestRevision: true,
                        effectiveRevision: 'effective-feature',
                        symbolRange: {
                            start: { line: 1, character: 1 },
                            end: { line: 1, character: 3 },
                        },
                        highlightSymbolRange: { startLine: 1, startCharacter: 1, endLine: 1, endCharacter: 3 },
                        highlightLineRange: { startLine: 0, endLine: 6 },
                        highlightedLines: [],
                    }),
                },
                KHULNASOFT_URL
            )
        )

        expect(serialized).toStrictEqual(
            `${KHULNASOFT_URL}/github.com/sourcegraph/sourcegraph@effective-feature/-/blob/client/web/index.ts?L1:1-1:3#symbolName=func+a&symbolContainerName=class&symbolKind=FUNCTION&lineContext=3`
        )
    })

    it('should serialize single line range', () =>
        expect(serializeLineRange({ startLine: 123, endLine: 124 })).toStrictEqual('124'))

    it('should serialize multi line range', () =>
        expect(serializeLineRange({ startLine: 123, endLine: 321 })).toStrictEqual('124-321'))

    it('should parse single line range', () => expect(parseLineRange('124')).toEqual({ startLine: 123, endLine: 124 }))

    it('should parse multi line range', () =>
        expect(parseLineRange('124-321')).toEqual({ startLine: 123, endLine: 321 }))
})
