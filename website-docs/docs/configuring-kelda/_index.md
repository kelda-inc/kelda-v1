---
layout: docs
styleoverrides: docs
title: Writing the Kelda Configuration
subtitle: Overview
weight: 800
---

This guide describes how to set up the Kelda configuration required to deploy your application.

For an overview of Kelda's configuration model, see the docs [here](/kelda-v1/docs/reference/configuration).

## Required Configuration

Two types of configuration are required in order to use Kelda:

1. The [**Workspace**](/kelda-v1/docs/configuring-kelda/workspace/) configuration, which contains the Kubernetes YAML for deploying all the services in the development environment.
2. The [**Sync**](/kelda-v1/docs/configuring-kelda/sync/) configuration, which defines how local code should be synced to the remote cluster during development.

Before using Kelda, you must set up the Sync configuration for at least one service, and the Workspace configuration for all services that it depends on.

## Approach

We recommend the following approach when configuring a new application.

1. Pick the first service that you want to develop with Kelda.
2. Follow the [guide on writing the Workspace configuration](/kelda-v1/docs/configuring-kelda/workspace/). This configuration contains the Kubernetes YAML for deploying the services into the development environment.
3. Follow the [guide on writing the Sync configuration](/kelda-v1/docs/configuring-kelda/sync/) to enable development on your first service.
4. At this point, you can start using Kelda to develop your first service. You can write the Sync configuration for other services as necessary.
