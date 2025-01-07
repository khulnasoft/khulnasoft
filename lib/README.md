# KhulnaSoft lib module

This directory is the root of a separate go module from the primary module rooted at khulnasoft/khulnasoft. This module exists to hold code that we want to reuse outside of the khulnasoft/khulnasoft repo.

Code in this module should _not_ import from khulnasoft/khulnasoft or from other KhulnaSoft repositories to avoid complicated dependency relationships. Instead consider moving code from elsewhere into this module.
