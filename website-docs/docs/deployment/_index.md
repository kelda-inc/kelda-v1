---
layout: docs
styleoverrides: docs
title: Deploying Kelda to a Private Cluster
subtitle: Instructions
weight: 700
---

**Estimated Time**: 15 minutes<br/>

The public demo cluster is great for example applications, but you'll need your own Kelda cluster to work on private code.

By the end of this guide, you will have the [Kelda Minion](/kelda-v1/docs/reference/architecture#the-kelda-minion) running in your own Kubernetes cluster.  The [Kelda CLI](/kelda-v1/docs/reference/architecture#the-kelda-cli) interacts with the Minion to manage development environments in your cluster.

{{< img-docs src="architecture.png" alt="Architecture diagram" >}}

## Instructions

1. **[Install the Kelda CLI](/kelda-v1/docs/installing-cli)**

  Run the following command to download and install Kelda.

  ```
  curl -fsSL 'https://kelda.io/install.sh' | sh
  ```

2. **Create the [Kubernetes cluster](/kelda-v1/docs/reference/architecture#the-kubernetes-cluster) that Kelda will run on**

  {{< tip heading="Don't want to create a Kubernetes cluster? Try out [Hosted Kelda](/request-hosted-kelda-access) and skip straight to [configuring Kelda](/kelda-v1/docs/configuring-kelda)." />}}

  We recommend a cluster **3 times the size of your development laptop**.

  Create the cluster using one of the options [here](/kelda-v1/docs/deployment/creating-cluster).

3. **Point your local kubeconfig at the cluster**

  If you successfully created your Kubernetes cluster, it should show up under `kubectl config get-contexts`.

  Switch to the proper context with `kubectl config use-context <context name>`.

  As a sanity check, make sure that `kubectl get nodes` doesn't error.

4. **Deploy the Kelda Minion**

  The Kelda CLI communicates with the Kelda minion, which in turn communicates with the Kubernetes API to manage your development namespace.

  Run the following command to deploy the Minion to the cluster you created in step 1.

  ```
  kelda setup-minion
  ```

  {{< note heading="If you've been provided a customer license for Kelda.." >}}
  Download the license, and run `kelda setup-minion --license <path to license>` instead.
  {{< /note >}}

5. **Confirm that everything installed successfully**

  If everything installed successfully, running

  ```
  kelda version
  ```

  Should show something like

  ```
  local version:  0.14.2
  minion version: 0.14.2
  ```

## What's Next?

Congratulations! You've now setup your own instance of Kelda that can be used to run your applications. Check out the [configuration documentation](/kelda-v1/docs/configuring-kelda) to learn how to start developing on your own application.
