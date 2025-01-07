import mock from 'mock-require'
import { describe, it } from 'vitest'

// Stub Khulnasoft API
import { createStubKhulnasoftAPI } from '@sourcegraph/extension-api-stubs'

mock('sourcegraph', createStubKhulnasoftAPI())

describe('init', () => {
    it('placeholder', () => {})
})
