---
layout: default
title: Installation on Ubuntu
nav_order: 2
description: "A step by step tutorial on how to deploy KhulnaSoft Fair Web Analytics on an Ubuntu system using systemd."
permalink: /running-khulnasoft/tutorials/configuring-deploying-khulnasoft-ubuntu/
parent: For operators
---

<!--
Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-License-Identifier: Apache-2.0
-->

# Installation on Ubuntu

Configuring and deploying KhulnaSoft Fair Web Analytics on Ubuntu
{: .no_toc }

This tutorial walks you through the steps needed to setup and deploy a standalone, single-node KhulnaSoft Fair Web Analytics instance that is using a local SQLite file as its database backend. `systemd` is used for managing the KhulnaSoft Fair Web Analytics service.

<span class="label label-green">Note</span>

If you get stuck or need help, [file an issue][gh-issues], or send an [email][email]. If you have installed KhulnaSoft Fair Web Analytics and would like to spread the word, we're happy to feature you in our README. [Send a PR][edit-readme] adding your site or app and we'll merge it.

[gh-issues]: https://github.com/khulnasoft/khulnasoft/issues
[mastodon]: https://fosstodon.org/@khulnasoft
[email]: mailto:admin@khulnasoft.com
[edit-readme]: https://github.com/khulnasoft/khulnasoft/edit/development/README.md

---

## Table of contents
{: .no_toc }

1. TOC
{:toc}

---

## Prerequisites

This tutorial assumes the machine you are planning to run KhulnaSoft Fair Web Analytics on is connected to the internet and has DNS records for `khulnasoft.mysite.com` (or the domain you are actually planning to use) pointing to it. Ports 80 and 443 are expected to be accessible to the public. See the [documentation for subdomains][domain-doc] for further information on this topic.

[domain-doc]: ./../setting-up-using-subdomains/

## Downloading and installing the package

KhulnaSoft Fair Web Analytics version v0.1.6 and later is packaged as a Debian package, so installation on Ubuntu (and other Debian based distributions) is easy. First, download the package for the latest release:

```
curl -sSL https://get.khulnasoft.com/deb -o khulnasoft.deb
```

Next, you can verify the package's signature using `gpg` and `dpkg-sig` (this step is optional, but recommended):
```
curl https://keybase.io/khulnasoft/pgp_keys.asc | gpg --import
dpkg-sig --verify khulnasoft.deb
```

The package itself can be installed using `dpkg`:

```
sudo dpkg -i khulnasoft.deb
```

You can confirm that your installation is working as expected like this:
```
$ which khulnasoft
/usr/local/bin/khulnasoft
$ khulnasoft version
INFO[0000] Current build created using                   revision={{ site.khulnasoft_version }}
```

You can now safely remove the download:

```
rm khulnasoft.deb
```

[releases]: https://github.com/khulnasoft/khulnasoft/releases

---

## Configuring KhulnaSoft Fair Web Analytics

In this setup, KhulnaSoft Fair Web Analytics stores its runtime configuration in `/etc/khulnasoft/khulnasoft.env`. This file has already been created on installation, so you can now populate it with the values required for your install.

### Application Secret

KhulnaSoft Fair Web Analytics is using a secret to sign login cookies and tokens for resetting passwords or inviting users. You can generate a unique secret for your installation using the `khulnasoft secret` subcommand:

```
$ khulnasoft secret
INFO[0000] Created 16 bytes secret                       secret="S2dR9JYYTNG3+5QN+jxiwA=="
```

Populate the `KHULNASOFT_SECRET` key with the value you just generated:

```
KHULNASOFT_SECRET="S2dR9JYYTNG3+5QN+jxiwA==" # do not use this secret in production
```

__Heads Up__
{: .label .label-red }

If you do not set this config value, KhulnaSoft Fair Web Analytics will generate a random one every time it starts up. This means it works securely, yet all login sessions, password reset emails or invitations will be invalidated when the service restarts.

### Setting up AutoTLS

KhulnaSoft Fair Web Analytics requires a secure connection and can automatically acquire a renew SSL certificates from LetsEncrypt for your domain. Add the domain you want to use to serve KhulnaSoft Fair Web Analytics to `KHULNASOFT_SERVER_AUTOTLS`:

```
KHULNASOFT_SERVER_AUTOTLS="khulnasoft.mysite.com"
```

