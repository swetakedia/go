---
title: Overview
---

Aurora is an API server for the Hcnet ecosystem.  It acts as the interface between [hcnet-core](https://github.com/hcnet/hcnet-core) and applications that want to access the Hcnet network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Hcnet ecosystem](https://www.hcnet.org/developers/guides/) for details of where Aurora fits in. You can also watch a [talk on Aurora](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Hcnet.org developer Scott Fleckenstein:

[![Aurora: API webserver for the Hcnet network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Aurora: API webserver for the Hcnet network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Aurora provides a RESTful API to allow client applications to interact with the Hcnet network. You can communicate with Aurora using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Hcnet SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.hcnet.org/developers/js-hcnet-sdk/learn/index.html) for clients to use to interact with Aurora.

SDF runs a instance of Aurora that is connected to the test net: [https://aurora-testnet.hcnet.org/](https://aurora-testnet.hcnet.org/) and one that is connected to the public Hcnet network:
[https://aurora.hcnet.org/](https://aurora.hcnet.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/hcnet/js-hcnet-sdk)
- [Java](https://github.com/hcnet/java-hcnet-sdk)
- [Go](https://github.com/hcnet/go)

Community maintained libraries (in various states of completeness) for interacting with Aurora in other languages:<br>
- [Ruby](https://github.com/hcnet/ruby-hcnet-sdk)
- [Python](https://github.com/HcnetCN/py-hcnet-base)
- [C#](https://github.com/elucidsoft/dotnet-hcnet-sdk)
