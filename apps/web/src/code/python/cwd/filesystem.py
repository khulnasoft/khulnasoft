from khulnasoft import Sandbox

sandbox = Sandbox(
    template="base",
    cwd="/home/user/code"  # $HighlightLine
)
sandbox.filesystem.write("hello.txt", "Welcome to KHULNASOFT!")  # $HighlightLine
proc = sandbox.process.start("cat /home/user/code/hello.txt")
proc.wait()
print(proc.output.stdout)
# output: "Welcome to KHULNASOFT!"

sandbox.filesystem.write("../hello.txt", "We hope you have a great day!")  # $HighlightLine
proc2 = sandbox.process.start("cat /home/user/hello.txt")
proc2.wait()
print(proc2.output.stdout)
# output: "We hope you have a great day!"

sandbox.close()
