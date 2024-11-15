---
layout: default
title: Downloads and distributions
nav_order: 8
description: "Where and how to download KhulnaSoft Fair Web Analytics or install it from other sources."
permalink: /running-khulnasoft/downloads-distributions/
parent: For operators
---

<!--
Copyright 2020-2022 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-License-Identifier: Apache-2.0
-->

# Downloads and distributions
{: .no_toc }

KhulnaSoft Fair Web Analytics is currently distributed in two ways: packaged as a __single binary file__ for Linux, MacOS and Windows or as a __Docker image__. The most recent release [is available here][most-recent], Docker images are hosted on [Docker Hub][docker-hub].

The latest binary tarball can be downloaded from `https://get.khulnasoft.com`:

```
curl -sSL https://get.khulnasoft.com -o khulnasoft.tar.gz
```

On Docker Hub, `latest` will give you the last release:

```
docker pull khulnasoft/khulnasoft:latest
```

For Ubuntu and Debian, we also provide a `deb` package you can download:

```
curl -sSL https://get.khulnasoft.com/deb -o khulnasoft.deb
```

---

## Table of contents
{: .no_toc }

1. TOC
{:toc}

---

## Release channels

Both for binaries and Docker images you can use one of these channels to pick your download:

### Tagged releases
{: .no_toc }

In case you want to deploy KhulnaSoft Fair Web Analytics, this channel is what you are most likely going to use. When ready, we cut official releases and tag them with a version identifier (e.g. `v0.1.0`). These releases are immutable and will never change, so both a download and the Docker image are guaranteed to provide the exact same build every time.

### Development channel
{: .no_toc }

The `development` channel gives you the most recent build from the `development` branch of our repository. This is likely to contain things that are not production ready or bring other kinds of caveats. Only binaries and Docker images are available for this channel. __You probably should not use this__ unless you are participating in the development of KhulnaSoft Fair Web Analytics.

---

## Downloading binary files

Binary files can be downloaded from our [GitHub repository][repo-releases] or using `get.khulnasoft.com`. The most recent release [is available here][most-recent].

Downloading `https://get.khulnasoft.com` will give you a tarball containing the most recent tagged release. If you specify a version or channel like `https://get.khulnasoft.com/v0.1.0-alpha.8` you will download that specific version.

```sh
# most recent release
curl -sSL https://get.khulnasoft.com
# most recent build from the development channel
curl -sSL https://get.khulnasoft.com/development
# build for {{ site.khulnasoft_version }}
curl -sSL https://get.khulnasoft.com/{{ site.khulnasoft_version }}
```

---

__Heads Up__
{: .label .label-red }

The archive file currently contains the binaries __for all supported operating systems__, so no matter which OS you are targeting, you will always download the same file.

[repo-releases]: https://github.com/khulnasoft/khulnasoft/releases
[most-recent]: https://get.khulnasoft.com

## Verifying the binaries' signatures

To prevent unwanted modifications of our releases, we sign all binaries using GPG and include the signature in our distribution. You can run the following commands to verify the integrity of your download:

```
curl https://keybase.io/khulnasoft/pgp_keys.asc | gpg --import
gpg --verify khulnasoft-linux-amd64.asc khulnasoft-linux-amd64
```

Debian packages are also signed:

```
curl https://keybase.io/khulnasoft/pgp_keys.asc | gpg --import
dpkg-sig --verify khulnasoft.deb
```

## Pulling the Docker image

Docker images are available as `khulnasoft/khulnasoft` on [Docker Hub][docker-hub]. Tagged releases are available under the respective tag (e.g. `khulnasoft/khulnasoft:v0.1.0`). The `development` channel is available as an image tag as well. As per Docker convention `latest` maps to the latest tagged release.

```sh
# {{ site.khulnasoft_version }} release
docker pull khulnasoft/khulnasoft:{{ site.khulnasoft_version }}
# development channel
docker pull khulnasoft/khulnasoft:development
```

---

__Heads Up__
{: .label .label-red }

While our version tags on Docker Hub are immutable and __will always return the same build__, it is important to note that `development` and `latest` will be updated on a rolling basis. If you deploy KhulnaSoft Fair Web Analytics using Docker, make sure to use a version tag or pin your image's revision.

[docker-hub]: https://hub.docker.com/r/khulnasoft/khulnasoft

---

## Click-To-Deploy

We currently also offer "Click-To-Deploy" packages for the following hosting providers:

### Heroku

You can deploy KhulnaSoft Fair Web Analytics to Heroku using our [deployment template][heroku-repo]. We also offer a [tutorial on this website][heroku-tutorial].

[heroku-repo]: https://github.com/khulnasoft/heroku
[heroku-tutorial]: /running-khulnasoft/tutorials/configuring-deploying-khulnasoft-heroku/

### YunoHost

If you use YunoHost to self-host your applications, you can install our [packaged KhulnaSoft app][ynh-repo] to get KhulnaSoft Fair Web Analytics up and running:

```
$ sudo yunohost app install https://github.com/khulnasoft/khulnasoft_ynh
```

[ynh-repo]: https://github.com/khulnasoft/khulnasoft_ynh

## Building KhulnaSoft for other architectures other

Current distributions of KhulnaSoft Fair Web Analytics target `amd64`, `arm64` and `arm/v7` architectures only. If you want to run it on different hardware, it is possible to build a binary for your target OS and platform yourself. Assuming you have Docker and `make` installed, you can build the latest version like this (the example targets Linux on ARM v6):

```
git clone git@github.com:khulnasoft/khulnasoft.git
cd khulnasoft
git checkout {{ site.khulnasoft_version }}
TARGETS=linux/arm-6 make build
```

Once finished, the `bin` folder contains your binary of choice. Currently supported architectures are all architectures supported by [xgo][xgo]:

- 386
- amd64
- arm-5
- arm-6
- arm-7
- arm64
- mips
- mipsle
- mips64
- mips64le

Supported operating systems are:

- linux
- darwin
- windows

[xgo]: https://github.com/techknowlogick/xgo
