import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as random from "@khulnasoft/random";

const foo = new random.RandomShuffle("foo", {inputs: [
    `just one
newline`,
    `foo
bar
baz
qux
quux
qux`,
    `{
    "a": 1,
    "b": 2,
    "c": [
      "foo",
      "bar",
      "baz",
      "qux",
      "quux"
    ]
}
`,
]});
