from khulnasoft import Sandbox

sandbox = Sandbox(template="base")

url = sandbox.get_host()  # $HighlightLine
print("https://" + url)

sandbox.close()
