## About Khulnasoft and the Go Gin example application

Khulnasoft is a developer-friendly serverless platform to deploy apps globally. No-ops, servers, or infrastructure management.
This repository contains a Go Gin application you can deploy on the Khulnasoft serverless platform for testing.

This example application is designed to show how a Go Gin application can be deployed on Khulnasoft.

## Getting Started

Follow the steps below to deploy and run the Go Gin application on your Khulnasoft account.

### Requirements

You need a Khulnasoft account to successfully deploy and run this application. If you don't already have an account, you can sign-up for free [here](https://app.khulnasoft.com/auth/signup).

### Deploy using the Khulnasoft button

The fastest way to deploy the Go Gin application is to click the **Deploy to Khulnasoft** button below.

Clicking on this button brings you to the Khulnasoft App creation page with everything pre-set to launch this application.

_To modify this application example, you will need to fork this repository. Checkout the [fork and deploy](#fork-and-deploy-to-khulnasoft) instructions._

### Fork and deploy to Khulnasoft

If you want to customize and enhance this application, you need to fork this repository.

If you used the **Deploy to Khulnasoft** button, you can simply link your service to your forked repository to be able to push changes.
Alternatively, you can manually create the application as described below.

On the [Khulnasoft Control Panel](https://app.khulnasoft.com/), on the **Overview** tab, click the **Create Web Service** button to begin.

1. Select **GitHub** as the deployment method.
2. In the repositories list, select the repository you just forked.
3. Expand the **Builder** section and click the **override** toggle associated with the **Run command**.  In the field, enter `bin/example-go-gin`.
4. Choose a name for your App, i.e `go-gin-on-khulnasoft`, and click **Deploy**.

You land on the deployment page where you can follow the build of your Go Gin application. Once the build is completed, your application is being deployed and you will be able to access it via `<YOUR_APP_NAME>-<YOUR_ORG_NAME>.khulnasoft.app`.

## Contributing

If you have any questions, ideas or suggestions regarding this application sample, feel free to open an [issue](//github.com/khulnasoft/khulnasoft/tree/main/examples/go-gin/issues) or fork this repository and open a [pull request](//github.com/khulnasoft/khulnasoft/tree/main/examples/go-gin/pulls).

## Contact

[Khulnasoft](https://www.khulnasoft.com) - [@gokhulnasoft](https://twitter.com/gokhulnasoft) - [Slack](http://slack.khulnasoft.com/)
