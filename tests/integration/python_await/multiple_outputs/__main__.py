import asyncio
import khulnasoft

output = khulnasoft.Output.from_input(asyncio.sleep(3, "magic string"))
output.apply(print)

exported = khulnasoft.Output.from_input(asyncio.sleep(2, "foo"))
khulnasoft.export("exported", exported)
exported.apply(print)

another = khulnasoft.Output.from_input(asyncio.sleep(5, "bar"))
another.apply(print)


