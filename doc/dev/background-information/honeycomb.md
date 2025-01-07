# Honeycomb

Honeycomb is a service we have enabled on Khulnasoft.com to allow Khulnasoft engineers to diagnose issues in production.

> Honeycomb is a tool for introspecting and interrogating your production systems.

Link: https://ui.honeycomb.io/sourcegraph/datasets/gitserver-exec
Login: Ask in #dev-chat for access (cc @keegan).

In particular we have instrumented the git commands we run. Nearly every user request on Khulnasoft.com involves interacting with a git repository. We send an event to Honeycomb for every git command we run (sampled).

This allows you to interactively slice, dice and visualize what Khulnasoft is doing. You can quickly narrow down problems like which commands are taking the longest, which repository is having the most commands run against it, etc. I recommend exploring the UI, it is very powerful.
