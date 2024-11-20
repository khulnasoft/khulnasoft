import khulnasoft

key = (lambda path: open(path).read())("key.pub")
khulnasoft.export("result", key)
