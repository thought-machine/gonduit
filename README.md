# Gonduit [![GoDoc](https://godoc.org/github.com/etcinit/gonduit?status.svg)](https://godoc.org/github.com/etcinit/gonduit)

A Go client for interacting with [Phabricator](http://phabricator.org) via the [Conduit](https://secure.phabricator.com/book/phabdev/article/conduit/) API.

> Currently this library is incomplete. This is a fork from
> https://github.com/jpoehls/go-conduit with support for some additional call
> and options. The library is a bit messy right now.

# Getting started

## Installing the library

A simple `go get` should do it:

```
go get github.com/etcinit/gonduit
```

## Getting a conduit certificate

This library uses the `conduit.connect` method to establish an authenticated
session with the Phabricator instance.

This means that you'll need to have a valid username and and a Conduit
certificate in order to use the client.

To get a conduit certificate, go to
`https://{PHABRICATOR_URL}/settings/panel/conduit`. From there, you should be
able to copy your certificate. If you are creating a bot/automated script, you
should create a bot account on Phabricator rather than using your own.

# Basic Usage

## Connecting

Connecting to a Conduit API is a two-step process: First, `Dial` connects to
the API and checks compatibility, and finally creates a Client instance. From
there, you can use `client.Connect` to authenticate with the API.

```go
client, err := gonduit.Dial("https://phabricator.psyduck.info")

err = client.Connect("USERNAME", "CERTIFICATE")
```

## Errors

Any conduit error response will be returned as a `core.ConduitError` type:

```go
client, err := gonduit.Dial("https://phabricator.psyduck.info")
err = client.Connect("USERNAME", "CERTIFICATE")

ce, ok := err.(*core.ConduitError)
if ok {
	println("code: " + ce.Code())
	println("info: " + ce.Info())
}

// Or, use the built-in utility function:
if core.IsConduitError(err) {
	// do something else
}
```

## Supported Calls

All the supported API calls are available in the `Client` struct. Every
function is named after the Conduit method they call: For `phid.query`, we have
`Client.PHIDQuery`. The same applies for request and responses:
`requests.PHIDQueryRequest` and `responses.PHIDQueryResponse`.

Additionally, every general request method has the following signature:

```go
func (c *Conn) ConduitMethodName(req Request) (Response, error)
```

Some methods may also have specialized functions, you shhould refer the GoDoc
for more information on how to use them.

### List of supported calls:

- conduit.connect
- differential.query
- diffusion.querycommit
- file.download
- macro.creatememe
- maniphest.query
- paste.create
- paste.query
- phid.lookup
- phid.query
- repository.query

# Arbitrary calls

If you need to call an API method that is not supported by this client library,
you can use the `client.Call` method to make arbitrary calls.

You will need to provide a struct with the request body and a struct for the
response. The request has to be able to be able to be serialized into JSON,
and the response has be able to be unserialized from JSON.

```go
type phidLookupRequest struct {
	Names   []string         `json:"names"`
	Session *gonduit.Session `json:"__conduit__"`
}

type phidLookupResponse map[string]*struct{
	URI      string `json:"uri"`
	FullName string `json:"fullName"`
	Status   string `json:"status"`
}

req := &phidLookupRequest {
	Names: []string{"T1"},
	Session: client.Session,
}
var res phidLookupResponse

err := client.Call("phid.lookup", req, &res)
```
