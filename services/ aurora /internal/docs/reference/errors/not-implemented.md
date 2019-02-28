---
title: Not Implemented
---

If your [request method](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html) is not supported by Aurora, Aurora will return a `not_implemented` error. This is analogous to a [HTTP 501 Error](https://developer.mozilla.org/en-US/docs/Web/HTTP/Response_codes).

If you are encountering this error, Aurora does not have the functionality you are requesting yet.

## Attributes

As with all errors Aurora returns, `not_implemented` follows the [Problem Details for HTTP APIs](https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00) draft specification guide and thus has the following attributes:

| Attribute | Type   | Description                                                                                                                     |
| --------- | ----   | ------------------------------------------------------------------------------------------------------------------------------- |
| Type      | URL    | The identifier for the error.  This is a URL that can be visited in the browser.                                                |
| Title     | String | A short title describing the error.                                                                                             |
| Status    | Number | An HTTP status code that maps to the error.                                                                                     |
| Detail    | String | A more detailed description of the error.                                                                                       |
| Instance  | String | A token that uniquely identifies this request. Allows server administrators to correlate a client report with server log files. |


## Examples

```shell
$ curl -X GET "https://aurora-testnet.hcnet.org/ledgers/200/effects"
{
  "type": "not_implemented",
  "title": "Resource Not Yet Implemented",
  "status": 404,
  "detail": "While the requested URL is expected to eventually point to a valid resource, the work to implement the resource has not yet been completed.",
  "instance": "aurora-testnet-001.prd.hcnet001.internal.hcnet-ops.com/hCYL7oezXs-141917"
}
```

## Related

[Server Error](./server-error.md)
