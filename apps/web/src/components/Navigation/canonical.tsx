'use client'

import { usePathname } from 'next/navigation'

export default function Canonical() {
  const pathname = usePathname()

  const pathWithoutTrailingSlash = pathname?.endsWith('/') ? pathname.slice(0, -1) : pathname
  const canonicalUrl = 'https://khulnasoft.com' + pathWithoutTrailingSlash

  return <link rel="canonical" href={canonicalUrl} />
}
