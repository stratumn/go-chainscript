# Go ChainScript

[![GoDoc](https://godoc.org/github.com/stratumn/go-chainscript?status.svg)](https://godoc.org/github.com/stratumn/go-chainscript)
[![build status](https://travis-ci.org/stratumn/go-chainscript.svg?branch=master)](https://travis-ci.org/stratumn/go-chainscript)
[![codecov](https://codecov.io/gh/stratumn/go-chainscript/branch/master/graph/badge.svg)](https://codecov.io/gh/stratumn/go-chainscript)
[![Go Report Card](https://goreportcard.com/badge/github.com/stratumn/go-chainscript)](https://goreportcard.com/report/github.com/stratumn/go-chainscript)

Official Go implementation of [ChainScript](https://github.com/stratumn/chainscript).
This is the recommended way to use ChainScript in your Go projects.

However, it is opinionated and tries to keep the application data as abstract
as possible. Some applications will have different requirements and might
want to optimize for specific use-cases. If you are in that case, don't
hesitate to implement your own ChainScript library.
If you do so, don't forget to set the `client_id` field to make it easy for
others to deserialize and validate your data.

## Updating ChainScript

The ChainScript definitions are imported as a `git subtree`.
Changes to the protobuf files should be done in the
[ChainScript repository](https://github.com/stratumn/chainscript).

To get the latest ChainScript definitions in this project, run:

```bash
make update_chainscript
```