To make sure the automatic certificate creation and renewal works, make sure your host system exposes __both port 80 and 443__ to the public internet.

### Setting up email

KhulnaSoft Fair Web Analytics needs to send transactional email for the following features:

- Inviting a new user to an account
- Resetting your password in case you forgot it

To enable this, you can add SMTP credentials, namely __Host, Sender, User, Password and Port__ to the `khulnasoft.env` file:

```
KHULNASOFT_SMTP_HOST="smtp.mysite.com"
KHULNASOFT_SMTP_SENDER="khulnasoft@mysite.com"
KHULNASOFT_SMTP_USER="me"
KHULNASOFT_SMTP_PASSWORD="my-password"
KHULNASOFT_SMTP_PORT="587"
```

__Heads Up__
{: .label .label-red }

KhulnaSoft Fair Web Analytics will run without these values being set and try to fall back to a local `sendmail` install, yet please be aware that if you rely on any of the above features email delivery will be __very unreliable if not configured correctly__. You can always add this at a later time though.

---

### Verifying your config file

Before you start the application, it's a good idea to double check the setup. Your config file at `/etc/khulnasoft/khulnasoft.env` should now contain an entry for each of these values:

```
KHULNASOFT_SECRET="uNrZP7r5fY3sfS35tbzR9w==" # do not use this secret in production
KHULNASOFT_SERVER_AUTOTLS="khulnasoft.mysite.com"
KHULNASOFT_SMTP_HOST="smtp.mysite.com"
KHULNASOFT_SMTP_USER="me"
KHULNASOFT_SMTP_PASSWORD="my-password"
KHULNASOFT_SMTP_PORT="587"
```

If all of this is populated with the values you expect, you're ready to use KhulnaSoft Fair Web Analytics.

---

## Starting the `systemd` service

`systemd` is used to make sure KhulnaSoft Fair Web Analytics is up and running at all times (e.g. after rebooting or crashing) and accepts events. The `deb` package has already creating a `systemd` service for you on installation, so all you need to do now is start it:

```
sudo systemctl enable khulnasoft
sudo systemctl start khulnasoft
```

You can check whether this worked correctly using `status:`

```
$ sudo systemctl status khulnasoft
● khulnasoft.service - KhulnaSoft Service
   Loaded: loaded (/etc/systemd/system/khulnasoft.service; enabled; vendor preset: enabled)
   Active: active (running) since Mon 2020-01-27 15:57:58 CET; 1min ago
 Main PID: 6701 (khulnasoft)
    Tasks: 11 (limit: 4915)
   CGroup: /system.slice/khulnasoft.service
           └─6701 /usr/local/bin/khulnasoft
```


Your instance is now ready to use.

---

## Setting up the instance

Now that KhulnaSoft Fair Web Analytics is up and running, you can create your login user and a first account by navigating to `https://khulnasoft.mysite.com/setup`. You can create one user and one account here, but you can always add more later on.

After submitting the form, your KhulnaSoft Fair Web Analytics instance is ready to use.

## Maintenance

### Accessing logs

The easiest way for accessing application logs in this setup is using `journald`

```
$ sudo journalctl -u khulnasoft
khulnasoft[6573]: time="2020-01-27T15:57:41+01:00" level=info msg="Successfully applied database migrations"
khulnasoft[6573]: time="2020-01-27T15:57:41+01:00" level=info msg="Server now listening on port 80 and 443 using AutoTLS"
khulnasoft[6573]: time="2020-01-27T15:57:41+01:00" level=info msg="Cron successfully pruned expired events" removed=0
```

### Uninstalling the service

If you want to uninstall the service from your system, stop and disable the `khulnasoft` service:

```
sudo systemctl stop khulnasoft
sudo systemctl disable khulnasoft
```

### Updating the version in use

To update to a new version of KhulnaSoft Fair Web Analytics, download the package for the newer version and install:

```
curl https://get.khulnasoft.com/deb -o khulnasoft.deb
dpkg-sig --verify khulnasoft.deb
sudo dpkg -i khulnasoft.deb
```

Confirm that this worked by having `khulnasoft` print its updated version:

```
$ khulnasoft version
INFO[0000] Current build created using                   revision=v0.2.12
```

You can now restart your service to pick up the changes:

```
sudo systemctl restart khulnasoft
```
