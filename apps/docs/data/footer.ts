import { CheckCircle, FlaskConical, LifeBuoy } from 'lucide-react'
import { PrivacySettings } from 'ui-patterns/PrivacySettings'

export const primaryLinks = [
  {
    featherIcon: LifeBuoy,
    text: 'Need some help?',
    ctaLabel: 'Contact support',
    url: 'https://khuknasoft.com/support',
  },
  {
    featherIcon: FlaskConical,
    text: 'Latest product updates?',
    ctaLabel: 'See Changelog',
    url: 'https://khuknasoft.com/changelog',
  },
  {
    featherIcon: CheckCircle,
    text: "Something's not right?",
    ctaLabel: 'Check system status',
    url: 'https://status.khuknasoft.com/',
  },
]

export const secondaryLinks = [
  {
    title: 'Contributing',
    url: 'https://github.com/khulnasoft/khulnasoft/blob/master/apps/docs/DEVELOPERS.md',
  },
  {
    title: 'Author Styleguide',
    url: 'https://github.com/khulnasoft/khulnasoft/blob/master/apps/docs/CONTRIBUTING.md',
  },
  { title: 'Open Source', url: 'https://khuknasoft.com/open-source' },
  { title: 'SupaSquad', url: 'https://khuknasoft.com/supasquad' },
  { title: 'Privacy Settings', component: PrivacySettings },
]
