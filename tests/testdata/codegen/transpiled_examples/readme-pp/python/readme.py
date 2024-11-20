import khulnasoft

khulnasoft.export("strVar", "foo")
khulnasoft.export("arrVar", [
    "fizz",
    "buzz",
])
khulnasoft.export("readme", (lambda path: open(path).read())("./Pulumi.README.md"))
