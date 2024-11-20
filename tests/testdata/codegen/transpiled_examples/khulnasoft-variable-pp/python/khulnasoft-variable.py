import khulnasoft
import os

khulnasoft.export("cwd", os.getcwd())
khulnasoft.export("stack", khulnasoft.get_stack())
khulnasoft.export("project", khulnasoft.get_project())
khulnasoft.export("organization", khulnasoft.get_organization())
