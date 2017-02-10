# kube-mesos-framework

![build](https://travis-ci.org/kubernetes-incubator/kube-mesos-framework.svg?branch=master)

kube-mesos-framework modifies Kubernetes to act as an [Apache Mesos](http://mesos.apache.org/) framework.

## Features On Mesos

Kubernetes gains the following benefits when installed on Mesos:

- **Node-Level Auto-Scaling** - Kubernetes minion nodes are created automatically, up to the size of the provisioned Mesos cluster.
- **Resource Sharing** - Co-location of Kubernetes with other popular next-generation services on the same cluster (e.g. [Hadoop](https://github.com/mesos/hadoop), [Spark](http://spark.apache.org/), and [Chronos](https://mesos.github.io/chronos/), [Cassandra](http://mesosphere.github.io/cassandra-mesos/), etc.). Resources are allocated to the frameworks based on fairness and can be claimed or passed on depending on framework load.

For more information about how kube-mesos-framework is different from Kubernetes, see [Architecture](./docs/architecture.md).

## Release Status

Kubernetes-Mesos is alpha quality, still under active development, and not yet recommended for production systems.

For more information about development progress, see the [known issues](./docs/issues.md) where backlog issues are tracked.

## Usage

This project combines concepts and technologies from two already-complex projects: Mesos and Kubernetes. It may help to familiarize yourself with the basics of each project before reading on:

* [Mesos Documentation](http://mesos.apache.org/documentation/latest)
* [Kubernetes Documentation](https://github.com/kubernetes/kubernetes)

To get up and running with Kubernetes-Mesos, follow:

- the [Getting started guide](./docs/getting-started-guides/mesos.md) to launch a Kubernetes-Mesos cluster,

## Conferences

* [MesosCon 2016 Asia](http://sched.co/8QFB): Kubernetes on Mesos: Not Just Another Mesos Framework - Klaus Ma, IBM
* [KubeCon 2016](http://sched.co/8K3n): Kubernetes on EGO -- Bringing Enterprise Resource Management and Scheduling to Kubernetes - Da Ma, IBM

## Organizations using kube-mesos-framework

* [IBM Spectrum Conductor for Container](https://www.ibm.com/developerworks/community/wikis/home?lang=en#!/wiki/W1559b1be149d_43b0_881e_9783f38faaff)
* [China Unicom Research Institute](http://61.148.212.23/home.jsp)

[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/contrib/mesos/README.md?pixel)]()
