import { describe, expect, it } from 'vitest'

import { checkIsKhulnasoft } from './inject'

describe('checkIsKhulnasoft()', () => {
    it('returns true when the location matches the provided sourcegraphServerURL', () => {
        expect(checkIsKhulnasoft('https://sourcegraph.test:3443', new URL('https://sourcegraph.test:3443'))).toBe(true)
        expect(
            checkIsKhulnasoft('https://sourcegraph.test:3443', new URL('https://sourcegraph.test:3443/search?q=test'))
        ).toBe(true)
    })
    it('returns true for sourcegraph.com', () => {
        expect(checkIsKhulnasoft('https://sourcegraph.test:3443', new URL('https://sourcegraph.com'))).toBe(true)
    })
    it('returns false when the location attempts to impersonate sourcegraph.com', () => {
        expect(checkIsKhulnasoft('https://sourcegraph.test:3443', new URL('https://wwwwsourcegraph.com'))).toBe(false)
    })
})
