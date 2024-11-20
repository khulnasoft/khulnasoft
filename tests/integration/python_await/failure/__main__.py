import asyncio
import khulnasoft

output = khulnasoft.Output.from_input(asyncio.sleep(1, []))
output.apply(lambda x: x[0])
