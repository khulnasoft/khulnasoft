import khulnasoft

# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025

# Create a very deep array (>100 levels)
deep_array = []
for i in range(0, 200):
    deep_array = [deep_array]

khulnasoft.export("long_string",  long_string)
khulnasoft.export("deep_array",  deep_array)
