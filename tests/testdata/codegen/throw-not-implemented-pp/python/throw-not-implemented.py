import khulnasoft


def not_implemented(msg):
    raise NotImplementedError(msg)

khulnasoft.export("result", not_implemented("expression here is not implemented yet"))
