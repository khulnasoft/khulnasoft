import asyncio
import khulnasoft


def unknownIfDryRun(value):
    if khulnasoft.runtime.is_dry_run():
        return khulnasoft.Output(resources=set(), future=fut(None), is_known=fut(False))
    return khulnasoft.Output.from_input(value)


def fut(x):
    f = asyncio.Future()
    f.set_result(x)
    return f
