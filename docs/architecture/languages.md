(languages)=
(language-hosts)=
# Language hosts

*Language hosts*, or *language runtimes*, are the means by which Pulumi is able
to support a variety of programming languages. Officially, the term "language
host" refers to the combination of two parts:

* A runtime, which is a Pulumi [plugin](plugins) that exposes the ability to
  execute programs written in a particular language according to a [standardized
  gRPC interface](khulnasoftrpc.LanguageRuntime). The plugin will be named
  `khulnasoft-language-<language>` (e.g. `khulnasoft-language-nodejs` for NodeJS, or
  `khulnasoft-language-python` for Python).
* An SDK, which is a set of libraries that provide the necessary abstractions
  and utilities for writing Pulumi programs in that language (e.g.
  `@khulnasoft/khulnasoft` in NodeJS, or `khulnasoft` in Python).

Often however, the term "language host" is used to refer to the runtime alone.
Aside from providing the ability to [](khulnasoftrpc.LanguageRuntime.Run) programs,
the runtime also supports a number of other operations:

* *Code generation* methods enable callers to generate both [SDKs](sdkgen)
  ([](khulnasoftrpc.LanguageRuntime.GeneratePackage)) and [programs](programgen)
  ([](khulnasoftrpc.LanguageRuntime.GenerateProject)) in the language.
* *Query* endpoints allow callers to calculate the set of language-specific
  dependencies ([](khulnasoftrpc.LanguageRuntime.GetProgramDependencies)) or Pulumi
  plugins ([](khulnasoftrpc.LanguageRuntime.GetRequiredPlugins)) that might be
  required by a program.
* The *[](khulnasoftrpc.LanguageRuntime.Pack)* method allows callers to package up
  bundles of code written in the language into a format suitable for consumption
  by other code (for instance, packaging an SDK for use as a dependency, or
  packaging a program for execution).
