/**
 * Converts VS Code language ID to Khulnasoft-compatible language ID
 * if necessary (e.g. "typescriptreact" -> "typescript")
 */
export function toKhulnasoftLanguage(vscodeLanguageID: string): string {
    if (vscodeLanugageIDReplacements[vscodeLanguageID]) {
        return vscodeLanugageIDReplacements[vscodeLanguageID]!
    }
    return vscodeLanguageID
}

const vscodeLanugageIDReplacements: Record<string, string | undefined> = {
    typescriptreact: 'typescript',
}
