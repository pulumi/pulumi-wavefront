[![Actions Status](https://github.com/pulumi/pulumi-wavefront/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-wavefront/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fwavefront.svg)](https://www.npmjs.com/package/@pulumi/wavefront)
[![Python version](https://badge.fury.io/py/pulumi-wavefront.svg)](https://pypi.org/project/pulumi-wavefront)
[![NuGet version](https://badge.fury.io/nu/pulumi.wavefront.svg)](https://badge.fury.io/nu/pulumi.wavefront)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-wavefront/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-wavefront/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-wavefront/blob/master/LICENSE)

# Wavefront Resource Provider

The Wavefront Resource Provider lets you manage Wavefront resources.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/wavefront

or `yarn`:

    $ yarn add @pulumi/wavefront

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_wavefront

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-wavefront/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Wavefront

## Configuration

The following configuration points are available:

- `wavefront:token` - (Required) This is a either a Users token or Service Account token with permissions necessary to
  manage your Wavefront account. It can also be sourced from the `WAVEFRONT_TOKEN` environment variable.
- `wavefront:address` - (Required) This is the URL of your cluster that you access Wavefront from without the leading `https://`
  or trailing `/` (e.g. `https://longboard.wavefront.com/` becomes `longboard.wavefront.com`). It can also be sourced
  from the `WAVEFRONT_ADDRESS` environment variable.

## Reference

For further information, please visit [the Wavefront provider docs](https://www.pulumi.com/docs/intro/cloud-providers/wavefront)
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/wavefront).
