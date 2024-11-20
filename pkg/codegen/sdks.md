(sdkgen)=
# SDKs

*[Provider](providers) SDKs* ("software development kits") are generated from a
[Pulumi Schema](schema) definition. Often referred to as "SDKgen", this process
is used by the myriad providers supported by Pulumi to expose their resources,
components, and functions in an idiomatic way for a given language. SDKgen is
generally exposed through the [](khulnasoftrpc.LanguageRuntime.GeneratePackage)
method of a [language host](language-hosts), which in turn is exposed by the
CLI's [`khulnasoft package
gen-sdk`](https://www.khulnasoft.com/docs/cli/commands/khulnasoft_package_gen-sdk/)
command. At a code level, SDKgen starts with the relevant `GeneratePackage` Go
function in `gen.go` -- see <gh-file:khulnasoft#pkg/codegen/nodejs/gen.go> for
NodeJS, <gh-file:khulnasoft#pkg/codegen/python/gen.go> for Python, and so on.

:::{note}
The `khulnasoft package gen-sdk` command is not really intended to be used by
external users or customers, and instead offers a convenient interface for
generating provider SDKs as part of e.g. the various provider CI jobs used
to automate provider build and release processes.
:::
