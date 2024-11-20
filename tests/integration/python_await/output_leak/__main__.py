# Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

import asyncio
import khulnasoft

a = khulnasoft.Output.from_input([1, 2])
# This output has to await an asyncio.task
async def fn():
    await asyncio.sleep(1)
    return 42
b = khulnasoft.Output.from_input(asyncio.to_thread(fn))

# this asyncio task will run forever, we shouldn't block the program on that
async def loop():
    while True:
        await asyncio.sleep(1)
c = asyncio.create_task(loop())

# we should wait for this because it's an output apply, not just an asyncio task.
def printer(i: int):
    print("PRINT:", i)
d = b.apply(printer)

# but we only explicitly await for a
khulnasoft.export("export", a)

# all of the above should eventually resolve leaving the queue empty.
while True:
    # need to pump the event loop so the above outputs resolve.
    khulnasoft.runtime.sync_await._sync_await(asyncio.sleep(1))
    with khulnasoft.runtime.settings.SETTINGS.lock:
        if len(khulnasoft.runtime.settings.SETTINGS.outputs) == 0:
            break
