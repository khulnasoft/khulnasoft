<h4 align="center">
  <a href="https://pypi.org/project/khulnasoft/">
    <img alt="Last 1 month downloads for the Python SDK" loading="lazy" width="200" height="20" decoding="async" data-nimg="1"
    style="color:transparent;width:auto;height:100%" src="https://img.shields.io/pypi/dm/khulnasoft?label=PyPI%20Downloads">
  </a>
  <a href="https://www.npmjs.com/package/khulnasoft">
    <img alt="Last 1 month downloads for the JavaScript SDK" loading="lazy" width="200" height="20" decoding="async" data-nimg="1"
    style="color:transparent;width:auto;height:100%" src="https://img.shields.io/npm/dm/khulnasoft?label=NPM%20Downloads">
  </a>
</h4>

<!---
<img width="100%" src="/readme-assets/preview.png" alt="Cover image">
--->
## What is KHULNASOFT?
[KHULNASOFT](https://www.khulnasoft.com/) is an open-source infrastructure that allows you to run AI-generated code in secure isolated sandboxes in the cloud. To start and control sandboxes, use our [JavaScript SDK](https://www.npmjs.com/package/@khulnasoft/code-interpreter) or [Python SDK](https://pypi.org/project/khulnasoft_code_interpreter).

> [!NOTE]
> This repository contains the core KhulnaSoft SDK that's used in our main [KhulnaSoft Code Interpreter SDK](https://github.com/khulnasoft/code-interpreter).

## Run your first Sandbox

### 1. Install SDK

JavaScript / TypeScript
```
npm i @khulnasoft/code-interpreter
```

Python
```
pip install khulnasoft-code-interpreter
```

### 2. Get your KHULNASOFT API key
1. Sign up to KhulnaSoft [here](https://khulnasoft.com).
2. Get your API key [here](https://khulnasoft.com/dashboard?tab=keys).
3. Set environment variable with your API key
```
KHULNASOFT_API_KEY=khulnasoft_***
```     

### 3. Execute code with code interpreter inside Sandbox

JavaScript / TypeScript
```ts
import { Sandbox } from '@khulnasoft/code-interpreter'

const sandbox = await Sandbox.create()
await sandbox.runCode('x = 1')

const execution = await sandbox.runCode('x+=1; x')
console.log(execution.text)  // outputs 2
```

Python
```py
from khulnasoft_code_interpreter import Sandbox

with Sandbox() as sandbox:
    sandbox.run_code("x = 1")
    execution = sandbox.run_code("x+=1; x")
    print(execution.text)  # outputs 2
```

### 4. Check docs
Visit [KhulnaSoft documentation](https://khulnasoft.com/docs).

### 5. KhulnaSoft cookbook
Visit our [Cookbook](https://github.com/khulnasoft/khulnasoft-cookbook/tree/main) to get inspired by examples with different LLMs and AI frameworks.
