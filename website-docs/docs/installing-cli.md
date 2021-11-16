---
layout: docs
styleoverrides: docs
title: Installing the Kelda CLI
weight: 500
---

Kelda works on macOS, Linux, and Windows Subsystem for Linux.

Run the following command to download and install Kelda.

```
curl -fsSL 'https://kelda.io/install.sh' | sh
```

You should see the following output:

```
Downloading the latest Kelda release...
######################################################################## 100.0%

The latest Kelda release has been downloaded to the current working directory.

Copy the binary into /usr/local/bin? (y/N) y
You may be prompted for your sudo password in order to write to /usr/local/bin.
Password:
```

## What's Next?

You're now ready to [try out an example application](/kelda-v1/docs/example-apps/polls).

If you'd like to run Kelda on your own infrastructure, you'll have to [install the Minion](/kelda-v1/docs/deployment) as well.
