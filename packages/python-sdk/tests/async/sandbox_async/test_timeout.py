import pytest

from time import sleep

from khulnasoft import AsyncSandbox


@pytest.mark.skip_debug()
async def test_shorten_timeout(async_sandbox: AsyncSandbox):
    await async_sandbox.set_timeout(5)
    sleep(6)

    is_running = await async_sandbox.is_running()
    assert is_running is False


@pytest.mark.skip_debug()
async def test_shorten_then_lengthen_timeout(async_sandbox: AsyncSandbox):
    await async_sandbox.set_timeout(5)
    sleep(1)
    await async_sandbox.set_timeout(10)
    sleep(6)
    await async_sandbox.is_running()
