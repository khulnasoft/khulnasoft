<a href="https://www.khulnasoft.com?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=top-logo" title="Pulumi - Modern Infrastructure as Code - AWS Azure Kubernetes Containers Serverless">
    <img src="https://www.khulnasoft.com/images/logo/logo-on-white-box.svg?" width="350">
</a>

[![Slack](http://www.khulnasoft.com/images/docs/badges/slack.svg)](https://slack.khulnasoft.com?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=slack-badge)
![GitHub Discussions](https://img.shields.io/github/discussions/khulnasoft/khulnasoft)
[![NPM version](https://badge.fury.io/js/%40khulnasoft%2Fkhulnasoft.svg)](https://npmjs.com/package/@khulnasoft/khulnasoft)
[![Python version](https://badge.fury.io/py/khulnasoft.svg)](https://pypi.org/project/khulnasoft)
[![NuGet version](https://badge.fury.io/nu/khulnasoft.svg)](https://badge.fury.io/nu/khulnasoft)
[![GoDoc](https://godoc.org/github.com/khulnasoft/khulnasoft?status.svg)](https://godoc.org/github.com/khulnasoft/khulnasoft)
[![License](https://img.shields.io/github/license/khulnasoft/khulnasoft)](LICENSE)
[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/khulnasoft/khulnasoft)

<a href="https://www.khulnasoft.com/docs/get-started/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=get-started-button" title="Get Started">
    <img src="https://www.khulnasoft.com/images/get-started.svg?" align="right" width="120">
</a>

**Pulumi's Infrastructure as Code SDK** is the easiest way to build and deploy infrastructure, of any architecture and on any cloud, using programming languages that you already know and love. Code and ship infrastructure faster with your favorite languages and tools, and embed IaC anywhere with [Automation API](https://www.khulnasoft.com/docs/guides/automation-api/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=automation+api).

Simply write code in your favorite language and Pulumi automatically provisions and manages your resources on
[AWS](https://www.khulnasoft.com/docs/reference/clouds/aws/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=aws-reference-link),
[Azure](https://www.khulnasoft.com/docs/reference/clouds/azure/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=azure-reference-link),
[Google Cloud Platform](https://www.khulnasoft.com/docs/reference/clouds/gcp/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=gcp-reference-link), 
[Kubernetes](https://www.khulnasoft.com/docs/reference/clouds/kubernetes/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=kuberneters-reference-link), and [120+ providers](https://www.khulnasoft.com/registry/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=providers-reference-link) using an
[infrastructure-as-code](https://www.khulnasoft.com/what-is/what-is-infrastructure-as-code/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=infrastructure-as-code) approach.
Skip the YAML, and use standard language features like loops, functions, classes,
and package management that you already know and love.

For example, create three web servers:

```typescript
const aws = require("@khulnasoft/aws");
const sg = new aws.ec2.SecurityGroup("web-sg", {
    ingress: [{ protocol: "tcp", fromPort: 80, toPort: 80, cidrBlocks: ["0.0.0.0/0"] }],
});
for (let i = 0; i < 3; i++) {
    new aws.ec2.Instance(`web-${i}`, {
        ami: "ami-7172b611",
        instanceType: "t2.micro",
        vpcSecurityGroupIds: [sg.id],
        userData: `#!/bin/bash
            echo "Hello, World!" > index.html
            nohup python -m SimpleHTTPServer 80 &`,
    });
}
```

Or a simple serverless timer that archives Hacker News every day at 8:30AM:

```typescript
const aws = require("@khulnasoft/aws");

const snapshots = new aws.dynamodb.Table("snapshots", {
    attributes: [{ name: "id", type: "S", }],
    hashKey: "id", billingMode: "PAY_PER_REQUEST",
});

aws.cloudwatch.onSchedule("daily-yc-snapshot", "cron(30 8 * * ? *)", () => {
    require("https").get("https://news.ycombinator.com", res => {
        let content = "";
        res.setEncoding("utf8");
        res.on("data", chunk => content += chunk);
        res.on("end", () => new aws.sdk.DynamoDB.DocumentClient().put({
            TableName: snapshots.name.get(),
            Item: { date: Date.now(), content },
        }).promise());
    }).end();
});
```

Many examples are available spanning containers, serverless, and infrastructure in
[khulnasoft/examples](https://github.com/khulnasoft/examples).

Pulumi is open source under the [Apache 2.0 license](https://github.com/khulnasoft/khulnasoft/blob/master/LICENSE), supports many languages and clouds, and is easy to extend.  This
repo contains the `khulnasoft` CLI, language SDKs, and core Pulumi engine, and individual libraries are in their own repos.

## Welcome

<img align="right" width="400" src="https://www.khulnasoft.com/images/docs/quickstart/console.png" />

* **[Get Started with Pulumi](https://www.khulnasoft.com/docs/get-started/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=get+started+with+khulnasoft)**: Deploy a simple application in AWS, Azure, Google Cloud, or Kubernetes using Pulumi.

* **[Learn](https://www.khulnasoft.com/learn/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=learn)**: Follow Pulumi learning pathways to learn best practices and architectural patterns through authentic examples.

* **[Examples](https://github.com/khulnasoft/examples)**: Browse several examples across many languages,
  clouds, and scenarios including containers, serverless, and infrastructure.

* **[Docs](https://www.khulnasoft.com/docs/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=docs)**: Learn about Pulumi concepts, follow user-guides, and consult the reference documentation.

* **[Registry](https://www.khulnasoft.com/registry/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=registry)**: Find the Pulumi Package with the resources you need. Install the package directly into your project, browse the API documentation, and start building.

* **[Secrets Management](https://www.khulnasoft.com/esc/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=secrets+management)**: Tame secrets sprawl and configuration complexity securely across all your cloud infrastructure and applications with Pulumi ESC.

* **[Pulumi Roadmap](https://github.com/orgs/khulnasoft/projects/44)**: Review the planned work for the upcoming quarter and a selected backlog of issues that are on our mind but not yet scheduled.

* **[Community Slack](https://slack.khulnasoft.com/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=welcome-slack)**: Join us in Pulumi Community Slack. All conversations and questions are welcome.

* **[GitHub Discussions](https://github.com/khulnasoft/khulnasoft/discussions)**: Ask questions or share what you're building with Pulumi.

## <a name="getting-started"></a>Getting Started

[![Watch the video](/youtube_preview_image.png)](https://www.youtube.com/watch?v=6f8KF6UGN7g)

See the [Get Started](https://www.khulnasoft.com/docs/quickstart/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=getting-started-quickstart) guide to quickly get started with
Pulumi on your platform and cloud of choice.

Otherwise, the following steps demonstrate how to deploy your first Pulumi program, using AWS
Serverless Lambdas, in minutes:

1. **Install**:

    To install the latest Pulumi release, run the following (see full
    [installation instructions](https://www.khulnasoft.com/docs/reference/install/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=getting-started-install) for additional installation options):

    ```bash
    $ curl -fsSL https://get.khulnasoft.com/ | sh
    ```

2. **Create a Project**:

    After installing, you can get started with the `khulnasoft new` command:

    ```bash
    $ mkdir khulnasoft-demo && cd khulnasoft-demo
    $ khulnasoft new hello-aws-javascript
    ```

    The `new` command offers templates for all languages and clouds.  Run it without an argument and it'll prompt
    you with available projects.  This command created an AWS Serverless Lambda project written in JavaScript.

3. **Deploy to the Cloud**:

    Run `khulnasoft up` to get your code to the cloud:

    ```bash
    $ khulnasoft up
    ```

    This makes all cloud resources needed to run your code.  Simply make edits to your project, and subsequent
    `khulnasoft up`s will compute the minimal diff to deploy your changes.

4. **Use Your Program**:

    Now that your code is deployed, you can interact with it.  In the above example, we can curl the endpoint:

    ```bash
    $ curl $(khulnasoft stack output url)
    ```

5. **Access the Logs**:

    If you're using containers or functions, Pulumi's unified logging command will show all of your logs:

    ```bash
    $ khulnasoft logs -f
    ```

6. **Destroy your Resources**:

    After you're done, you can remove all resources created by your program:

    ```bash
    $ khulnasoft destroy -y
    ```

To learn more, head over to [khulnasoft.com](https://khulnasoft.com/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=getting-started-learn-more-home) for much more information, including
[tutorials](https://www.khulnasoft.com/docs/reference/tutorials/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=getting-started-learn-more-tutorials), [examples](https://github.com/khulnasoft/examples), and
details of the core Pulumi CLI and [programming model concepts](https://www.khulnasoft.com/docs/reference/concepts/?utm_campaign=khulnasoft-khulnasoft-github-repo&utm_source=github.com&utm_medium=getting-started-learn-more-concepts).

## <a name="platform"></a>Platform

### Languages

|    | Language | Status | Runtime | Versions |
| -- | -------- | ------ | ------- | -------- |
| <img src="https://www.khulnasoft.com/logos/tech/logo-js.png" height=38 />     | [JavaScript](https://www.khulnasoft.com/docs/intro/languages/javascript/) | Stable  | Node.js | [Current, Active and Maintenance LTS versions](https://nodejs.org/en/about/previous-releases)  |
| <img src="https://www.khulnasoft.com/logos/tech/logo-ts.png" height=38 />     | [TypeScript](https://www.khulnasoft.com/docs/intro/languages/javascript/) | Stable  | Node.js | [Current, Active and Maintenance LTS versions](https://nodejs.org/en/about/previous-releases)  |
| <img src="https://www.khulnasoft.com/logos/tech/logo-python.svg" height=38 /> | [Python](https://www.khulnasoft.com/docs/intro/languages/python/)     | Stable  | Python | [Supported versions](https://devguide.python.org/versions/#versions) |
| <img src="https://www.khulnasoft.com/logos/tech/logo-golang.png" height=38 /> | [Go](https://www.khulnasoft.com/docs/intro/languages/go/)             | Stable  | Go | [Supported versions](https://go.dev/doc/devel/release#policy) |
| <img src="https://www.khulnasoft.com/logos/tech/dotnet.svg" height=38 />      | [.NET (C#/F#/VB.NET)](https://www.khulnasoft.com/docs/intro/languages/dotnet/)     | Stable  | .NET | [Supported versions](https://dotnet.microsoft.com/en-us/platform/support/policy/dotnet-core#lifecycle)  |
| <img src="https://www.khulnasoft.com/logos/tech/java.svg" height=38 />      | [Java](https://www.khulnasoft.com/docs/intro/languages/java/)     | Public Preview  | JDK | 11+  |
| <img src="https://www.khulnasoft.com/logos/tech/yaml.svg" height=38 />      | [YAML](https://www.khulnasoft.com/docs/intro/languages/yaml/)     | Stable  | n/a  | n/a  |

### EOL Releases

The Pulumi CLI v1 and v2 are no longer supported. If you are not yet running v3, please consider migrating to v3 to continue getting the latest and greatest Pulumi has to offer! :muscle:

* To migrate from v2 to v3, please see our [v3 Migration Guide](https://www.khulnasoft.com/docs/install/migrating-3.0/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=v3+migration+guide).

### Clouds

Visit the [Registry](https://www.khulnasoft.com/registry/?utm_source=github.com&utm_medium=referral&utm_campaign=khulnasoft-khulnasoft-github-repo&utm_content=registry) for the full list of supported cloud and infrastructure providers.

## Contributing

Visit [CONTRIBUTING.md](https://github.com/khulnasoft/khulnasoft/blob/master/CONTRIBUTING.md) for information on building Pulumi from source or contributing improvements.
