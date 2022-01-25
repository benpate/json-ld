# JSON-LD

## Simple JSON-LD in Go

This is an experimental package for working with JSON-LD data in Go.  

## But Why?

There are other packages that do this quite well, so why make another one?  I needed something to dynamically evaluate ActivityStream data without requireing a specific vocabulary at compile time.  Unlike [the fantastic go-fed](https://go-fed.org) library, which uses automated code generation to create type-safe definitions for a specific library, this package evaluates generic data structures at run time.  This probably makes it slower, but the benefit is an easy-to-use API for safely accessing data within a JSON-LD document.
