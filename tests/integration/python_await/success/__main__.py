import asyncio
import khulnasoft

output = khulnasoft.Output.from_input(asyncio.sleep(1, "magic string"))
output.apply(print)
