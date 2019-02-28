---
title: Overview
---

The Go SDK contains packages for interacting with most aspects of the hcnet ecosystem.  In addition to generally useful, low-level packages such as [`keypair`](https://godoc.org/github.com/hcnet/go/keypair) (used for creating hcnet-compliant public/secret key pairs), the Go SDK also contains code for the server applications and client tools written in go.

## Godoc reference

The most accurate and up-to-date reference information on the Go SDK is found within godoc.  The godoc.org service automatically updates the documentation for the Go SDK everytime github is updated.  The godoc for all of our packages can be found at (https://godoc.org/github.com/hcnet/go).

## Client Packages

The Go SDK contains packages for interacting with the various hcnet services:

- [`aurora`](https://godoc.org/github.com/hcnet/go/clients/aurora) provides client access to a aurora server, allowing you to load account information, stream payments, post transactions and more.
- [`hcnettoml`](https://godoc.org/github.com/hcnet/go/clients/hcnettoml) provides the ability to resolve Hcnet.toml files from the internet.  You can read about [Hcnet.toml concepts here](../../guides/concepts/hcnet-toml.md).
- [`federation`](https://godoc.org/github.com/hcnet/go/clients/federation) makes it easy to resolve a hcnet addresses (e.g. `scott*hcnet.org`) into a hcnet account ID suitable for use within a transaction.

