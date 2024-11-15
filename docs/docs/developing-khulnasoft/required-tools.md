---
layout: default
title: Required tools
nav_order: 1
description: "Tools required for developing KhulnaSoft Fair Web Analytics."
permalink: /developing-khulnasoft/required-tools/
parent: For developers
---

<!--
Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-License-Identifier: Apache-2.0
-->

# Required tools

Developing KhulnaSoft Fair Web Analytics currently requires [Docker][docker] and [the compose plugin][docker-compose] to be installed and working correctly. Both tools are free and available for Linux, Windows and Mac.

If your system knows how to run `make` it will help a lot with common tasks like installing dependencies and building, yet you will still be able to work on KhulnaSoft Fair Web Analytics without it (this will likely be the case if you are developing on Windows).

If you are on an older version of Windows and want to use `make`, consider installing [Git Bash][git-bash], if you are using Windows 10, you could use the [Windows Subsystem for Linux][wsl]. Alternatively, you can look up the commands in the Makefile and resort to manually running them in sequence.

[docker]: https://docs.docker.com/install/
[docker-compose]: https://docs.docker.com/compose/install/
[git-bash]: https://gitforwindows.org/
[wsl]: https://docs.microsoft.com/en-us/windows/wsl/install-win10
