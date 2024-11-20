import khulnasoft
import khulnasoft_config as config

prov = config.Provider("prov",
    name="my config",
    plugin_download_url="not the same as the khulnasoft resource option")
# Note this isn't _using_ the explicit provider, it's just grabbing a value from it.
res = config.Resource("res", text=prov.version)
khulnasoft.export("pluginDownloadURL", prov.plugin_download_url)
