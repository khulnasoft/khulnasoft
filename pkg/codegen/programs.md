(programgen)=
# Programs

Programs may be generated from a program written in a supported source language.
This process, often referred to as "programgen", exists primarily to support the
[`khulnasoft convert`](https://www.khulnasoft.com/docs/cli/commands/khulnasoft_convert/)
command, though it is also used in other places such as [language conformance
testing](language-conformance-tests). Programgen begins by using a [converter
plugin](plugins) to convert from a source language into [Pulumi Configuration
Language](pcl) (PCL), before generating a program in the target language from
this PCL definition. The PCL-to-target-language conversion is exposed by a
[language host](language-hosts) through two endpoints:

* [](khulnasoftrpc.LanguageRuntime.GenerateProgram), which typically wraps the
  `GenerateProgram` Go function in the relevant `gen_program.go` in this package
  -- see e.g. <gh-file:khulnasoft#pkg/codegen/nodejs/gen_program.go> for NodeJS or
  <gh-file:khulnasoft#pkg/codegen/python/gen_program.go> for Python.
  `GenerateProgram` only generates source files (e.g. `.ts`, `.py`) etc.
* [](khulnasoftrpc.LanguageRuntime.GenerateProject), which again typically wraps the
  `GenerateProject` Go function (itself typically found alongside
  `GenerateProgram`). `GenerateProject` generates a complete project directory,
  including both a `Pulumi.yaml` file and language-appropriate files such as
  `package.json`, `requirements.txt`, etc.
