# Flux version 2

[![e2e](https://github.com/fluxcd/flux2/workflows/e2e/badge.svg)](https://github.com/fluxcd/flux2/actions)
[![report](https://goreportcard.com/badge/github.com/fluxcd/flux2)](https://goreportcard.com/report/github.com/fluxcd/flux2)
[![license](https://img.shields.io/github/license/fluxcd/flux2.svg)](https://github.com/fluxcd/flux2/blob/main/LICENSE)
[![release](https://img.shields.io/github/release/fluxcd/flux2/all.svg)](https://github.com/fluxcd/flux2/releases)

Flux is a tool for keeping Kubernetes clusters in sync with sources of
configuration (like Git repositories), and automating updates to
configuration when there is new code to deploy.

Flux version 2 ("v2") is built from the ground up to use Kubernetes'
API extension system, and to integrate with Prometheus and other core
components of the Kubernetes ecosystem. In version 2, Flux supports
multi-tenancy and support for syncing an arbitrary number of Git
repositories, among other long-requested features.

Flux v2 is constructed with the [GitOps Toolkit](#gitops-toolkit), a
set of composable APIs and specialized tools for building Continuous
Delivery on top of Kubernetes.

## Flux installation

With Homebrew:

```sh
brew install fluxcd/tap/flux
```

With Bash:

```sh
curl -s https://toolkit.fluxcd.io/install.sh | sudo bash

# enable completions in ~/.bash_profile
. <(flux completion bash)
```

Arch Linux (AUR) packages:

- [flux-bin](https://aur.archlinux.org/packages/flux-bin): install the latest
  stable version using a pre-build binary (recommended)
- [flux-go](https://aur.archlinux.org/packages/flux-go): build the latest
  stable version from source code
- [flux-scm](https://aur.archlinux.org/packages/flux-scm): build the latest
  (unstable) version from source code from our git `main` branch

Binaries for macOS, Windows and Linux AMD64/ARM are available to download on the
[release page](https://github.com/fluxcd/flux2/releases).

Verify that your cluster satisfies the prerequisites with:

```sh
flux check --pre
```

## Get started

To get started with Flux, start [browsing the
documentation](https://toolkit.fluxcd.io) or get started with one of
the following guides:

- [Get started with Flux](https://toolkit.fluxcd.io/get-started/)
- [Manage Helm Releases](https://toolkit.fluxcd.io/guides/helmreleases/)
- [Automate image updates to Git](https://toolkit.fluxcd.io/guides/image-update/)  
- [Manage Kubernetes secrets with Mozilla SOPS](https://toolkit.fluxcd.io/guides/mozilla-sops/)  

If you need help, please refer to our **[Support page](https://fluxcd.io/support/)**.

## GitOps Toolkit

The GitOps Toolkit is the set of APIs and controllers that make up the
runtime for Flux v2. The APIs comprise Kubernetes custom resources,
which can be created and updated by a cluster user, or by other
automation tooling.

![overview](docs/_files/gitops-toolkit.png)

You can use the toolkit to extend Flux, or to build your own systems
for continuous delivery -- see [the developer
guides](https://toolkit.fluxcd.io/dev-guides/source-watcher/).

### Components

- [Source Controller](https://toolkit.fluxcd.io/components/source/controller/)
    - [GitRepository CRD](https://toolkit.fluxcd.io/components/source/gitrepositories/)
    - [HelmRepository CRD](https://toolkit.fluxcd.io/components/source/helmrepositories/)
    - [HelmChart CRD](https://toolkit.fluxcd.io/components/source/helmcharts/)
    - [Bucket CRD](https://toolkit.fluxcd.io/components/source/buckets/)
- [Kustomize Controller](https://toolkit.fluxcd.io/components/kustomize/controller/)
    - [Kustomization CRD](https://toolkit.fluxcd.io/components/kustomize/kustomization/)
- [Helm Controller](https://toolkit.fluxcd.io/components/helm/controller/)
    - [HelmRelease CRD](https://toolkit.fluxcd.io/components/helm/helmreleases/)
- [Notification Controller](https://toolkit.fluxcd.io/components/notification/controller/)
    - [Provider CRD](https://toolkit.fluxcd.io/components/notification/provider/)
    - [Alert CRD](https://toolkit.fluxcd.io/components/notification/alert/)
    - [Receiver CRD](https://toolkit.fluxcd.io/components/notification/receiver/)
- [Image Automation Controllers](https://toolkit.fluxcd.io/components/image/controller/)
  - [ImageRepository CRD](https://toolkit.fluxcd.io/components/image/imagerepositories/)
  - [ImagePolicy CRD](https://toolkit.fluxcd.io/components/image/imagepolicies/)
  - [ImageUpdateAutomation CRD](https://toolkit.fluxcd.io/components/image/imageupdateautomations/)
  
## Community

Need help or want to contribute? Please see the links below. The Flux project is always looking for
new contributors and there are a multitude of ways to get involved.

- Getting Started?
    - Look at our [Get Started guide](https://toolkit.fluxcd.io/get-started/) and give us feedback
- Need help?
    - First: Ask questions on our [GH Discussions page](https://github.com/fluxcd/flux2/discussions)
    - Second: Talk to us in the #flux channel on [CNCF Slack](https://slack.cncf.io/)
    - Please follow our [Support Guidelines](https://fluxcd.io/support/)
      (in short: be nice, be respectful of volunteers' time, understand that maintainers and
      contributors cannot respond to all DMs, and keep discussions in the public #flux channel as much as possible).
- Have feature proposals or want to contribute?
    - Propose features on our [GH Discussions page](https://github.com/fluxcd/flux2/discussions)
    - Join our upcoming dev meetings ([meeting access and agenda](https://docs.google.com/document/d/1l_M0om0qUEN_NNiGgpqJ2tvsF2iioHkaARDeh6b70B0/view))
    - [Join the flux-dev mailing list](https://lists.cncf.io/g/cncf-flux-dev).
    - Check out [how to contribute](CONTRIBUTING.md) to the project

### Events

Check out our **[events calendar](https://fluxcd.io/community/#talks)**,
both with upcoming talks you can attend or past events videos you can watch.

We look forward to seeing you with us!
