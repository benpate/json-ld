# JSON-LD 🔗

## Simple JSON-LD in Go

This is an experimental package for working with [JSON-LD](https://json-ld.org) documents in Go.  

## Experimental, Pre-Alpha Quality

Please do not use this software until it is complete.

## What It Does

JSON-LD is a very flexible format that does not fit well into Go's strict type system.  This library provides a set of low level functions for accessing this data in a friendly and null-safe way.

Here's an Example:
```json
{
	"@context": "https://www.w3.org/ns/activitystreams",
	"summary": "Sally offered the Foo object",
	"type": "Offer",
	"actor": "https://sally.example.org",
	"object": "http://example.org/foo"
}
```

```go
object.Property("actor").AsString("id") // returns "https://sally.example.org"
object.Property("actor").AsObject("id") // returns the object {"id":"https://sally.example.org"}
```

## But Why?

There are other packages that do this quite well, so why make another one?  I needed something to dynamically evaluate ActivityStream data without requireing a specific vocabulary at compile time.  Unlike [the fantastic go-fed](https://go-fed.org) library, which uses automated code generation to create type-safe definitions for a specific library, this package evaluates generic data structures at run time.  This probably makes it slower, but the benefit is an easy-to-use API for safely accessing data within a JSON-LD document.

## Pull Requests Welcome

If this library is useful (or might be useful) to you, feel free to contribute with requests, fixes, and enhancements.  We're all in this together!
