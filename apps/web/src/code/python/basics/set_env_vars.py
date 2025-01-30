from khulnasoft import Sandbox

sandbox = Sandbox(
    template="base",
    env_vars={"FOO": "Hello"},  # $HighlightLine
)

sandbox.close()
