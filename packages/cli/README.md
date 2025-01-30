<p align="center">
  <img width="100" src="https://raw.githubusercontent.com/khulnasoft/KHULNASOFT/refs/heads/main/readme-assets/logo-circle.png" alt="khulnasoft logo">
</p>

# KhulnaSoft CLI

This CLI tool allows you to build manager your running KhulnaSoft sandbox and sandbox templates. Learn more in [our documentation](https://khulnasoft.com/docs).

### 1. Install the CLI

```bash
npm install -g @khulnasoft/cli
```

### 2. Authenticate

```bash
khulnasoft auth login
```

> [!NOTE] 
> To authenticate without the ability to open the browser, provide
> `KHULNASOFT_ACCESS_TOKEN` as an environment variable. Get your `KHULNASOFT_ACCESS_TOKEN`
> from the Personal tab at [khulnasoft.com/dashboard](https://khulnasoft.com/dashboard). Then use the CLI like this:
> `KHULNASOFT_ACCESS_TOKEN=sk_khulnasoft_... khulnasoft build`.

> [!IMPORTANT]  
> Note the distinction between `KHULNASOFT_ACCESS_TOKEN` and `KHULNASOFT_API_KEY`.

### 3. Check out docs
Visit our [CLI documentation](https://khulnasoft.com/docs) to learn more.
